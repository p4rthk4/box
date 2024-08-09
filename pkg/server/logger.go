package server

import (
	"os"

	"github.com/p4rthk4/u2smtp/pkg/logx"
)

// TODO dont set log from server pkg, use by extrenal program

func (s *SMTPServer) SetLogger() {
	file := os.Stdout
	s.logFile = file

	if !config.Dev {
		file, err := os.OpenFile(config.LogDirPath+"/"+config.LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logx.LogError("Error opening file:", err)
			return
		}
		s.logFile = file
	}

	s.logger = logx.NewLogger(s.logFile)
}
