package main

import (
	"fmt"
	"time"

	smtpclient "github.com/p4rthk4/u2smtp/pkg/client"
)

func main() {	

	clinet := smtpclient.NewClinet()
	// try cockatielone.biz
	// clinet.SetHost("alt3.gmail-smtp-in.l.google.com")
	clinet.SetTimeout(5 * time.Second)
	clinet.SetFrom("degama@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetRcpt("sahello@parthka.dev")
	clinet.SetRcpt("parthka.2005@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetHostname("home.lope.dev")

	clinet.Size = 1024
	clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `Content-Type: multipart/mixed;
 boundary="----sinikael-?=_1-17221480878150.018640130471523797"
From: degama@cockatielone.biz
To: parthka.2005@gmail.com
Subject: Hello Lope =?UTF-8?B?8J+Siw==?=
Date: Sun, 28 Jul 2024 06:28:07 +0000
Message-Id: <1722148087817-999db44e-b6240e4d-3497e6f5@cockatielone.biz>
MIME-Version: 1.0

------sinikael-?=_1-17221480878150.018640130471523797
Content-Type: text/plain
Content-Transfer-Encoding: base64

VGhpcyBpcyB0aGUgcGxhaW4gdGV4dCB2ZXJzaW9uICBvZiB0aGUgZW1haWwuIA==
------sinikael-?=_1-17221480878150.018640130471523797--
`
