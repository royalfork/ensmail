package ensmail_test

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"testing"
	"time"

	gsmtp "github.com/emersion/go-smtp"
	"github.com/foxcpp/go-mockdns"
	"github.com/foxcpp/maddy/tests"
	"github.com/royalfork/ensmail/pkg/ens"
	"github.com/royalfork/ensmail/pkg/ensmail"
)

var (
	tlsCertFile string
	tlsKeyFile  string
)

func init() {
	flag.StringVar(&tlsCertFile, "cert", "", "maddy TLS cert file (absolute path)")
	flag.StringVar(&tlsKeyFile, "key", "", "maddy TLS key file (absolute path)")
}

// TestIntegration runs maddy with the production config, runs the
// ENSMail forwarding LMTP server, and performs an end-to-end
// integration test.  SMTP mail is sent to maddy where the RCPT
// address is resolved by ENSMail, which then forwards the mail to the
// relevant remote SMTP server.
func TestIntegration(t *testing.T) {
	var (
		sender = "sender@anywhere.test"
		// Test message has rcptLabel local-part (rcptLabel@smtpDomain)
		rcptLabel = "ensrecipient"
		// Maddy receives mail from this domain.
		smtpDomain = "ensmail.test"
		rcpt       = fmt.Sprintf("%s@%s", rcptLabel, smtpDomain)
		// On this SMTP port
		smtpPort = 25252

		// ENS text/email record for rcptLabel.eth returns resolvedLabel@resolvedDomain.
		resolvedLabel  = "recipient"
		resolvedDomain = "public.test"
		resolvedRcpt   = fmt.Sprintf("%s@%s", resolvedLabel, resolvedDomain)

		// ENSMail listens for incoming LMTP connections on ensMailSock.
		ensMailSock = filepath.Join(t.TempDir(), "ensmail.sock")
		// Maddy listens for incoming LMTP connections on maddySock
		// for outgoing, forwarded mail.
		maddySock = filepath.Join(t.TempDir(), "maddy.sock")

		// Test message data.
		msg = []byte("To: ensrecipient@ensmail.test\r\n" +
			"Subject: discount Gophers!\r\n" +
			"\r\n" +
			"This is the email body.\r\n")
	)

	closeENS := runENSMail(t, rcptLabel, fmt.Sprintf("%s@%s", resolvedLabel, resolvedDomain), ensMailSock, maddySock)
	defer closeENS()

	remoteSMTPPort, closeMaddy := runMaddy(t, smtpPort, smtpDomain, resolvedDomain, ensMailSock, maddySock)
	defer closeMaddy()

	// Mock smtp server will receive the forwarded/resolved mail.
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", remoteSMTPPort))
	if err != nil {
		t.Fatal(err)
	}

	rcv := make(chan Session, 1)
	srv := gsmtp.NewServer(Backend{Received: rcv})
	go func() {
		if err := srv.Serve(l); err != nil {
			panic(err)
		}
	}()

	if err := gsmtp.SendMail(fmt.Sprintf("127.0.0.1:%d", smtpPort), nil, sender, []string{rcpt}, bytes.NewReader(msg)); err != nil {
		t.Fatal(err)
	}

	select {
	case rcvd := <-rcv:
		if rcvd.From != sender {
			t.Errorf("want sender: %s, got: %s", sender, rcvd.From)
		}
		if len(rcvd.To) != 1 {
			t.Errorf("want receivers: %d, got: %d", 1, len(rcvd.To))
		} else if rcvd.To[0] != resolvedRcpt {
			t.Errorf("want receiver: %s, got: %s", resolvedRcpt, rcvd.To[0])
		}

	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

// Run ensmail server builds a private ENS chain and registers
// rcptLabel.eth with resolved email text record set to resolvedEmail.
// The ENSMail LMTP server email resolver is set to an ENSResolver
// which queries the private chain. The ENSMail LMTP server creats a
// new go-smtp LMTP client for each incoming mail session (which it
// accepts from ensMailSock), which forwards the "resolved" mail
// messages over LMTP on maddySock.
func runENSMail(t *testing.T, rcptLabel, resolvedEmail, ensMailSock, maddySock string) (close func()) {
	testENS, err := ens.NewTest()
	if err != nil {
		t.Fatal(err)
	}
	rcptNode, err := testENS.Register(testENS.Accts[1].Addr, rcptLabel)
	if err != nil {
		t.Fatal(err)
	}
	if !testENS.Chain.Succeed(testENS.Registry.SetResolver(testENS.Accts[1].Auth, rcptNode, testENS.ResolverAddr)) {
		t.Fatal("unable to set resolver")
	}
	if !testENS.Chain.Succeed(testENS.Resolver.SetText(testENS.Accts[1].Auth, rcptNode, "email", resolvedEmail)) {
		t.Fatal("unable to set resolver")
	}

	resolver, err := ensmail.NewENSResolver(testENS.RegistryAddr, testENS.Chain)
	if err != nil {
		t.Fatal(err)
	}

	newForwarderClient := func() (ensmail.ForwarderClient, error) {
		conn, err := net.Dial("unix", maddySock)
		if err != nil {
			return nil, err
		}
		return gsmtp.NewClientLMTP(conn, "ensmail-testclient.local")
	}

	ensServer, err := ensmail.NewLMTPServer(resolver.Email, newForwarderClient)
	if err != nil {
		t.Fatal(err)
	}
	l, err := net.Listen("unix", ensMailSock)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		ensServer.Serve(l)
	}()

	return func() {
		ensServer.Close()
		l.Close()
	}
}

// runMaddy uses the maddy testing package to execute maddy with the
// production configs/maddy.conf, listens for SMTP on tcp smtpPort,
// accepts mail only for recipients in smtpDomain, forwards the mail
// over LMTP to ensMailSock, and listens for outgoing mail on LMTP
// maddySock.  Outgoing mail can only be sent to rcptDomain; a local
// DNS route is setup where rcptDomain resolves to 127.0.0.1. Outgoing
// mail is sent on remoteSMTPPort (which the maddy testing framework
// must set itself).
func runMaddy(t *testing.T, smtpPort int, smtpDomain, rcptDomain, ensMailSock, maddySock string) (remoteSMTPPort uint16, close func()) {
	const config = "../configs/maddy.conf"
	conf, err := ioutil.ReadFile(config)
	if err != nil {
		t.Fatal(err)
	}

	mdy := tests.NewT(t)

	// Because maddy must forward mail to the resolved rcpt domain,
	// set the resolved rcpt domain to localhost.
	mdy.DNS(map[string]mockdns.Zone{
		rcptDomain + ".": {
			MX: []net.MX{{Host: "mx." + rcptDomain + ".", Pref: 10}},
		},
		"mx." + rcptDomain + ".": {
			A: []string{"127.0.0.1"},
		},
	})

	remoteSMTPPort = mdy.Port("remote_smtp")

	mdy.Env("HOSTNAME=mx." + smtpDomain)
	mdy.Env("DOMAIN=" + smtpDomain)
	mdy.Env(fmt.Sprintf("SMTP_PORT=%d", smtpPort))
	mdy.Env("LMTP_ENSMAIL_SOCK=" + ensMailSock)
	mdy.Env("LMTP_FORWARD_SOCK=" + maddySock)
	mdy.Env("TLS_CERT_FILE=" + tlsCertFile)
	mdy.Env("TLS_KEY_FILE=" + tlsKeyFile)

	mdy.Config(string(conf))

	mdy.Run(1)

	return remoteSMTPPort, mdy.Close
}

// The Backend implements SMTP server methods.
type Backend struct {
	Received chan<- Session
}

func (b Backend) NewSession(c gsmtp.ConnectionState, hostname string) (gsmtp.Session, error) {
	return &Session{
		onData: func(s Session) {
			b.Received <- s
		},
	}, nil
}

// A Session is returned after EHLO.
type Session struct {
	From    string
	To      []string
	MsgData []byte
	onData  func(Session)
}

func (s *Session) AuthPlain(username, password string) error {
	return gsmtp.ErrAuthUnsupported
}

func (s *Session) Mail(from string, opts *gsmtp.MailOptions) error {
	s.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	s.To = append(s.To, to)
	return nil
}

func (s *Session) Data(r io.Reader) (err error) {
	s.MsgData, err = ioutil.ReadAll(r)
	s.onData(*s)
	return err
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
