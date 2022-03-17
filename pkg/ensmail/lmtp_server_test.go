package ensmail

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/go-kit/log"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockForwarder struct {
	mailFunc  func(from string, opts *smtp.MailOptions) error
	rcptFunc  func(to string) error
	dataFunc  func(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error)
	resetFunc func() error
	closeFunc func() error
}

func (m mockForwarder) Mail(from string, opts *smtp.MailOptions) error {
	if m.mailFunc != nil {
		return m.mailFunc(from, opts)
	}
	return nil
}

func (m mockForwarder) Rcpt(to string) error {
	if m.rcptFunc != nil {
		return m.rcptFunc(to)
	}
	return nil
}

func (m mockForwarder) LMTPData(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error) {
	if m.dataFunc != nil {
		return m.dataFunc(statusCb)
	}
	return nil, nil
}

func (m mockForwarder) Reset() error {
	if m.resetFunc != nil {
		return m.resetFunc()
	}
	return nil
}

func (m mockForwarder) Close() error {
	if m.closeFunc != nil {
		return m.closeFunc()
	}
	return nil
}

type sessionRecorder struct {
	sessions []*testSession
}

type testSession struct {
	mockForwarder
	From string
	To   []string
	Data bytes.Buffer
}

type Closer struct {
	io.Writer
	closeFunc func() error
}

func (c Closer) Close() error {
	return c.closeFunc()
}

func (sr *sessionRecorder) Forwarder() (ForwarderClient, error) {
	var ts testSession
	ts.mailFunc = func(from string, opts *smtp.MailOptions) error {
		ts.From = from
		return nil
	}

	ts.rcptFunc = func(to string) error {
		if ts.To == nil {
			ts.To = []string{to}
		} else {
			ts.To = append(ts.To, to)
		}
		return nil
	}

	ts.dataFunc = func(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error) {
		return Closer{
			Writer: &ts.Data,
			closeFunc: func() error {
				for _, rcpt := range ts.To {
					statusCb(rcpt, nil)
				}
				return nil
			},
		}, nil
	}

	sr.sessions = append(sr.sessions, &ts)
	return ts, nil
}

