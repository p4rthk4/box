package client

import (
	"fmt"
	"net"
	"time"
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

	// errors
	allClientServerError := []ClientServerError{}

	// loop of sand mail to mx servers
	for _, m := range mxRecords {
		ips, err := getIPFromString(m.Host)
		if err != nil {
			allClientServerError = append(allClientServerError, ClientServerError{
				domainName:  m.Host,
				errorString: err.Error(),
			})
			continue
		}

		for _, v := range ips {
			conn, err := client.createNewConn(v, client.RcptPort)
			if err != nil {
				allClientServerError = append(allClientServerError, ClientServerError{
					domainName:  m.Host,
					errorString: err.Error(),
				})
				continue
			} else {
				fmt.Println("connect", m.Host)
			}

			err = conn.handleConn()
			if err != nil{
				allClientServerError = append(allClientServerError, ClientServerError{
					domainName:  m.Host,
					errorString: err.Error(),
				})
				continue
			}

		}

		fmt.Println(ips)

	}

	fmt.Println(allClientServerError)

	return nil
}
