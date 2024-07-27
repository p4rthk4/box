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
	clinet.SetFrom("degama@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetRcpt("parthka.2005@cockatielone.biz")
	clinet.SetRcpt("hello@parthka.dev")
	clinet.SetHostname("home.parthka.dev")
	clinet.Size = 1024
	clinet.DSNReturn = smtpclient.DSNReturnFull
	clinet.SetData(mailContent)
	err := clinet.SendMail()
	fmt.Println(err)
}
	