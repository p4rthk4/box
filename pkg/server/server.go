package server

import (
	"net"
	"sync"

	"github.com/p4rthk4/u2smtp/pkg/logx"
)

type SMTPServer struct {
	Host   string
	Port   int
	IsIPv6 bool

	listener net.Listener

	ServerWaitGroup *sync.WaitGroup
	logger          *logx.Log
}

var clientCount int = 0 // goble and overridable parallels client connection count

func (s *SMTPServer) SetLogger(l *logx.Log) {
	s.logger = l
}
