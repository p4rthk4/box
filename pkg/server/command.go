// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"fmt"
	"io"

	"github.com/p4rthk4/u2smtp/pkg/config"
)

type HandleCommandStatus int

const (
	HandleCommandOk HandleCommandStatus = iota
	HandleCommandClose
)

func (conn *Connection) handleCommand(cmd string, args string) HandleCommandStatus {
	switch cmd {
	case "HELO":
		conn.handleHello(args)
	case "MAIL":
		conn.handleMail(args)
	case "RCPT":
		conn.handleRcpt(args)
	case "DATA":
		conn.handleData()
	case "RSET":
		conn.handleReset()
	case "NOOP":
		conn.handleNoop()
	case "QUIT":
		conn.handleQuit()
		return HandleCommandClose
	case "SEND", "SOML", "SAML", "EXPN", "HELP", "TURN", "EHLO", "LHLO", "STARTTLS", "AUTH", "BDAT", "VRFY":
		conn.text.cmdNotImplemented()
	default:
		conn.text.cmdNotRecognized()
	}

	return HandleCommandOk
}

func (conn *Connection) handleHello(args string) {
	domain, err := parseHelloArguments(args)
	if err == HelloArgDomainInvalid {
		conn.text.reply(501, "domain required for HELO")
		return
	}

	conn.client.domain = domain

	conn.text.reply(250, "%s ready for you", config.ConfOpts.HostName)

}

func (conn *Connection) handleMail(args string) {
	if conn.client.domain == "" {
		conn.text.reply(503, "Error: send HELO/EHLO first")
		return
	}

	p := newParser(args)

	if ok := p.cutPrefix("FROM:"); !ok {
		conn.text.syntaxError("Syntax: MAIL FROM:<address>")
		return
	}

	p.trim()
	from, err := p.parseReversePath()
	if err != nil {
		conn.text.syntaxError("invalid address")
		return
	}

	conn.client.mailFrom = from

	conn.text.reply(250, "Ok")
}

func (conn *Connection) handleRcpt(args string) {
	if conn.client.mailFrom == "" {
		conn.text.reply(503, "Error: send MAIL first")
		return
	}

	if len(conn.client.recipients) == config.ConfOpts.MaxRecipients {
		conn.text.reply(452, "Maximum limit of %v recipients reached", config.ConfOpts.MaxRecipients)
		return
	}

	p := newParser(args)

	if ok := p.cutPrefix("TO:"); !ok {
		conn.text.syntaxError("Syntax: RCPT TO:<address>")
		return
	}

	p.trim()
	rcpt, err := p.parseReversePath()
	if err != nil {
		conn.text.syntaxError("invalid address")
		return
	}

	conn.client.recipients = append(conn.client.recipients, rcpt)

	conn.text.reply(250, "Ok")
}

func (conn *Connection) handleNoop() {
	conn.text.reply(250, "Yeep")
}

func (conn *Connection) handleQuit() {
	conn.text.byyy()
	conn.closeWithSuccess()
}

func (conn *Connection) handleReset() {
	conn.reset()
	conn.text.reply(250, "Flushed")
}

func (conn *Connection) handleData() {
	if len(conn.client.recipients) == 0 {
		conn.text.reply(503, "Error: send RCPT first")
		return
	}

	conn.text.reply(354, "Start mail input end with <CRLF>.<CRLF>")

	data, err := conn.text.readData()
	if err == io.ErrUnexpectedEOF || err != nil {
		fmt.Println("Lol Error...")
		return
	}
	conn.client.data = data

	conn.text.reply(250, "Ok")
	conn.client.forwardStatus = MailForwardSuccess
	conn.reset()

	conn.logger.Success("%d email received successfully from %s[%s]:%d", conn.mailCount, conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)

	conn.mailCount += 1
}
