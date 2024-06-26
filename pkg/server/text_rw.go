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

	textReader := LimitReader{
		r:         conn,
		lineLimit: 2000, // Doubled maximum line length per RFC 5321 (Section 4.5.3.1.6)
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

func (rw *TextReaderWriter) reply(code int, format string, a ...any) {
	rw.t.PrintfLine("%d %s", code, fmt.Sprintf(format, a...))
}

func (rw *TextReaderWriter) greet() {
	rw.t.PrintfLine("%d %s %s", 220, config.ConfOpts.HostName, config.ConfOpts.ClientGreet)
}

func (rw *TextReaderWriter) byyy() {
	rw.t.PrintfLine("%d %s", 221, config.ConfOpts.ClientByyy)
}

func (rw *TextReaderWriter) busy() {
	rw.t.PrintfLine("%d %s Service not available, max clients exceeded", 421, config.ConfOpts.HostName)
}

func (rw *TextReaderWriter) timeout() {
	rw.t.PrintfLine("%d %s Error: timeout exceeded", 421, config.ConfOpts.HostName)
}

func (rw *TextReaderWriter) longLine() {
	rw.t.PrintfLine("%d %s Error: too long line", 500, config.ConfOpts.HostName)
}

func (rw *TextReaderWriter) syntaxError(format string, a ...any) {
	rw.t.PrintfLine("%d %s", 501, fmt.Sprintf(format, a...))
}

func (rw *TextReaderWriter) cmdNotRecognized() {
	rw.t.PrintfLine("500 Error: command not recognized")
}

func (rw *TextReaderWriter) cmdNotImplemented() {
	rw.t.PrintfLine("502 Error: command not implemented")
}

func (rw *TextReaderWriter) readLine() (string, error) {
	rw.setTimeout(2 * time.Minute)
	defer rw.clearTimeout()

	return rw.t.ReadLine()
}

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
