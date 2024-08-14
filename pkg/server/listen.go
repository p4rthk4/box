package server

import (
	// "fmt"
	"fmt"
	"os"
	"time"

	"github.com/rellitelink/box/pkg/logx"
	reusesocket "github.com/rellitelink/box/pkg/reuse_socket"
)

func (s *SMTPServer) Listen() {

	fmt.Println(time.Now().UTC().String())
	fmt.Println(time.Now().String())

	address := s.getHostAddress()

	listener, err := reusesocket.Listen("tcp", address)
	if err != nil {
		logx.LogError("server listen faild", err)
		os.Exit(1)
	}

	s.listener = listener
	s.logger.Info("server start/listen on %s", address)
}

func (s *SMTPServer) AcceptConnections() {
	smtpServerPreProcess(s.logger) // pre process

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logx.LogError("fail to connect with client...", err)
			continue
		}

		go HandleNewConnection(conn, s.logger)
	}
}
