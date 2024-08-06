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
 q=dns/txt; s=draft; bh=UN1Luk0CfA92T8/KKG0F5isLo3kzJdhmFjMztTyXB40=;
 h=from:subject:to:mime-version:content-type;
 b=kDH+tTQLJWmFgs3DLszQ2u8kmxzay9A0I9xWIV3oQn+hA83Xr0FgOAWYPlBqcr2RwBvWOcRjb
 5fIpuF8Q1eWwDAteQWrewSAKzrvHlFLaBlTUoPPT4UrC034+bZwkAE8JFnXlXi02zNxE+OBGYDn
 sV/G/LC340A5CpegERri40TEk1RA6aY8LxkRJ5l/kODrP8X5c/sLM9kOukAxEO9Oxfq3opPbAQ+
 ALXg7KOrcffriHci8bLM7taPISRvGDpUI1Vn8U7dmHydVWU9pwnghxhOaB2V7XDECuTJUqvzGDH
 +FLx6v+qsZYHbj7i3IxFWywQlFUOhsKL0/fykZOK2+vw==
From: parthka@myworkspacel.ink
To: Parth <parthka.2005@gmail.com>
Subject: Test Message!
Message-ID: <32ff42dc-d901-16e0-6c5a-2c1b5baf1843@myworkspacel.ink>
Date: Tue, 06 Aug 2024 16:17:03 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-5e2249bfe0044c95-Part_1"

----_NmP-5e2249bfe0044c95-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-5e2249bfe0044c95-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-5e2249bfe0044c95-Part_1--
`
