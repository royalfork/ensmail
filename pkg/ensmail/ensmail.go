package ensmail

import (
	"errors"
	"io"
	"net"

	"github.com/emersion/go-smtp"
)

// ResolveFunc resolves an email address to a forward address.
type ResolveFunc func(string) (string, error)

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
	srv          *smtp.Server
	resolver     ResolveFunc
	newForwarder NewForwarderClient
}

func NewLMTPServer(r ResolveFunc, nf NewForwarderClient) (*LMTPResolveForwarder, error) {
	l := LMTPResolveForwarder{
		resolver:     r,
		newForwarder: nf,
	}
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
	return s.srv.Serve(l)
}

// Close immediately closes all active server connections, and causes
// Serve to return.
func (s *LMTPResolveForwarder) Close() error {
	return s.srv.Close()
}

type session struct {
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
		return nil, err
	}

	return &session{
		resolver:   s.resolver,
		forwarder:  fwdr,
		unresolved: make(map[string]string),
	}, nil
}

func (s *session) Reset() {
	s.forwarder.Reset()
}

func (s *session) AuthPlain(username, password string) error {
	return smtp.ErrAuthUnsupported
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return s.forwarder.Mail(from, opts)
}

// Rcpt will resolve "to", and pass the resolved value to the
// forwarder.
func (s *session) Rcpt(to string) error {
	resolved, err := s.resolver(to)
	if err != nil {
		return err
	}
	s.unresolved[resolved] = to
	return s.forwarder.Rcpt(resolved)
}

func (s *session) Data(r io.Reader) error {
	return errors.New("LMTPData method should be called")
}

func (s *session) LMTPData(r io.Reader, status smtp.StatusCollector) error {
	w, err := s.forwarder.LMTPData(func(rcpt string, err *smtp.SMTPError) {
		if err != nil {
			status.SetStatus(s.unresolved[rcpt], err)
			return
		}
		// Because err is half-nil, a full-nil err must be sent into
		// SetStatus.
		status.SetStatus(s.unresolved[rcpt], nil)
	})
	if err != nil {
		return err
	}
	defer w.Close()

	// TODO add "Received:" header?  Or other header to document resolution?

	_, err = io.Copy(w, r)
	return err
}

func (s *session) Logout() error {
	return s.forwarder.Close()
}
