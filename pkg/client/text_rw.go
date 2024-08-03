package client

import (
	"bytes"
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
	rwc *ReadWriteClose
	t   *textproto.Conn
	*limitlinereader.LimitLineReader
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

func (err SMTPServerError) Error() string {
	s := fmt.Sprintf("SMTP error %03d", err.Code)
	if err.Message != "" {
		s += ": " + err.Message
	}
	return s
}

func newTextReaderWriter(conn net.Conn) *TextReaderWriter {
	textReader := &limitlinereader.LimitLineReader{
		Reader:      conn,
		MaxLineSize: 2000, // Doubled maximum line length per RFC 5321 (Section 4.5.3.1.6)
	}

	rwc := ReadWriteClose{
		Reader: textReader,
		Writer: conn,
		Closer: conn,
	}

	text := textproto.NewConn(rwc)
	return &TextReaderWriter{
		netConn:         conn,
		rwc:             &rwc,
		t:               text,
		LimitLineReader: textReader,
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

func (rw *TextReaderWriter) cmd(expectCode int, format string, args ...interface{}) (int, string, error) {
	id, err := rw.t.Cmd(format, args...)
	if err != nil {
		if protoErr, ok := err.(*textproto.Error); ok {
			err = toSMTPServerErr(protoErr)
		}
		return 0, "", err
	}

	rw.t.StartResponse(id)
	defer rw.t.EndResponse(id)

	return rw.readResponse(expectCode)
}

func (rw *TextReaderWriter) data(data []byte) (int, string, error) {
	dataReader := bytes.NewReader(data)
	dataWriter := rw.t.DotWriter()

	_, err := io.Copy(dataWriter, dataReader)
	if err != nil {
		return 0, "", err
	}

	err = dataWriter.Close()
	if err != nil {
		if protoErr, ok := err.(*textproto.Error); ok {
			err = toSMTPServerErr(protoErr)
		}
		return 0, "", err
	}

	return rw.readResponse(250)
}

func toSMTPServerErr(protoErr *textproto.Error) SMTPServerError {
	smtpErr := SMTPServerError{
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
