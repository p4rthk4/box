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
	clinet.SetRcpt("pthreeh@outlook.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=cwVlsOsYr8jpWoqSEZl2vgUNnNMbmD+g62zYpDEYxl8=;
 h=from:subject:to:mime-version:content-type;
 b=KbKQy4rPpiJnyRknxp9j2bASfB+89wu93G5lsP6tIVO2ZHIdSCYVii8KqgaDkW0UjlhaZ5X9L
 CiHTOaqsAMNzj/mG9MJ0dJ23K7wISQU3XPpXevjZ1uKfgz36mnBtZTGT1fTDQzDXVQIP3pKO+Xu
 hX6jSKN++15s5rFokUxr2DC3GhMPUobpPnsWEEfcxSY+mu2leERjInqPyI59+ffJmW2NQ/Tdghv
 xkbrfaR13ZGcDx4zspr77FiQggmW6YOTrtjbpRihlPfFviGfSJPxMB07GtOeEBj02r8MRelDNDj
 rTK128YG5L3bhUf0qc2zwaRO797PlL5rfbyc0L9zqYaA==
From: lo🤑pel111@cockatielone.biz
To: pthreeh@outlook.com
Subject: Hello Test Message!
Message-ID: <437538c5-ca70-3c6b-85fe-d24002c92f97@cockatielone.biz>
Date: Wed, 07 Aug 2024 03:22:52 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-1badf63e3634940f-Part_1"

----_NmP-1badf63e3634940f-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

Hello World =F0=9F=A4=91=F0=9F=A4=91
----_NmP-1badf63e3634940f-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-1badf63e3634940f-Part_1--

`
