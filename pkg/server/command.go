package server

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/p4rthk4/u2smtp/pkg/config"
)

type HandleCommandStatus int

const (
	HandleCommandOk HandleCommandStatus = iota
	HandleCommandClose
)

type BodyType string

const (
	Body7Bit       BodyType = "7BIT"
	Body8BitMIME   BodyType = "8BITMIME"
	BodyBinaryMIME BodyType = "BINARYMIME"
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
	case "BDAT":
		conn.handleBdat(args)
	case "RSET":
		conn.handleReset()
	case "NOOP":
		conn.handleNoop()
	case "QUIT":
		conn.handleQuit()
		return HandleCommandClose
	case "STARTTLS":
		conn.handleStartTls()
	case "SEND", "SOML", "SAML", "EXPN", "HELP", "TURN", "LHLO", "AUTH", "VRFY":
		conn.rw.cmdNotImplemented(cmd)
	default:
		conn.rw.cmdNotRecognized(cmd)
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

	conn.domain = domain
	conn.useEsmtp = true
	replyMsg := []string{"Hello " + domain}
	replyMsg = append(replyMsg, greetReplyMessage...)
	conn.rw.replyLines(250, replyMsg)
}

func (conn *Connection) handleStartTls() {
	if !config.ConfOpts.ESMTP.Enable || !config.ConfOpts.ESMTP.Tls {
		conn.rw.esmtpDisable()
		return
	}

	conn.rw.reply(220, "Ready to start TLS")

	tlsConn := tls.Server(conn.conn, tlsConfig)
	if err := tlsConn.Handshake(); err != nil {
		fmt.Println("TLS Handshake faild", err) // TODO: Remove...
		conn.rw.reply(550, "TLS Handshake error")
		return
	}

	conn.conn = tlsConn
	conn.useTls = true
	conn.rw = newTextReaderWriter(conn.conn)
}

func (conn *Connection) handleHello(args string) {
	domain, err := parseHelloArguments(args)
	if err == HelloArgDomainInvalid {
		conn.rw.reply(501, "domain required for HELO")
		return
	}

	conn.domain = domain
	conn.useEsmtp = false
	conn.rw.reply(250, "%s ready for you", config.ConfOpts.HostName)
}

func (conn *Connection) handleMail(args string) {
	if conn.domain == "" {
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
	conn.mailFrom = from

	if !conn.useEsmtp {
		conn.rw.reply(250, "Ok")
		return
	}

	mailArgs, err := parseArgs(p.s)
	if err != nil {
		conn.rw.reply(501, "Unable to parse MAIL ESMTP parameters")
		return
	}

	for key, value := range mailArgs {
		switch key {
		case "SIZE":
			size, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				conn.rw.reply(501, "Unable to parse SIZE as an integer")
				return
			}
			if config.ConfOpts.ESMTP.MessageSize > 0 && int(size) > config.ConfOpts.ESMTP.MessageSize {
				conn.rw.reply(552, "Max message size exceeded")
				return
			}

			conn.size = int(size)
		case "SMTPUTF8":
			if !config.ConfOpts.ESMTP.Utf8 {
				conn.rw.reply(504, "SMTPUTF8 is not implemented")
				return
			}
			conn.put8 = true
		case "BODY":
			value = strings.ToUpper(value)
			switch BodyType(value) {
			case BodyBinaryMIME:
				if !config.ConfOpts.ESMTP.BinaryMime {
					conn.rw.reply(504, "BINARYMIME is not implemented")
					return
				}
			case Body7Bit, Body8BitMIME:
				// This space is intentionally left blank
			default:
				conn.rw.reply(501, "Unknown BODY value")
				return
			}
			conn.body = BodyType(value)
		case "RET":
			// TODO: // DSN
			// if !c.server.EnableDSN {
			// 	c.writeResponse(504, EnhancedCode{5, 5, 4}, "RET is not implemented")
			// 	return
			// }
			// value = strings.ToUpper(value)
			// switch DSNReturn(value) {
			// case DSNReturnFull, DSNReturnHeaders:
			// 	// This space is intentionally left blank
			// default:
			// 	c.writeResponse(501, EnhancedCode{5, 5, 4}, "Unknown RET value")
			// 	return
			// }
			// opts.Return = DSNReturn(value)
		case "ENVID":
			// TODO: // ENVID
			// if !c.server.EnableDSN {
			// 	c.writeResponse(504, EnhancedCode{5, 5, 4}, "ENVID is not implemented")
			// 	return
			// }
			// value, err := decodeXtext(value)
			// if err != nil || value == "" || !isPrintableASCII(value) {
			// 	c.writeResponse(501, EnhancedCode{5, 5, 4}, "Malformed ENVID parameter value")
			// 	return
			// }
		case "AUTH":
			// value, err := decodeXtext(value)
			// if err != nil || value == "" {
			// 	c.writeResponse(500, EnhancedCode{5, 5, 4}, "Malformed AUTH parameter value")
			// 	return
			// }
			// if value == "<>" {
			// 	value = ""
			// } else {
			// 	p := parser{s: value}
			// 	value, err = p.parseMailbox()
			// 	if err != nil || p.s != "" {
			// 		c.writeResponse(500, EnhancedCode{5, 5, 4}, "Malformed AUTH parameter mailbox")
			// 		return
			// 	}
			// }
		default:
			conn.rw.reply(500, "Unknown MAIL FROM argument")
			return
		}
	}

	conn.rw.reply(250, "Ok")
}

