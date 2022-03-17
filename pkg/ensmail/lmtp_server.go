package ensmail

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/go-kit/log"
	"github.com/google/uuid"
)

// ResolveFunc resolves the local-part of an incoming email address to
// a forward email address.
type ResolveFunc func(context.Context, string) (string, error)

// NewForwarderClient returns a Forwarder, used to forward emails after
// address resolution.
type NewForwarderClient func() (ForwarderClient, error)

// ForwarderClient receives SMTP commands to forward emails.
type ForwarderClient interface {
	Mail(from string, opts *smtp.MailOptions) error
	Rcpt(to string) error
	LMTPData(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error)
	Reset() error
	Close() error
}

// LMTPResolveForwarder is an LMTP server which receives mail on a
// unix socket, resolves all mail receipients of that mail to another
// email address (recipients are based on the SMTP envelope "RCPT TO"
// command), and forwards the mail, with newly resolved recipients,
// over LMTP to a "Forwarder".
type LMTPResolveForwarder struct {
	logger       log.Logger
	srv          *smtp.Server
	resolver     ResolveFunc
	newForwarder NewForwarderClient
}

func NewLMTPServer(logger log.Logger, r ResolveFunc, nf NewForwarderClient) (*LMTPResolveForwarder, error) {
	l := LMTPResolveForwarder{
		logger:       log.With(logger, "app", "ensmail"),
		resolver:     r,
		newForwarder: nf,
	}
	// TODO: set timeouts? set max bytes received?
	l.srv = smtp.NewServer(&l)
	l.srv.LMTP = true
	return &l, nil
}

// Serve accepts incoming LMTP connections on the unix domain socket
// listener l.  Serve blocks until Close is called.
func (s *LMTPResolveForwarder) Serve(l net.Listener) error {
	if l.Addr().Network() != "unix" {
		return errors.New("not a unix domian socket listener")
	}
	s.logger.Log("serve", fmt.Sprintf("%s://%s", l.Addr().Network(), l.Addr().String()))
	return s.srv.Serve(l)
}

// Close immediately closes all active server connections, and causes
// Serve to return.
func (s *LMTPResolveForwarder) Close() error {
	s.logger.Log("serve", "close")
	return s.srv.Close()
}

type session struct {
	logger     log.Logger
	resolver   ResolveFunc
	unresolved map[string]string // k: resolved addr, v: unresolved addr
	forwarder  ForwarderClient
}

// NewSession implements the smtp.Backend interface, and is called for
// each new connection made to LMTP server.  A new forwarder client is
// created for each new session.
func (s *LMTPResolveForwarder) NewSession(c smtp.ConnectionState, hostname string) (smtp.Session, error) {
	fwdr, err := s.newForwarder()
	if err != nil {
		s.logger.Log("call", "s.newForwarder", "err", err)
		return nil, err
	}

	return &session{
		logger:     log.With(s.logger, "sessid", uuid.New().String()[:8]),
		resolver:   s.resolver,
		forwarder:  fwdr,
		unresolved: make(map[string]string),
	}, nil
}

func (s *session) Reset() {
	s.logger.Log("smtp", "RESET")
	s.forwarder.Reset()
}

func (s *session) AuthPlain(username, password string) error {
	return smtp.ErrAuthUnsupported
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	s.logger.Log("smtp", "MAIL", "from", from)
	return s.forwarder.Mail(from, opts)
}

// Rcpt will resolve "to", and pass the resolved value to the
// forwarder.
func (s *session) Rcpt(to string) error {
	logger := log.With(s.logger, "smtp", "RCPT", "to", to)

	at := strings.LastIndex(to, "@")
	if at <= 0 {
		logger.Log("err", "invalid addr")
		return fmt.Errorf("invalid recipient email: %s", to)
	}

	// TODO: use proper context
	resolved, err := s.resolver(context.Background(), to[:at])
	if err != nil {
		logger.Log("call", "s.resolver", "err", err)
		return err
	}
	logger = log.With(logger, "resolved", resolved)

	// TODO: what happens if s.unresolved[resolved] != ""?
	s.unresolved[resolved] = to

	if err := s.forwarder.Rcpt(resolved); err != nil {
		logger.Log("call", "s.forwarder.Rcpt", "err", err)
		return err
	}

	logger.Log("forward", "success")
	return nil
}

func (s *session) Data(r io.Reader) error {
	return errors.New("LMTPData method should be called")
}

// LMTPData copies data from r into forwarder DATA, waits for return
// status for every recipient.  It returns err only if forwarder DATA
// call fails.
func (s *session) LMTPData(r io.Reader, status smtp.StatusCollector) error {
	type statusRsp struct {
		rcpt string
		err  error
	}
	logger := log.With(s.logger, "smtp", "DATA")

	// Collect data responses per recipient.
	// TODO: this is subtly broken, because it's possible that Rcpt is
	// called with same "to" string, multiple times.  In that case,
	// status.SetStatus is supposed to be called multiple times for
	// each rcpt.
	dataRsps := make(chan statusRsp, len(s.unresolved))

	w, err := s.forwarder.LMTPData(func(rcpt string, serr *smtp.SMTPError) {
		// Convert half-nil serr to full-nil err interface value
		var err error
		if serr != nil {
			err = serr
		}
		dataRsps <- statusRsp{rcpt, err}
	})
	if err != nil {
		logger.Log("call", "s.forwarder.LMTPData", "err", err)
		return err
	}

	// TODO add "Received:" header?  Or other header to document resolution?

	// Copy received data to forwarding server.
	n, err := io.Copy(w, r)
	w.Close()
	if err != nil {
		logger.Log("call", "io.Copy", "err", err)
		return err
	}

	// Wait for all statuses to return, and call SetStatus appropriately.
	for range s.unresolved {
		select {
		case rsp := <-dataRsps:
			status.SetStatus(s.unresolved[rsp.rcpt], rsp.err)
			delete(s.unresolved, rsp.rcpt)
		// TODO: This timeout should not be hardcoded.  What's a good
		// value for this?
		case <-time.After(5 * time.Second):
			var missingRcpt strings.Builder
			for _, missing := range s.unresolved {
				fmt.Fprintf(&missingRcpt, "%s, ", missing)
			}
			err := fmt.Errorf("timeout waiting for forward LMTP status: %s", strings.TrimRight(missingRcpt.String(), ", "))
			logger.Log("call", "<-dataRsps", "err", err)
			return err
		}
	}

	logger.Log("forward", "success", "bytes", n)
	return nil
}

func (s *session) Logout() error {
	s.logger.Log("smtp", "LOGOUT")
	return s.forwarder.Close()
}
