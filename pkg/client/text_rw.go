package client

import (
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	limitlinereader "github.com/p4rthk4/u2smtp/pkg/limit_line_reader"
)

type TextReaderWriter struct {
	rwc     *ReadWriteClose
	t       *textproto.Conn
	netConn net.Conn
}

type ReadWriteClose struct {
	io.Reader
	io.Writer
	io.Closer
}

type EnhancedCode [3]int
type SMTPServerError struct {
	Code         int
	EnhancedCode EnhancedCode
	Message      string
}

func (err *SMTPServerError) Error() string {
	s := fmt.Sprintf("SMTP error %03d", err.Code)
	if err.Message != "" {
		s += ": " + err.Message
	}
	return s
}

func newTextReaderWriter(conn net.Conn) *TextReaderWriter {

	textReader := limitlinereader.LimitLineReader{
		Reader:      conn,
		MaxLineSize: 2000, // Doubled maximum line length per RFC 5321 (Section 4.5.3.1.6)
	}

	rwc := ReadWriteClose{
		Reader: &textReader,
		Writer: conn,
		Closer: conn,
	}

	text := textproto.NewConn(rwc)

	return &TextReaderWriter{
		rwc:     &rwc,
		t:       text,
		netConn: conn,
	}
}

func (rw *TextReaderWriter) readResponse(expectCode int) (int, string, error) {
	rw.netConn.SetReadDeadline(time.Now().Add(10 * time.Second)) // TODO: get in config file
	defer rw.netConn.SetReadDeadline(time.Time{})

	code, msg, err := rw.t.ReadResponse(expectCode)
	if protoErr, ok := err.(*textproto.Error); ok {
		err = toSMTPServerErr(protoErr)
	}
	return code, msg, err
}

func toSMTPServerErr(protoErr *textproto.Error) *SMTPServerError {
	smtpErr := &SMTPServerError{
		Code:    protoErr.Code,
		Message: protoErr.Msg,
	}

	parts := strings.SplitN(protoErr.Msg, " ", 2)
	if len(parts) != 2 {
		return smtpErr
	}

	enchCode, err := parseEnhancedCode(parts[0])
	if err != nil {
		return smtpErr
	}

	msg := parts[1]

	// Per RFC 2034, enhanced code should be prepended to each line.
	msg = strings.ReplaceAll(msg, "\n"+parts[0]+" ", "\n")

	smtpErr.EnhancedCode = enchCode
	smtpErr.Message = msg
	return smtpErr
}

func parseEnhancedCode(s string) (EnhancedCode, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return EnhancedCode{}, fmt.Errorf("wrong amount of enhanced code parts")
	}

	code := EnhancedCode{}
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return code, err
		}
		code[i] = num
	}
	return code, nil
}