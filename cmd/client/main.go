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
	clinet.SetRcpt("parthka.2005@proton.me")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=s+LisPZ/kV8OBCx+/82+4XwlUr7IYG1h+a1OXFmrcUQ=;
 h=from:subject:to:mime-version:content-type;
 b=R8KZBpBOL5mYiuiKOFtF6dpOrqKY8NXg144Wi/kMAXXJElx+Ot0nBJnuNcd2OfGg40Yy14wRk
 KEYnBWPh97z0gnDmOi4vy1ijS4neJ7XQW9Op0DqY59IvgkO/4qZ5RMxLbCo1gsk2RPU8ihmsFOz
 VI3lWWA8Ile973W6SCU8kV24tKJTQxy9o1p5eNHgyGe4+MIMjaafDUGili3WgkG1TZfdos+19EE
 O0ptnC53xSMuidECfWkvFWAZbrD4VcZGIcipIIX6nyyif3R5nMfsHAzqpSJ9U7MqeaBT4atLZID
 LZoq+koklgA+eCWkna0G9AjuyA7jHoWZfDWwwp94tGBw==
From: lo🤑pel111@cockatielone.biz
To: parthka.2005@proton.me
Subject: Hello Test Message!
Message-ID: <cb4db153-9dc9-65e9-492a-aabee9d7b730@cockatielone.biz>
Date: Wed, 07 Aug 2024 03:20:17 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-7477e98eb12c786a-Part_1"

----_NmP-7477e98eb12c786a-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

Hello World =F0=9F=A4=91=F0=9F=A4=91
----_NmP-7477e98eb12c786a-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-7477e98eb12c786a-Part_1--

`
