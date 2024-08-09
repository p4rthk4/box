package logx

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Log struct {
	loggerInfo    *log.Logger
	loggerWarn    *log.Logger
	loggerError   *log.Logger
	loggerSuccess *log.Logger
	logFile       *os.File
}

func NewLogger(logFile *os.File) *Log {
	return NewLoggerWithPrefix(logFile, "")
}

func NewLoggerWithPrefix(logFile *os.File, prefix string) *Log {

	errorColor := color.New(color.FgRed).SprintFunc()
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	InfoColor := color.New(color.FgMagenta).SprintFunc()
	successColor := color.New(color.FgHiGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	var loggerInfo *log.Logger
	var loggerWarn *log.Logger
	var loggerError *log.Logger
	var loggerSuccess *log.Logger

	if prefix == "" {
		loggerInfo = log.New(logFile, bold(InfoColor("INFO: ")), log.LstdFlags)
		loggerWarn = log.New(logFile, bold(warnColor("WARN: ")), log.LstdFlags)
		loggerError = log.New(logFile, bold(errorColor("ERROR: ")), log.LstdFlags)
		loggerSuccess = log.New(logFile, bold(successColor("SUCCESS: ")), log.LstdFlags)

	} else {
		loggerInfo = log.New(logFile, bold(InfoColor(fmt.Sprintf("INFO %s: ", prefix))), log.LstdFlags)
		loggerWarn = log.New(logFile, bold(warnColor(fmt.Sprintf("WARN %s: ", prefix))), log.LstdFlags)
		loggerError = log.New(logFile, bold(errorColor(fmt.Sprintf("ERROR %s: ", prefix))), log.LstdFlags)
		loggerSuccess = log.New(logFile, bold(successColor(fmt.Sprintf("SUCCESS %s: ", prefix))), log.LstdFlags)
	}

	return &Log{
		loggerInfo:    loggerInfo,
		loggerWarn:    loggerWarn,
		loggerError:   loggerError,
		loggerSuccess: loggerSuccess,
		logFile:       logFile,
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

// dont end with "\n" it add automatic
func (l *Log) Success(format string, a ...any) {

	logString := fmt.Sprintf(format, a...)
	l.loggerSuccess.Println(logString)
}

func (l *Log) GetNewWithPrefix(prefix string) *Log {
	return NewLoggerWithPrefix(l.logFile, prefix)
}
