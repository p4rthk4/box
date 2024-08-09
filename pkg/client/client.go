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
	StartTls     bool
	CheckTlsHost bool

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

	Response ClientResponse

	chunkSize int
}

func NewClinet() SMTPClinet {
	return SMTPClinet{
		StartTls:     true,
		CheckTlsHost: false,
		hostname:     "localhost",

		RcptHost: "",
		RcptPort: 25,

		From: "",
		Rcpt: "",

		timeout: time.Minute * 2,

		chunkSize: 1024 * 2,
	}
}

// set rcpt port host use for make connection
func (client *SMTPClinet) SetRcptHost(host string) {
	client.RcptHost = host
}

// set rcpt port no use for make connection
func (client *SMTPClinet) SetRcptPort(port int) {
	client.RcptPort = port
}

// set client host name use in hello
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
	ServerError bool
}

type ClientResponse struct {
	Errors         []ClientServerError
	Success        bool
	TempError      bool // temp error (4yz)
	AnyClientError bool // any client unknown error
	Status         string
}

func (client *SMTPClinet) SendMail() {
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
			return
		}

		mxs, err := net.LookupMX(domainName)
		if err != nil {
			client.Response.Errors = append(client.Response.Errors, ClientServerError{
				domainName:  domainName,
				errorString: fmt.Sprintf("no any MX records found of %s domain", domainName),
			})
			return
		}
		mxRecords = mxs
	} else {
		client.Response.Errors = append(client.Response.Errors, ClientServerError{
			domainName:  "unknown",
			errorString: "no any host and rcpt found",
		})
		return
	}

	if IsSMTPUTF8(client.From) {
		client.UTF8 = true
	}
	if IsSMTPUTF8(client.Rcpt) {
		client.UTF8 = true
	}

	for _, m := range mxRecords {
		if client.Response.Success {
			break
		}

		ips, err := getIPFromString(m.Host)
		if err != nil {
			client.Response.Errors = append(client.Response.Errors, ClientServerError{
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

			// TODO: log
			fmt.Printf("send mail to: %s %s:%d\n", m.Host, rcptIP, client.RcptPort)

			client.RcptHost = m.Host
			conn, err := client.createNewConn(address)
			if err != nil {
				// fail log
				client.Response.Errors = append(client.Response.Errors, ClientServerError{
					domainName:  m.Host,
					errorString: err.Error(),
				})
				client.Response.TempError = true
				continue
			}

			err = conn.handleConn()
			if err != nil {
				switch e := err.(type) {
				case SMTPServerError:
					if e.GetErrorType() == SMTPErrorTemp {
						client.Response.TempError = true
					}
					client.Response.Errors = append(client.Response.Errors, ClientServerError{
						domainName:  m.Host,
						errorString: err.Error(),
					})
					continue
				case net.Error:
					if e.Timeout() {
						client.Response.Errors = append(client.Response.Errors, ClientServerError{
							domainName:  m.Host,
							errorString: fmt.Sprintf("connection timeout with %s by server", address),
						})
						client.Response.TempError = true
						continue
					}
				default:
					client.Response.Errors = append(client.Response.Errors, ClientServerError{
						domainName:  m.Host,
						errorString: err.Error(),
						ServerError: true,
					})
					client.Response.AnyClientError = true
					continue
				}
			}

			client.Response.Success = true
			break
		}
	}
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


func (client *SMTPClinet) GetResponse() ClientResponse {
	if client.Response.Success {
		client.Response.Status = "SUCCESS"
		return client.Response
	}

	if client.Response.TempError {
		client.Response.Status = "TRYAGAIN"
	} else {
		client.Response.Status = "FAIL"
	}
	
	return client.Response
}
