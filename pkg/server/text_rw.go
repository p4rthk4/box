package server

import (
	"fmt"
	"io"
	"net"
	"net/textproto"
	"time"

	"github.com/p4rthk4/u2smtp/pkg/config"
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

// get new reader and write form net.Conn or io.ReadWriter
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

func (rw *TextReaderWriter) reply(code int, format string, a ...any) {
	rw.t.PrintfLine("%d %s", code, fmt.Sprintf(format, a...))
}

func (rw *TextReaderWriter) greet(hostname string) {
	rw.t.PrintfLine("%d %s %s", 220, hostname, config.ConfOpts.ClientGreet)
}

func (rw *TextReaderWriter) byyy() {
	rw.t.PrintfLine("%d %s", 221, config.ConfOpts.ClientByyy)
}

func (rw *TextReaderWriter) busy(hostname string) {
	rw.t.PrintfLine("%d %s Service not available, max clients exceeded", 421, hostname)
}

func (rw *TextReaderWriter) timeout(hostname string) {
	rw.t.PrintfLine("%d %s Error: timeout exceeded", 421, hostname)
}

func (rw *TextReaderWriter) longLine(hostname string) {
	rw.t.PrintfLine("%d %s Error: too long line", 500, hostname)
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
