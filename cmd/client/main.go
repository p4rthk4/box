package main

import (
	"fmt"
	"strings"
	"time"

	smtpclient "github.com/p4rthk4/u2smtp/pkg/client"
	"github.com/p4rthk4/u2smtp/pkg/config"
)

func main() {

	config.LoadConfig() // load conifg in config file

	clinet := smtpclient.NewClinet()
	// try cockatielone.biz
	// clinet.SetHost("alt3.gmail-smtp-in.l.google.com")
	clinet.SetTimeout(5 * time.Second)
	clinet.SetFrom("parthdegama@cockatielone.biz")
	clinet.SetRcpt("pthreeh@outlook.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	mail = strings.ReplaceAll(mail, "\n", "\r\n")
	clinet.SetData([]byte(mail))
	fmt.Println([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From Parth!
Message-ID: <29b208f0-89f4-9150-2ce5-723d5f948c90@cockatielone.biz>
Date: Wed, 07 Aug 2024 05:00:17 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-537d6703f74f3988-Part_1"

----_NmP-537d6703f74f3988-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-537d6703f74f3988-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGI+SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1lPC9i
Pg==
----_NmP-537d6703f74f3988-Part_1--

`