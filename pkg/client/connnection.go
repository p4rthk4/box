package client

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type ClientConn struct {
	conn net.Conn
	rw   *TextReaderWriter

	helloDone bool
	extension map[string]string

	smtpClient *SMTPClinet
}

type ClientError struct {
	tryNext bool
	faild   bool
	wait    bool

	err  string
	code int
}

func (clientErr ClientError) Error() string {
	err := ""
	if clientErr.tryNext {
		err = "try next connection"
	}
	if clientErr.faild {
		err = "fail to send mail"
	}
	if clientErr.wait {
		err = "wait 1 * (n) minute and try again"
	}

	return fmt.Sprintf("%s\nserver reply with %03d\n%s", err, clientErr.code, clientErr.err)
}

func (conn *ClientConn) handleConn() error {
	defer conn.close()

	err := conn.hello()
	if err != nil {
		return serverErrToClientErr(err)
	}

	fmt.Println(conn.extension)

	err = conn.mail()
	if err != nil {
		return serverErrToClientErr(err)
	}

	err = conn.rcpt()
	if err != nil {
		return serverErrToClientErr(err)
	}

	return nil
}

func (conn *ClientConn) hello() error {
	if conn.helloDone {
		return nil
	}

	if err := conn.greet(); err != nil {
		return err
	}

	err := conn.ehlo()
	if err != nil {
		var smtpServerError SMTPServerError
		if errors.As(err, &smtpServerError) && (smtpServerError.Code == 500 || smtpServerError.Code == 502) {
			err = conn.helo()
		}
	}
	return err
}

func (conn *ClientConn) greet() error {
	if conn.helloDone {
		return nil
	}
	_, _, err := conn.rw.readResponse(220)
	return err
}

func (conn *ClientConn) ehlo() error {
	_, msg, err := conn.rw.cmd(250, "EHLO %s", conn.smtpClient.hostname)
	if err != nil {
		return err
	}
	conn.helloDone = true

	ext := make(map[string]string)
	extList := strings.Split(msg, "\n")
	if len(extList) > 1 {
		extList = extList[1:]
		for _, line := range extList {
			args := strings.SplitN(line, " ", 2)
			if len(args) > 1 {
				ext[args[0]] = args[1]
			} else {
				ext[args[0]] = ""
			}
		}
	}
	conn.extension = ext
	return err
}

func (conn *ClientConn) helo() error {
	_, _, err := conn.rw.cmd(250, "HELO %s", conn.smtpClient.hostname)
	if err == nil {
		conn.helloDone = true
	}
	return err
}

func (conn *ClientConn) mail() error {
	var sb strings.Builder
	sb.Grow(2048)

	fmt.Fprintf(&sb, "MAIL FROM:<%s>", conn.smtpClient.From)
	if _, ok := conn.extension["8BITMIME"]; ok {
		sb.WriteString(" BODY=8BITMIME")
	}
	if _, ok := conn.extension["SIZE"]; ok && conn.smtpClient.Size != 0 {
		fmt.Fprintf(&sb, " SIZE=%v", conn.smtpClient.Size)
	}

	if conn.smtpClient.RequireTLS {
		if _, ok := conn.extension["REQUIRETLS"]; ok {
			sb.WriteString(" REQUIRETLS")
		} else {
			return errors.New("smtp: server does not support REQUIRETLS")
		}
	}

	// if domain and email address is i18n so enable this
	if conn.smtpClient.UTF8 {
		if _, ok := conn.extension["SMTPUTF8"]; ok {
			sb.WriteString(" SMTPUTF8")
		} else {
			return errors.New("smtp: server does not support SMTPUTF8")
		}
	}

	if _, ok := conn.extension["DSN"]; ok {
		switch conn.smtpClient.DSNReturn {
		case DSNReturnFull, DSNReturnHeaders:
			fmt.Fprintf(&sb, " RET=%s", string(conn.smtpClient.DSNReturn))
		case "":
			// This space is intentionally left blank
		default:
			return errors.New("smtp: Unknown RET parameter value")
		}
		// TODO: idk
		// if opts.EnvelopeID != "" {
		// 	if !isPrintableASCII(opts.EnvelopeID) {
		// 		return errors.New("smtp: Malformed ENVID parameter value")
		// 	}
		// 	fmt.Fprintf(&sb, " ENVID=%s", encodeXtext(opts.EnvelopeID))
		// }
	}

	fmt.Println(sb.String())

	_, _, err := conn.rw.cmd(250, "%s", sb.String())
	return err
}

func (conn *ClientConn) rcpt() error {
	var sb strings.Builder
	sb.Grow(2048)

	fmt.Fprintf(&sb, "RCPT TO:<%s>", conn.smtpClient.Rcpt)
	if _, ok := conn.extension["DSN"]; ok {
		// TODO: dsn ...
		_ = ok
	}

	if _, _, err := conn.rw.cmd(25, "%s", sb.String()); err != nil {
		return err
	}

	return nil
}

func serverErrToClientErr(err error) error {
	if err != nil {
		switch e := err.(type) {
		case SMTPServerError:
			// TODO: error handling
			fmt.Println("Got it, this is error...")
			return ClientError{
				tryNext: true,
				err:     e.Message,
				code:    e.Code,
			}
		}
		return err
	}
	return err
}

func (conn *ClientConn) close() {
	err := conn.conn.Close()
	if err != nil {
		// TODO: logs
		fmt.Println("connecrion close error...")
	}
}
