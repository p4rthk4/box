package server

import (
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
	case "RSET":
		conn.handleReset()
	case "NOOP":
		conn.handleNoop()
	case "QUIT":
		conn.handleQuit()
		return HandleCommandClose
	case "SEND", "SOML", "SAML", "EXPN", "HELP", "TURN", "LHLO", "STARTTLS", "AUTH", "BDAT", "VRFY":
		conn.rw.cmdNotImplemented(cmd)
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
	conn.useEsmtp = true
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
	conn.useEsmtp = false
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

	if !conn.useEsmtp {
		conn.rw.reply(250, "Ok")
		return
	}

	fmt.Println("args", args)

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
		case "REQUIRETLS":
			if !config.ConfOpts.ESMTP.Tls {
				conn.rw.reply(504, "STARTTLS and REQUIRETLS is not implemented")
				return
			}
			conn.requireTls = true
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

	// TODO: parse args

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
