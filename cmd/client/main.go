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
	clinet.SetFrom("parthka@myworkspacel.ink")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=myworkspacel.ink;
	q=dns/txt; s=draft; bh=ddgIlflXn5KXzZYAkEbeLhBKFYGaZHohlzlMrIoaCzo=;
	h=from:subject:to:mime-version:content-type;
	b=kV9LYU1Vytq5UcquD8ZKCUGBVmqy3nYE0qfzIlqP9og5oYd2JSa89+eB9xE6lztODW6jtodDp
	 cilS9tcA9ko1uKLFWZ0Gj1FSEtitNLQY6/Y48TDeQXDTmuwOaf4YkRpIPIGRSKpQWWNWPYhglaD
	 HpOYLt/JvZNhfmwSORPbE6s=
From: parthka@myworkspacel.ink
To: Parth <parthka.2005@gmail.com>
Subject: Test Message!
Message-ID: <7dc4a9e5-5e35-0e02-6edd-1b53e4c46e5d@myworkspacel.ink>
Date: Tue, 06 Aug 2024 16:00:25 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-a6e2095a9a0e0ea4-Part_1"

----_NmP-a6e2095a9a0e0ea4-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-a6e2095a9a0e0ea4-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-a6e2095a9a0e0ea4-Part_1--


`
