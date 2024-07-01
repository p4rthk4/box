// U2SMTP - smpt client
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package client

import (
	"fmt"
	"net"
)

type SMTPClinet struct {
	StartTLS bool
	Host     string
	Port     int
	From     string
	Rcpt     string

	hostAddress ServerAddress
	mxRecords []string
	mail      []byte
}

func NewClinet() SMTPClinet {
	return SMTPClinet{
		Host:     "",
		Port:     25,
		StartTLS: true,
		From:     "",
		Rcpt:     "",
	}
}

func (client *SMTPClinet) SetHost(host string) {
	client.Host = host
}

func (client *SMTPClinet) SetPort(port int) {
	client.Port = port
}

func (client *SMTPClinet) SetFrom(from string) {
	client.From = from
}

func (client *SMTPClinet) SetRcpt(rcpt string) {
	client.Rcpt = rcpt
}

func (clinet *SMTPClinet) SendMail(mail []byte) error {
	clinet.mail = mail
	fmt.Println("Send Mail...")
	if (clinet.Host != "") {
		ipAddr := net.ParseIP(clinet.Host)
		if ipAddr == nil {

		}
	}

	return nil
}