func (conn *Connection) handleRcpt(args string) {
	if conn.mailFrom == "" {
		conn.rw.reply(503, "Error: send MAIL first")
		return
	}

	if len(conn.recipients) == config.ConfOpts.MaxRecipients {
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

	// TODO: parse args

	conn.recipients = append(conn.recipients, rcpt)
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
	if len(conn.recipients) == 0 {
		conn.rw.reply(503, "Error: send RCPT first")
		return
	}

	// TODO: through error when body type is binerymime

	conn.rw.reply(354, "Start mail input end with <CRLF>.<CRLF>")

	data, err := conn.rw.readData()
	if err == io.ErrUnexpectedEOF || err != nil {
		fmt.Println("Lol Error...")
		return
	}
	conn.data = data

	conn.rw.reply(250, "Ok")
	conn.forwardStatus = MailForwardSuccess
	conn.reset()

	conn.logger.Success("%d email received successfully from %s[%s]:%d", conn.mailCount, conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
	conn.mailCount += 1
}

func (conn *Connection) handleBdat(arg string) {
	if len(conn.recipients) == 0 {
		conn.rw.reply(503, "Error: send RCPT first")
		return
	}
	args := strings.Fields(arg)
	if len(args) == 0 {
		conn.rw.reply(501, "Missing chunk size argument")
		return
	}
	if len(args) > 2 {
		conn.rw.reply(501, "Too many arguments")
		return
	}

	last := false
	if len(args) == 2 {
		if !strings.EqualFold(args[1], "LAST") {
			conn.rw.reply(501, "Unknown BDAT argument")
			return
		}
		last = true
	}

	size, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		conn.rw.reply(501, "Malformed size argument")
		return
	}

	if conn.dataBuffer == nil {
		conn.dataBuffer = new(bytes.Buffer)
	}

	if config.ConfOpts.ESMTP.MessageSize > 0 && conn.dataBuffer.Len() > config.ConfOpts.ESMTP.MessageSize {
		conn.rw.reply(552, "Max message size exceeded")
		return
	}

	oldSize := conn.rw.setMaxLineSize(0)
	defer conn.rw.setMaxLineSize(oldSize)

	lr := io.LimitReader(conn.rw.t.R, int64(size))
	n, err := io.Copy(conn.dataBuffer, lr)
	if err != nil || n != int64(size) {
		conn.rw.reply(554, "Error: Transaction failed, data reading error.")
		conn.reset()
		return
	}

	if last {
		conn.data = conn.dataBuffer.Bytes()
		conn.rw.reply(250, "Ok, last %d bytes received, total %d", size, conn.dataBuffer.Len())

		conn.forwardStatus = MailForwardSuccess
		conn.reset()

		conn.logger.Success("%d email received successfully from %s[%s]:%d", conn.mailCount, conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String(), conn.remoteAddress.port)
		conn.mailCount += 1
	} else {
		conn.rw.reply(250, "%d bytes received, total %d", size, conn.dataBuffer.Len())
	}
}
