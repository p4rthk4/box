package client

import (
	"fmt"
	"net"
	"os"
	"time"
)

type DSNReturnType string

const (
	DSNReturnFull    DSNReturnType = "FULL"
	DSNReturnHeaders DSNReturnType = "HDRS"
)

type SMTPClinet struct {
	StartTLS bool
	hostname string

	RcptHost string
	RcptPort int

	From string
	Rcpt string

	data []byte

	timeout time.Duration

	Size int
	UTF8 bool
	// RequireTLS bool
	// DSNReturn  DSNReturnType

	chunkSize int
}

func NewClinet() SMTPClinet {
	return SMTPClinet{
		StartTLS: true,
		hostname: "localhost",

		RcptHost: "",
		RcptPort: 25,

		From: "",
		Rcpt: "",

		timeout: time.Minute * 2,

		chunkSize: 1024 * 2,
	}
}

// set rcpt port host
func (client *SMTPClinet) SetRcptHost(host string) {
	client.RcptHost = host
}

// set rcpt port no
func (client *SMTPClinet) SetRcptPort(port int) {
	client.RcptPort = port
}

// set client host name
func (client *SMTPClinet) SetHostname(host string) {
	client.hostname = host
}

func (client *SMTPClinet) SetFrom(from string) {
	client.From = from
}

func (client *SMTPClinet) SetRcpt(rcpt string) {
	client.Rcpt = rcpt
}

func (client *SMTPClinet) SetData(data []byte) {
	client.data = data
}

func (client *SMTPClinet) SetTimeout(t time.Duration) {
	client.timeout = t
}

type ClientServerError struct {
	domainName  string
	errorString string
}

func (client *SMTPClinet) SendMail() error {

	client.Size = len(client.data) + 5 // add 5 for <crlf>.<crlf>

	mxRecords := []*net.MX{}
	if client.RcptHost != "" {
		mxRecords = append(mxRecords, &net.MX{
			Host: client.RcptHost,
			Pref: 0,
		})
	} else if client.Rcpt != "" {
		domainName, err := getDomainFromEmail(client.Rcpt)
		if err != nil {
			return err
		}

		mxs, err := net.LookupMX(domainName)
		if err != nil {
			return fmt.Errorf("no any MX records found of %s domain", domainName)
		}
		mxRecords = mxs
	} else {
		return fmt.Errorf("no any host and rcpt found")
	}

	if IsSMTPUTF8(client.From) {
		client.UTF8 = true
	}
	if IsSMTPUTF8(client.Rcpt) {
		client.UTF8 = true
	}

	// errors
	allErrors := []ClientServerError{}

	// loop of sand mail to mx servers
	for _, m := range mxRecords {
		ips, err := getIPFromString(m.Host)
		if err != nil {
			allErrors = append(allErrors, ClientServerError{
				domainName:  m.Host,
				errorString: err.Error(),
			})
			continue
		}

		for _, rcptIP := range ips {
			address := ""
			if isIPv6(rcptIP) {
				address = fmt.Sprintf("[%s]:%d", rcptIP, client.RcptPort)
			} else {
				address = fmt.Sprintf("%s:%d", rcptIP, client.RcptPort)
			}

			client.RcptHost = m.Host
			conn, err := client.createNewConn(address)
			if err != nil {
				allErrors = append(allErrors, ClientServerError{
					domainName:  m.Host,
					errorString: err.Error(),
				})
				continue
			} else {
				fmt.Println("connect", m.Host)
			}

			err = conn.handleConn()
			if err != nil {
			retryErr:
				switch e := err.(type) {
				case ClientError:
					allErrors = append(allErrors, ClientServerError{
						domainName:  m.Host,
						errorString: err.Error(),
					})
					continue
				case net.Error:
					if e.Timeout() {
						fmt.Println("time...")
						allErrors = append(allErrors, ClientServerError{
							domainName:  m.Host,
							errorString: fmt.Sprintf("connection timeout with %s by server", address),
						})
						continue
					}
				case SMTPServerError:
					fmt.Println("Error trnfer in thire area")
					err = serverErrToClientErr(err)
					goto retryErr
				default:
					allErrors = append(allErrors, ClientServerError{
						domainName:  m.Host,
						errorString: err.Error(),
					})
					continue
				}

			}

			break // TODO: tmp break for nothing...!
		}
		fmt.Println(ips)
	}

	fmt.Println(allErrors, len(allErrors))
	return nil
}

func (client *SMTPClinet) createNewConn(address string) (ClientConn, error) {
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
