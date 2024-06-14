// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	// "fmt"
	"os"

	"github.com/p4rthk4/u2smtp/pkg/log"
	reusesocket "github.com/p4rthk4/u2smtp/pkg/reuse_socket"
)

func (s *SMTPServer) Listen() {

	address := s.getHostAddress()

	listener, err := reusesocket.Listen("tcp", address)
	if err != nil {
		log.LogError("server listen faild", err)
		os.Exit(1)
	}

	s.listener = listener

}

func (s *SMTPServer) AcceptConnections() {

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.LogError("fail to connect with client...", err)
			continue
		}

		go HandleNewConnection(conn)

	}

}
