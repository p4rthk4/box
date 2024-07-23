package server

import (
	"fmt"
	"io"
	"net"

	"github.com/p4rthk4/u2smtp/pkg/config"
	limitlinereader "github.com/p4rthk4/u2smtp/pkg/limit_line_reader"
	"github.com/p4rthk4/u2smtp/pkg/logx"
	"github.com/p4rthk4/u2smtp/pkg/uid"
)

type MailForwardCode int

const (
	MailForwardFaild   MailForwardCode = iota // is mail not recive perfect or interrupt connection
	MailForwardSuccess                        // mail recive full
	MailForwardIdle                           // if connection close without send data
)

type Client struct {
	domain        string
	mailFrom      string
	recipients    []string
	data          []byte
	forwardStatus MailForwardCode
}

type Connection struct {
	conn net.Conn

	remoteAddress RoLAddress
	localAddress  RoLAddress
	text          *TextReaderWriter // text protocal for mail

	uid       string
	mailCount int // it is count of how many mail tranfare in this connection
	client    Client

	logger       *logx.Log
	serverLogger *logx.Log // print server level log
}

// handle new client connection
func HandleNewConnection(conn net.Conn, serverLogger *logx.Log) {
	connection := Connection{
		conn:         conn,
		serverLogger: serverLogger,
	}
	clientCount += 1

	err := connection.init()
	if err {
		// conn.Close() // importent
		return
	}

	connection.handle()
}

// init client connection
// return true if error
func (conn *Connection) init() bool {

	conn.text = newTextReaderWriter(conn.conn)

	uid, err := uid.GetNewId()
	if err != nil {
		conn.serverLogger.Error("generate email uid error: %v", err)
		return true
	}

	conn.mailCount = 1
	conn.uid = uid
	conn.logger = conn.serverLogger.GetNewWithPrefix(conn.uid)

	if ok := conn.remoteAddress.SetAddress(conn.conn.RemoteAddr().Network(), conn.conn.RemoteAddr().String()); !ok {
		conn.logger.Warn("no PTR record or faild to find PTR records of %s", conn.remoteAddress.String())
	}

	if ok := conn.localAddress.SetAddress(conn.conn.LocalAddr().Network(), conn.conn.LocalAddr().String()); !ok {
		conn.logger.Warn("no PTR record or faild to find PTR records of local address %s", conn.remoteAddress.String())
	}

	conn.client.domain = ""
	conn.client.mailFrom = ""
	conn.client.recipients = []string{}
	conn.client.data = nil
	conn.client.forwardStatus = MailForwardIdle

	conn.logger.Info("client %s[%s]:%d connected", conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)

	return false // err!
}

// handle client connection
func (conn *Connection) handle() {

	if config.ConfOpts.MaxClients > 0 && clientCount > config.ConfOpts.MaxClients { // if max clients
		conn.text.busy(conn.localAddress.GetPTR())
		conn.closeForMaxClientsExceeded()
		return
	} else {
		conn.text.greet(conn.localAddress.GetPTR()) // send 220 for conncetion establishment
	}

	for {

		line, err := conn.text.readLine()

		if err == nil {
			cmd, args, err := parseCommand(line)
			if err == CmdParseOk {
				status := conn.handleCommand(cmd, args)
				if status == HandleCommandClose { // if connect close with QUIT...
					break
				}
			} else {
				conn.text.cmdNotRecognized()
			}
			continue
		}

		// if error not nil
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() { // if time out
			conn.text.timeout(conn.localAddress.GetPTR())
			conn.closeWithFailAnd("timeout exceeded")
			break
		} else if err == limitlinereader.ErrTooLongLine {
			conn.text.longLine(conn.localAddress.GetPTR())
			conn.closeWithFailAnd("too long line")
			break
		} else if err == io.ErrUnexpectedEOF { // eof or connection close
			conn.closeWithFail()
			break
		} else { // if unexpected
			conn.closeWithFail()
			break
		}

	}
}

func (conn *Connection) closeWithFail() {
	conn.client.forwardStatus = MailForwardFaild
	conn.close()
	conn.logger.Error("disconnected unfortunately client %s[%s]:%d", conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
}

func (conn *Connection) closeWithSuccess() {
	conn.close()
	conn.logger.Info("disconnected client %s[%s]:%d", conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
}

func (conn *Connection) closeForMaxClientsExceeded() {
	conn.close()
	conn.logger.Warn("disconnected client by server for max clients exceeded %s[%s]:%d", conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
}

func (conn *Connection) closeWithFailAnd(resone string) {
	conn.client.forwardStatus = MailForwardFaild
	conn.close()
	conn.logger.Warn("disconnected client by server for %s %s[%s]:%d", resone, conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
}

func (conn *Connection) reset() {
	conn.forward()

	conn.client.mailFrom = ""
	conn.client.recipients = []string{}
	conn.client.data = nil
	conn.client.forwardStatus = MailForwardIdle
}

func (conn *Connection) close() {
	conn.forward()

	clientCount -= 1
	conn.conn.Close()
}

func (conn *Connection) forward() {
	if conn.client.mailFrom == "" {
		return
	}

	go func(uid string, count int, client Client) {

		email := Email{
			Uid:        fmt.Sprintf("%s_%d", uid, count),
			Domain:     client.domain,
			From:       client.mailFrom,
			Recipients: client.recipients,
			dataByte:   client.data,
			// Data:       string(client.data),
		}

		if client.forwardStatus == MailForwardSuccess {
			email.Success = true
		} else {
			email.Success = false
		}

		mailFwd.ForwardMail(email)
	}(conn.uid, conn.mailCount, conn.client)
}
