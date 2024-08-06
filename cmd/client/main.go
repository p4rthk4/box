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
        q=dns/txt; s=draft; bh=s2qpmZEFQIbCjNBfNLLS3B4ppmEvvhogCEM45G2nna4=;
        h=from:subject:to:mime-version:content-type;
        b=IGvFQePiKUulKsPpf3V0PYSdDB409wQHs7LOijutvGq7BEhG4o9w1SUl+2FksleaRSB+ueZ2L
         OonPqN/HcMv1IgOO52PuFl02EoOG7hLybiEGFBjUgMEt8Xu3grmGznydbDopVA3Msxy+79/k2Sa
         qesB0BHf293lZ87MKgSzLHYIhSvUNtVmpfDtBLGqYWd1PH7uLlkMxmbTkuJ/JXfoPm+8N2DqUJF
         zVn+tcxjp8BlJcCOhK6scrggvQoDqtF8iD+DHVFpDwPt0sf6SvG9DmJ5TiP4mdbNQeQd3PHWgLU
         dwc0PzJP5OP3CGcj5h3YRFKMoNePTZHwdL9wQuYFP+ew==
From: parthka@myworkspacel.ink
To: Parth <parthka.2005@gmail.com>
Subject: Test Message!
Message-ID: <5e199cdf-564b-b17e-81fd-b759e613795a@myworkspacel.ink>
Date: Tue, 06 Aug 2024 15:49:17 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-164e1edfd71a02c6-Part_1"

----_NmP-164e1edfd71a02c6-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-164e1edfd71a02c6-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-164e1edfd71a02c6-Part_1--
`
