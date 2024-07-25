package client

import (
	"fmt"
	"net"
	"os"
)

type ClientConn struct {
	conn net.Conn
	rw   *TextReaderWriter

	helloDone bool

	smtpClient *SMTPClinet
}

func (client *SMTPClinet) createNewConn(ip net.IP, port int) (ClientConn, error) {

	address := ""
	if isIPv6(ip) {
		address = fmt.Sprintf("[%s]:%d", ip.String(), port)
	} else {
		address = fmt.Sprintf("%s:%d", ip.String(), port)
	}

	conn, err := net.DialTimeout("tcp", address, client.timeout)
	if err != nil {
		switch e := err.(type) {
		case *net.OpError:
			switch e.Op {
			case "dial":
				switch e := e.Err.(type) {
				case *os.SyscallError:
					if e.Syscall == "connect" {
						return ClientConn{}, fmt.Errorf("connection refused when server connect with %s", address)
					}
				case net.Error:
					if e.Timeout() {
						return ClientConn{}, fmt.Errorf("connection timeout with %s by server", address)
					}
				}
			}
		case net.Error:
			if e.Timeout() {
				return ClientConn{}, fmt.Errorf("connection timeout with %s by server", address)
			}
		}

		return ClientConn{}, fmt.Errorf("internal server error for connect with %s	%v", address, err)
	}

	return ClientConn{
		conn: conn,
		rw:   newTextReaderWriter(conn),

		smtpClient: client,
	}, nil
}

func (conn *ClientConn) handleConn() error {

	err := conn.hello()
	if err != nil {
		return err
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

	conn.helloDone = true
	// if err := conn.ehlo(); err != nil {
	// 	var smtpServerError *SMTPServerError
	// 	if errors.As(err, &smtpServerError) && (smtpServerError.Code == 500 || smtpServerError.Code == 502) {
	// 		// The server doesn't support EHLO, fallback to HELO
	// 		conn.helloDone = conn.helo()
	// 	} else {
	// 		conn.helloDone = err
	// 	}
	// }
	return nil
}

func (conn *ClientConn) greet() error {
	if conn.helloDone {
		return nil
	}

	conn.helloDone = true
	code, message, err := conn.rw.readResponse(220)
	if err != nil {
		return err
	}
	fmt.Println("response:", code, message)

	return nil
}

func (conn *ClientConn) close() {
	conn.conn.Close()
}
