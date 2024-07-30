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
	case "EHLO":
		conn.handleEHello(args)
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
	case "SEND", "SOML", "SAML", "EXPN", "HELP", "TURN", "LHLO", "STARTTLS", "AUTH", "BDAT", "VRFY":
		conn.rw.cmdNotImplemented()
	default:
		conn.rw.cmdNotRecognized()
	}

	return HandleCommandOk
}

func (conn *Connection) handleEHello(args string) {
	if !config.ConfOpts.ESMTP.Enable {
		conn.rw.esmtpDisable()
		return
	}

	domain, err := parseHelloArguments(args)
	if err == HelloArgDomainInvalid {
		conn.rw.reply(501, "domain required for HELO")
		return
	}

	conn.client.domain = domain
	replyMsg := []string{"Hello " + domain}
	replyMsg = append(replyMsg, greetReplyMessage...)
	conn.rw.replyLines(250, replyMsg)
}

func (conn *Connection) handleHello(args string) {
	domain, err := parseHelloArguments(args)
	if err == HelloArgDomainInvalid {
		conn.rw.reply(501, "domain required for HELO")
		return
	}

	conn.client.domain = domain
	conn.rw.reply(250, "%s ready for you", config.ConfOpts.HostName)
}

func (conn *Connection) handleMail(args string) {
	if conn.client.domain == "" {
		conn.rw.reply(503, "Error: send HELO/EHLO first")
		return
	}

	p := newParser(args)
	if ok := p.cutPrefix("FROM:"); !ok {
		conn.rw.syntaxError("Syntax: MAIL FROM:<address>")
		return
	}

	p.trim()
	from, err := p.parseReversePath()
	if err != nil {
		conn.rw.syntaxError("invalid address")
		return
	}

	conn.client.mailFrom = from
	conn.rw.reply(250, "Ok")
}

func (conn *Connection) handleRcpt(args string) {
	if conn.client.mailFrom == "" {
		conn.rw.reply(503, "Error: send MAIL first")
		return
	}

	if len(conn.client.recipients) == config.ConfOpts.MaxRecipients {
		conn.rw.reply(452, "Maximum limit of %v recipients reached", config.ConfOpts.MaxRecipients)
		return
	}

	p := newParser(args)
	if ok := p.cutPrefix("TO:"); !ok {
		conn.rw.syntaxError("Syntax: RCPT TO:<address>")
		return
	}

	p.trim()
	rcpt, err := p.parseReversePath()
	if err != nil {
		conn.rw.syntaxError("invalid address")
		return
	}

	if config.ConfOpts.CheckMailBoxExist {
		if !mailFwd.ExistMailBox(rcpt) {
			conn.rw.reply(550, "mailbox unavailable")
			return
		}
	}

	conn.client.recipients = append(conn.client.recipients, rcpt)
	conn.rw.reply(250, "Ok")
}

func (conn *Connection) handleNoop() {
	conn.rw.reply(250, "Yeop")
}

func (conn *Connection) handleQuit() {
	conn.rw.byyy()
	conn.closeWithSuccess()
}

func (conn *Connection) handleReset() {
	conn.reset()
	conn.rw.reply(250, "Flushed")
}

func (conn *Connection) handleData() {
	if len(conn.client.recipients) == 0 {
		conn.rw.reply(503, "Error: send RCPT first")
		return
	}

	conn.rw.reply(354, "Start mail input end with <CRLF>.<CRLF>")

	data, err := conn.rw.readData()
	if err == io.ErrUnexpectedEOF || err != nil {
		fmt.Println("Lol Error...")
		return
	}
	conn.client.data = data

	conn.rw.reply(250, "Ok")
	conn.client.forwardStatus = MailForwardSuccess
	conn.reset()

	conn.logger.Success("%d email received successfully from %s[%s]:%d", conn.mailCount, conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
	conn.mailCount += 1
}
