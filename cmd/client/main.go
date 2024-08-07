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
 q=dns/txt; s=draft; bh=PbL7LhUCQqop8uUlTECK0E4k2GVsFBaEhVQXUCoDvMo=;
 h=from:subject:to:mime-version:content-type;
 b=QkLUwguLM4vrvbjjOrdwnSGlGDQ8gEf0hqn3V37XF3UoLtSQzIcHGL8Q6+YzcDAOlU9wGw3+U
 Ca0l5gA2hx31AtLGtFuESW6LHneKPHXsJHO9Cjgg21V3I/pQ468DSezoJKSKKIh5siPrXQgDfNE
 aKyW5zKZ2W8t2eH2s6TUFDOO1rlY5uR1daA82RW4yCwhZZ8vDgED895k9Jz7UBl3L7FXWgje/O0
 5cyeWDXeJO4/9FXYicdfioCyx1RfGCWeX1vuYF6D/3i1Gi5skpDk8JnbLoXi8ceiEcKs1BlG0fj
 WaQTJWyja/p9cTerSO10qHm/InIVUAPmMUBkyopOFQ8A==
From: lo🤑pel111@cockatielone.biz
To: deg4m4@gmail.com
Subject: Hello Test Message!
Message-ID: <784db069-057c-c727-3a2a-1182cd67b741@cockatielone.biz>
Date: Wed, 07 Aug 2024 03:16:54 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-1192bd904e2e4d3f-Part_1"

----_NmP-1192bd904e2e4d3f-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

Hello World =F0=9F=A4=91=F0=9F=A4=91
----_NmP-1192bd904e2e4d3f-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-1192bd904e2e4d3f-Part_1--

`