func (sr *sessionRecorder) check(t *testing.T, exp []*testSession) {
	cmpBuf := cmp.Comparer(func(x, y bytes.Buffer) bool {
		return x.String() == y.String()
	})

	if !cmp.Equal(exp, sr.sessions, cmpBuf, cmpopts.IgnoreFields(testSession{}, "mockForwarder")) {
		t.Errorf("forwardedSessions (-want, +got) %s", cmp.Diff(exp, sr.sessions, cmpBuf, cmpopts.IgnoreFields(testSession{}, "mockForwarder")))
	}
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
			continue
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

var (
	// logger = log.With(log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr)), "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	logger = log.NewNopLogger()
)

func TestLMTPServer(t *testing.T) {

	// Upon receiving LHLO, a connection to the forwarding server is
	// established.  If this connection can't be established, bubble
	// error back to original LHLO sender.
	t.Run("errForwarderCl", func(t *testing.T) {
		// Create server
		srv, err := NewLMTPServer(logger, nil, func() (ForwarderClient, error) {
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

	// After sender finishes completes DATA command, if forwarding
	// server has issue, ensure that error bubbles back to sender.
	t.Run("errForwardData", func(t *testing.T) {
		resolver := func(ctx context.Context, in string) (string, error) { return in, nil }

		errBadForward := errors.New("bad forward")
		srv, err := NewLMTPServer(logger, resolver, func() (ForwarderClient, error) {
			return mockForwarder{
				dataFunc: func(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error) {
					return nil, errBadForward
				},
			}, nil
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

		closed := make(chan error)
		go func() {
			closed <- srv.Serve(l)
		}()

		if err := sendMail(sock, "sender@public.com", []string{"rcpt@ensmail.org"}, testMsg); err == nil {
			// 220  ESMTP Service Ready
			// LHLO localhost
			// 250-Hello localhost
			// 250-PIPELINING
			// 250-8BITMIME
			// 250-ENHANCEDSTATUSCODES
			// 250-CHUNKING
			// 250 SIZE
			// MAIL FROM:<sender@public.com> BODY=8BITMIME
			// 250 2.0.0 Roger, accepting mail from <sender@public.com>
			// RCPT TO:<rcpt@ensmail.org>
			// 250 2.0.0 I'll make sure <rcpt@ensmail.org> gets this
			// DATA
			// 354 2.0.0 Go ahead. End your data with <CR><LF>.<CR><LF>
			// 554 5.0.0 <rcpt@ensmail.org> Error: transaction failed, blame it on the weather: bad forward
			t.Fatal("unexpected err:", err)
		}
	})

	// Some rcpt resolve, some don't.
	t.Run("errMultiRcptResolve", func(t *testing.T) {
		resolver := func(ctx context.Context, in string) (string, error) {
			if strings.HasPrefix(in, "BAD") {
				return "", errors.New("invalid resolve input")
			}
			return fmt.Sprintf("RESOLVED%s@resolved.test", in), nil
		}

		var recorder sessionRecorder
		srv, err := NewLMTPServer(logger, resolver, recorder.Forwarder)
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

		if err := sendMail(sock, "sender@public.com", []string{"rcpt1@ensmail.org", "BADrcpt2@ensmail.org", "rcpt3@ensmail.org", "BADrcpt4@ensmail.org"}, testMsg); err != nil {
			t.Fatal("unexpected err:", err)
		}
		time.Sleep(100 * time.Millisecond)

		if err := srv.Close(); err != nil {
			t.Fatal(err)
		}

		recorder.check(t, []*testSession{
			{
				From: "sender@public.com",
				To: []string{
					"RESOLVEDrcpt1@resolved.test",
					"RESOLVEDrcpt3@resolved.test",
				},
				Data: *bytes.NewBuffer(testMsg),
			},
		})
	})

	// LMTP provides per-recipient status on data.  If some fail,
	// ensure error is returned correctly.
	t.Run("errMultiRcptForward", func(t *testing.T) {
		resolver := func(ctx context.Context, in string) (string, error) {
			return in, nil
		}

		srv, err := NewLMTPServer(logger, resolver, func() (ForwarderClient, error) {
			rcpts := make([]string, 0)
			return mockForwarder{
				// Collects rcpts
				rcptFunc: func(to string) error {
					rcpts = append(rcpts, to)
					return nil
				},
				dataFunc: func(statusCb func(rcpt string, status *smtp.SMTPError)) (io.WriteCloser, error) {
					// On Close(), iterate over rcpts, and if it starts with "BAD", return 500 SMTP err
					return Closer{
						Writer: io.Discard,
						closeFunc: func() error {
							for _, rcpt := range rcpts {
								if strings.HasPrefix(rcpt, "BAD") {
									statusCb(rcpt, &smtp.SMTPError{
										Code:    500,
										Message: "test bad data",
									})
									continue
								}
								statusCb(rcpt, nil)
							}
							return nil
						},
					}, nil
				},
			}, nil
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

		closed := make(chan error)
		go func() {
			closed <- srv.Serve(l)
		}()

		if err := sendMail(sock, "sender@public.com", []string{"rcpt1@ensmail.org", "BADrcpt2@ensmail.org", "rcpt3@ensmail.org", "BADrcpt4@ensmail.org"}, testMsg); err == nil {
			t.Fatal("expected non-nil err")
		}

		if err := srv.Close(); err != nil {
			t.Fatal(err)
		}
	})

	// Mail with single rcpt is correctly resolved and forwarded.
	t.Run("success", func(t *testing.T) {
		resolver := func(ctx context.Context, in string) (string, error) {
			return fmt.Sprintf("RESOLVED%s@resolved.test", in), nil
		}

		var recorder sessionRecorder
		srv, err := NewLMTPServer(logger, resolver, recorder.Forwarder)
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

		if err := sendMail(sock, "sender1@public.com", []string{"rcpt@ensmail.org"}, testMsg); err != nil {
			t.Fatal("unexpected err:", err)
		}

		if err := sendMail(sock, "sender2@public.com", []string{"rcpt1@ensmail.org", "rcpt2@ensmail.org", "rcpt3@ensmail.org"}, testMsg); err != nil {
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
				From: "sender1@public.com",
				To:   []string{"RESOLVEDrcpt@resolved.test"},
				Data: *bytes.NewBuffer(testMsg),
			},
			{
				From: "sender2@public.com",
				To: []string{
					"RESOLVEDrcpt1@resolved.test",
					"RESOLVEDrcpt2@resolved.test",
					"RESOLVEDrcpt3@resolved.test",
				},
				Data: *bytes.NewBuffer(testMsg),
			},
		})
	})
}
