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
	clinet.SetFrom("lo🤑pel111@cockatielone.biz")
	clinet.SetRcpt("deg4m4@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=eZaSOTcz4Jd30CNA9/D0uabLRmgjJPgtB3GV/IMOeno=;
 h=from:subject:to:mime-version:content-type;
 b=Vyx+RLqTwm7E8ARSD5avyBRKkooyTfrt+4OLuVvdcJVSgpi5j35Asb8DtIpsN2WiqyJPLRzT5
 xquHowpOMhfJQ9xEDLFXGM3x/mDHYP1Rc8hA/s4RH3Mur4GX5AsPybgXe0lloxOcnK0bm/zbN7v
 3F09t59K2NJGsZF4AmoPMZQtTDc5NkaOBxAv+UWDsuDWU2JsPUUVJMPT4XiNT/kedzUsQfvoobi
 uOBIkAZPTSVq0MAkI88sSfheduk9FjlE4ifVAyAIbF2q+Z4jdPRRYEzpE8/q+zEj7U5LiMVYN4t
 OCpuiNfzNo3d4rnxnwjnT49GXOZKGUTD3GsVuKImKf/Q==
From: lo🤑pel111@cockatielone.biz
To: deg4m4@gmail.com
Subject: Hello Test Message!
Message-ID: <2264435b-4da1-4be5-fba2-36d43bcc86a9@cockatielone.biz>
Date: Wed, 07 Aug 2024 03:09:18 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-5df745e086c847f7-Part_1"

----_NmP-5df745e086c847f7-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

Hello World =F0=9F=A4=91=F0=9F=A4=91
----_NmP-5df745e086c847f7-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-5df745e086c847f7-Part_1--

`
