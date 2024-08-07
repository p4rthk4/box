package main

import (
	"fmt"
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

	clinet.SetData([]byte(mail))
	fmt.Println([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From Parth!
Message-ID: <7be02c46-c287-526f-8308-68016d92707b@cockatielone.biz>
Date: Wed, 07 Aug 2024 04:56:37 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-179eaccf9a33fea7-Part_1"

----_NmP-179eaccf9a33fea7-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-179eaccf9a33fea7-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGI+SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1lPC9i
Pg==
----_NmP-179eaccf9a33fea7-Part_1--

`