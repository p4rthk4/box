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
	clinet.SetFrom("parthdegama@myworkspacel.ink")
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
 q=dns/txt; s=draft; bh=10ksEvlgKQRSmdzeDKsGBQ3BXzRiT4pq9KuMlBwqjqg=;
 h=from:subject:to:mime-version:content-type;
 b=SkFp6oV3PBAguAGQE64B8SF1jUhuWXqdFxK6V0STCl3Iriey8DcPkASHqkJulE/C2CD6l/ys/
 PvaEAggikO8heKu2v2t5qrxANSYmISwj/Lva1Gza508neHbepruMbbRW2xMw7JuSIRn6jB8k9mI
 JolzwXWKkonluRI4yalqCAxCmRfMIafi6seliyW/0g6ewGqyUKYFtS/++OhEqTEZVAo//5uI5Cw
 bxqhsj33Pcw5ei0ZzsZINLFpqiCLAULejqEp6mbSllrbP1frQfOxJjrK88/0GOqQubCYYmXIshG
 jZzFqp4uWdjaZNrcgLAtnWqYPpKwzhjumHgmQISXYJBg==
From: parthdegama@myworkspacel.ink
To: parthka.2005@gmail.com
Subject: Test Message!
Message-ID: <74c858e5-79bf-f1ea-8334-111f39b67901@myworkspacel.ink>
Date: Tue, 06 Aug 2024 17:01:25 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-2bfb0752aee9fb57-Part_1"

----_NmP-2bfb0752aee9fb57-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-2bfb0752aee9fb57-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-2bfb0752aee9fb57-Part_1--

`
