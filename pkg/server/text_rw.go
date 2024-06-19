// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"fmt"
	"io"
	"net"
	"net/textproto"
	"time"

	"github.com/p4rthk4/u2smtp/pkg/config"
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

// get new reader and write form net.Conn or io.ReadWriter
func newTextReaderWriter(conn net.Conn) *TextReaderWriter {

	rwc := ReadWriteClose{
		Reader: conn,
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

// send reply
func (rw *TextReaderWriter) reply(code int, format string, a ...any) {
	rw.t.PrintfLine("%d %s", code, fmt.Sprintf(format, a...))
}

// send greet to cliet (220)
func (rw *TextReaderWriter) greet() {
	rw.t.PrintfLine("%d %s %s", 220, config.ConfOpts.HostName, config.ConfOpts.ClientGreet)
}

// send byyy (221)
func (rw *TextReaderWriter) byyy() {
	rw.t.PrintfLine("%d %s", 221, config.ConfOpts.ClientByyy)
}

// send busy (421)
func (rw *TextReaderWriter) busy() {
	rw.t.PrintfLine("%d %s Service not available, max clients exceeded", 421, config.ConfOpts.HostName)
}

// send timeout (421)
func (rw *TextReaderWriter) timeout() {
	rw.t.PrintfLine("%d %s Error: timeout exceeded", 421, config.ConfOpts.HostName)
}

// syntax error
func (rw *TextReaderWriter) syntaxError(format string, a ...any) {
	rw.t.PrintfLine("%d %s", 501, fmt.Sprintf(format, a...))
}

// command not recognized (502)
func (rw *TextReaderWriter) cmdNotRecognized() {
	rw.t.PrintfLine("500 Error: command not recognized")
}

// command not implemented (502)
func (rw *TextReaderWriter) cmdNotImplemented() {
	rw.t.PrintfLine("502 Error: command not implemented")
}

// read line end with \n
func (rw *TextReaderWriter) readLine() (string, error) {
	rw.setTimeout(2 * time.Minute)
	defer rw.clearTimeout()

	return rw.t.ReadLine()
}

// read data end with \r\n.\r\n
func (rw *TextReaderWriter) readData() ([]byte, error) {
	rw.setTimeout(15 * time.Minute)
	defer rw.clearTimeout()

	return rw.t.ReadDotBytes()
}

// set read time out
func (rw *TextReaderWriter) setTimeout(t time.Duration) {
	rw.netConn.SetReadDeadline(time.Now().Add(t))
}

// clear read time out
func (rw *TextReaderWriter) clearTimeout() {
	rw.netConn.SetReadDeadline(time.Time{})
}
