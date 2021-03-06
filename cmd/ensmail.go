package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/emersion/go-smtp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/royalfork/ensmail/pkg/ensmail"
)

var version = "dev"

func main() {
	var (
		ENSRegistry       common.Address
		Web3RTCURL        string
		LMTPServerSocket  string
		LMTPForwardSocket string

		ensRegistry string
	)

	flag.StringVar(&ensRegistry, "ens", "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e", "ENS Registry address")
	flag.StringVar(&Web3RTCURL, "web3", "", "WebRTC URL for web3 (overwrites HTTP_WEB3_PROVIDER env var)")
	flag.StringVar(&LMTPServerSocket, "s", "/run/ensmail/ensmail.sock", "LMTP server listens on this socket")
	flag.StringVar(&LMTPForwardSocket, "f", "/run/ensmail/forward.sock", "LMTP forwards mail to this socket")
	v := flag.Bool("v", false, "print version")
	flag.Parse()

	if *v {
		fmt.Println(version)
		os.Exit(0)
	}

	if Web3RTCURL == "" {
		Web3RTCURL = os.Getenv("HTTP_WEB3_PROVIDER")
	}

	ENSRegistry = common.HexToAddress(ensRegistry)

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger.Log("ens", ENSRegistry, "serveSocket", LMTPServerSocket, "fowardSocket", LMTPForwardSocket)

	client, err := ethclient.Dial(Web3RTCURL)
	if err != nil {
		logger.Log("call", "ethclient.Dial", "err", err)
		os.Exit(1)
	}

	resolver, err := ensmail.NewENSResolver(ENSRegistry, client)
	if err != nil {
		logger.Log("call", "ensmail.NewENSResolver", "err", err)
		os.Exit(1)
	}

	newForwarderClient := func() (ensmail.ForwarderClient, error) {
		conn, err := net.Dial("unix", LMTPForwardSocket)
		if err != nil {
			return nil, err
		}
		return smtp.NewClientLMTP(conn, "ensmail.local")
	}

	s, err := ensmail.NewLMTPServer(logger, resolver.Email, newForwarderClient)
	if err != nil {
		logger.Log("call", "ensmail.NewLMTPServer", "err", err)
		os.Exit(1)
	}

	l, err := net.Listen("unix", LMTPServerSocket)
	if err != nil {
		logger.Log("call", "new.Listen", "err", err)
		os.Exit(1)
	}
	defer l.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := s.Serve(l); err != nil {
			logger.Log("call", "s.Serve", "err", err)
			os.Exit(1)
		}
		wg.Done()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	s.Close()
	wg.Wait()
}
