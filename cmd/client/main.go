package main

import (
	"fmt"
	"time"

	smtpclient "github.com/p4rthk4/u2smtp/pkg/client"
)

func main() {
	mailContent := []byte{}

	clinet := smtpclient.NewClinet()
	// try cockatielone.biz
	// clinet.SetHost("alt3.gmail-smtp-in.l.google.com")
	clinet.SetTimeout(5 * time.Second)
	clinet.SetRcpt("hello@parthka.dev")
	clinet.SetRcpt("parthka.2005@cockatielone.biz")
	clinet.SetData(mailContent)
	err := clinet.SendMail()
	fmt.Println(err)
}
