// U2SMTP - log
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package logx

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Log struct {
	loggerInfo  *log.Logger
	loggerWarn  *log.Logger
	loggerError *log.Logger
}

func NewLogger(logFile *os.File) *Log {

	errorColor := color.New(color.FgRed).SprintFunc()
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	InfoColor := color.New(color.FgMagenta).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	loggerInfo := log.New(logFile, bold(InfoColor("INFO: ")), log.LstdFlags)
	loggerWarn := log.New(logFile, bold(warnColor("WARN: ")), log.LstdFlags)
	loggerError := log.New(logFile, bold(errorColor("ERROR: ")), log.LstdFlags)

	return &Log{
		loggerInfo:  loggerInfo,
		loggerWarn:  loggerWarn,
		loggerError: loggerError,
	}

}

// dont end with "\n" it add automatic
func (l *Log) Info(format string, a ...any) {
	logString := fmt.Sprintf(format, a...)
	l.loggerInfo.Println(logString)
}

// dont end with "\n" it add automatic
func (l *Log) Warn(format string, a ...any) {
	logString := fmt.Sprintf(format, a...)
	l.loggerWarn.Println(logString)
}

// dont end with "\n" it add automatic
func (l *Log) Error(format string, a ...any) {

	logString := fmt.Sprintf(format, a...)
	l.loggerError.Println(logString)
}
