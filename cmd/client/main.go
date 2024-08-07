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
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From Parth!
Message-ID: <b1f7e0c1-331f-eca6-a633-d724e596da43@cockatielone.biz>
Date: Wed, 07 Aug 2024 04:42:33 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative; boundary="000000000000ddaf2c061f104677"

--000000000000ddaf2c061f104677
Content-Type: text/plain; charset="UTF-8"

I Am *Parth* Degama

--000000000000ddaf2c061f104677
Content-Type: text/html; charset="UTF-8"

<meta http-equiv="Content-Type" content="text/html; charset=utf-8"><div dir="ltr"><div class="gmail_default" style="font-family:arial,sans-serif">I Am <b>Parth</b> Degama</div></div>

--000000000000ddaf2c061f104677--
`