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
	q=dns/txt; s=draft; bh=PKtoTjqHuudr64nYpizw9stW9CenBfl0sXHFI8KEfBA=;
	h=from:subject:to:mime-version:content-type:content-transfer-encoding;
	b=U4qYOMP8InhfY2pXhqVW8HlNS+xukXqgk+qL7N5rpmuuc9H2GAQBtvTJzvnGTdHu+V5MidvMB
	 vdhM9REw7Q9QYeMBDjvu4AB8qDoBc0tTTxpZkLB5Xr18xegu+ankojTppFfwlR1K/uuQcnk8ya9
	 7eQlshyXJ/U9K7vaF3u5Zlz2M84qg+ZlhBuFI8srVMPvNpzFvYM2DQhr2sukfBToXsdVxwCpmkt
	 xo3bTBeCwPStBQXeaQCxn9+mmEj3CwFG3Mwi4BI5A1+whvchYOCjVaewasqpyS5u2Rd44aU8Ejg
	 7x650TAbha6gStNmwMePlfICkI8OemV5FPtEQGCMa37A==
From: lopel@cockatielone.biz
To: pthreeh@outlook.com
Subject: Hello Test Message From Parth!
Message-ID: <2fdf2cf5-4117-7ce7-3cdc-44bd4c8f770e@cockatielone.biz>
Content-Transfer-Encoding: 7bit
Date: Wed, 07 Aug 2024 03:32:35 +0000
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8

Hello World!, my name is parth degama and your name
`