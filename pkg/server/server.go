// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"net"
	"sync"
)

type SMTPServer struct {
	Host   string
	Port   int
	IsIPv6 bool

	listener net.Listener
	
	ServerWaitGroup *sync.WaitGroup
}
