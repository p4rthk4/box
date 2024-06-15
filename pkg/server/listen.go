// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	// "fmt"
	"os"

	"github.com/p4rthk4/u2smtp/pkg/logx"
	reusesocket "github.com/p4rthk4/u2smtp/pkg/reuse_socket"
)

func (s *SMTPServer) Listen() {

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

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logx.LogError("fail to connect with client...", err)
			continue
		}

		go HandleNewConnection(conn, s.logger)
	}

}
