package ensmail

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"testing"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/google/go-cmp/cmp"
)

type sessionRecorder struct {
	sessions []*testSession
}

func (sr *sessionRecorder) Forwarder() (ForwarderClient, error) {
	var ts testSession
	sr.sessions = append(sr.sessions, &ts)
	return &ts, nil
}

func (sr *sessionRecorder) check(t *testing.T, exp []*testSession) {
	opt := cmp.Comparer(func(x, y bytes.Buffer) bool {
		return x.String() == y.String()
	})

	if !cmp.Equal(exp, sr.sessions, cmp.AllowUnexported(testSession{}), opt) {
		t.Errorf("forwardedSessions (-want, +got) %s", cmp.Diff(exp, sr.sessions, cmp.AllowUnexported(testSession{}), opt))
	}
}

type testSession struct {
	from string
	to   []string
	bytes.Buffer
}

func (s *testSession) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	return nil
}

func (s *testSession) Rcpt(to string) error {
	if s.to == nil {
		s.to = []string{to}
	} else {
		s.to = append(s.to, to)
	}
	return nil
}

func (s *testSession) LMTPData(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error) {
	statusCb(s.to[0], nil)
	return s, nil
}

func (s *testSession) Reset() error {
	return nil
}

func (s *testSession) Close() error {
	return nil
}

var testMsg = []byte("Received: from localhost (localhost [127.0.0.1]) by mx.maddy.test\r\n" +
	" (envelope-sender <sender@example.org>) with UTF8ESMTP id e6fa8a02; Fri, 25\r\n" +
	" Feb 2022 16:39:27 -0500\r\n" +
	"To: recipient@example.net\r\n" +
	"Subject: discount Gophers!\r\n" +
	"\r\n" +
	"This is the email body.\r\n")

func sendMail(sock string, from string, to []string, data []byte) error {
	conn, err := net.Dial("unix", sock)
	if err != nil {
		return err
	}
	cl, err := smtp.NewClientLMTP(conn, "ensmail-testclient.local")
	if err != nil {
		return err
	}

	if err := cl.Mail(from, nil); err != nil {
		return err
	}

	rcpts := make(map[string]chan error)
	for _, rcpt := range to {
		if err := cl.Rcpt(rcpt); err != nil {
			return err
		}
		rcpts[rcpt] = make(chan error, 1)
	}

	w, err := cl.LMTPData(func(rcpt string, status *smtp.SMTPError) {
		if status == nil {
			// Send a "full-nil" error interface value
			rcpts[rcpt] <- nil
		} else {
			rcpts[rcpt] <- status
		}
	})
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	for rcpt, errC := range rcpts {
		select {
		case err := <-errC:
			if err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			return fmt.Errorf("timeout for rcpt %s", rcpt)
		}
	}

	return nil
}

func TestLMTPServer(t *testing.T) {
	t.Run("errForwarderCl", func(t *testing.T) {
		// Create server
		srv, err := NewLMTPServer(nil, func() (ForwarderClient, error) {
			return nil, errors.New("TEST forward error")
		})
		if err != nil {
			t.Fatal(err)
		}

		// Serve on unix socket
		sock := filepath.Join(t.TempDir(), "lmtp.sock")
		l, err := net.Listen("unix", sock)
		if err != nil {
			t.Fatal(err)
		}
		defer l.Close()

		go srv.Serve(l)

		if err := sendMail(sock, "sender@public.com", []string{"rcpt1@ensmail.org"}, testMsg); err == nil {
			// The SMTP conversation is as follows:
			// S> 220  ESMTP Service Ready
			// C> LHLO localhost
			// S> 451 4.0.0 TEST forward error
			// C> HELO localhost
			// S> 500 5.5.1 This is a LMTP server, use LHLO
			// Client's LHLO command is sent when LMTPClient's Mail
			// method is called.  Upon receipt of LHLO,
			// LMTPResolveForwarder's NewSession method is called,
			// which returns the above "TEST forward error".
			// LMTPClient will attempt to degrade to a HELO command,
			// which results in the returns "500 use LHLO" error.
			t.Fatal("unexpected nil err")
		}

		srv.Close()
	})

	t.Run("singleRcpt", func(t *testing.T) {
		resolver := func(in string) (string, error) {
			return "RESOLVED" + in, nil
		}

		var recorder sessionRecorder
		srv, err := NewLMTPServer(resolver, recorder.Forwarder)
		if err != nil {
			t.Fatal(err)
		}

		// Serve on unix socket
		sock := filepath.Join(t.TempDir(), "lmtp.sock")
		l, err := net.Listen("unix", sock)
		if err != nil {
			t.Fatal(err)
		}
		defer l.Close()

		closed := make(chan error)
		go func() {
			closed <- srv.Serve(l)
		}()

		if err := sendMail(sock, "sender@public.com", []string{"rcpt@ensmail.org"}, testMsg); err != nil {
			t.Fatal("unexpected err:", err)
		}
		time.Sleep(100 * time.Millisecond)

		if err := srv.Close(); err != nil {
			t.Fatal(err)
		}

		select {
		case err := <-closed:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("server shutdown timeout")
		}

		recorder.check(t, []*testSession{
			{
				from:   "sender@public.com",
				to:     []string{"RESOLVEDrcpt@ensmail.org"},
				Buffer: *bytes.NewBuffer(testMsg),
			},
		})
	})
	t.Run("errForwardData", func(t *testing.T) {
	})
	t.Run("multiRcpt", func(t *testing.T) {
	})
	// Some rcpt resolve, some don't
	t.Run("multiRcptResolveErr", func(t *testing.T) {
	})
	// Some rcpt forward, some don't
	t.Run("multiRcptForwardErr", func(t *testing.T) {

	})
}
