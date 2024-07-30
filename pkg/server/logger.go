package server

import (
	"net"
	"os"
	"sync"

	"github.com/p4rthk4/u2smtp/pkg/config"
	"github.com/p4rthk4/u2smtp/pkg/logx"
)

type SMTPServer struct {
	Host   string
	Port   int
	IsIPv6 bool

	// server Listener
	listener net.Listener

	ServerWaitGroup *sync.WaitGroup
	logFile         *os.File
	logger          *logx.Log
}

var clientCount int = 0 // parallels client connection count

func (s *SMTPServer) SetLogger() {

	file := os.Stdout
	s.logFile = file

	if !config.ConfOpts.Dev {
		file, err := os.OpenFile(config.ConfOpts.LogDirPath+"/"+config.ConfOpts.LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logx.LogError("Error opening file:", err)
			return
		}
		s.logFile = file
	}

	s.logger = logx.NewLogger(s.logFile)

}
