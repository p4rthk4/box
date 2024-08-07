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
	clinet.SetFrom("lopel111@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=ZaFb4HjAGQRw753RecXgkuWa1t0gTDMY+HrG2YvFZ70=;
 h=from:subject:to:mime-version:content-type;
 b=ig2r9FPSBajz9iTt/rwWByInmAyzd7rrqohfvUB1N1VkQscwJ3d/vFjzhKLpJiII9arXhxUDm
 y/nlkHjqwVHd5m5zp6ljCPWtZ9dYNBVJW5dGQwlBrDnsjCBQkRuDd1dOWOi1Zbc0KVyc0IyMb72
 nB24AU47HSMOi2r7W+wffSf7vTtB+jBjo3uu1eo339J6U5qLx6rCNgWk0Fg9e1d0RovhsA5Vb/y
 kWOsgaVg2pFsgKyitRoSA9NawzTc72Pc0YORrcjxWvEaxM2TrU76jTQpTVKlSMGsGpIPnWyIWOp
 eZey3DFkE1ODba4ak+DrApMCOG4LWcfT1vWmrMs+jGGg==
From: lopel111@cockatielone.biz
To: parthka.2005@gmail.com
Subject: Hello Test Message!
Message-ID: <c2387efe-6915-d0ff-b9ee-fdaea1c125e8@cockatielone.biz>
Date: Wed, 07 Aug 2024 03:18:46 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-c19e3b851409c4e5-Part_1"

----_NmP-c19e3b851409c4e5-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

Hello World =F0=9F=A4=91=F0=9F=A4=91
----_NmP-c19e3b851409c4e5-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-c19e3b851409c4e5-Part_1--

`
