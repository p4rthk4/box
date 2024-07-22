package main

import smtpclient "github.com/p4rthk4/u2smtp/pkg/client"

func main() {
	clinet := smtpclient.NewClinet()
	clinet.SetHost("lolkongd.com")

	mailContent := []byte{}

	clinet.SendMail(mailContent)
}
