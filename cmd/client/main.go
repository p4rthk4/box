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
	clinet.SetFrom("parthdegama@cockatielone.biz")
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
 q=dns/txt; s=draft; bh=/He0rngpZD20KignV/pZ1EOiNFQ4twVmG8JfDoFc6LQ=;
 h=from:subject:to:mime-version:content-type;
 b=ZYV6V8rMzYcI8nym0RjZPh+JRrgNcE+Kob0D1/4B5mxkcdW1tr9m+WPmHCQ5oEP5rzO5lOfKG
 m4JVA5c2FQFvCiIu0hyyxycZw2bAxA+I1cMPtcVLR59UYVJI8tmpbnL7hycKyTMCIP70aBgVerL
 gMKCJIwShQFNEFDw8nHq9wksWpXt7kiJ4Sr69+32MRyDh2+zrdGbcAepPzkt74btk7DFQ1O2+c6
 Km+d5MoPYHE6i5S9qLojpGIMmtgG0EAJunQ1jo0bIauISwUWStJivDxbs3XOVTZ/PfT8zVP250T
 39XyG8n8OENUmo4niA9w1vxJqBfuhlJXHCrD5VEboYFQ==
Content-Type: multipart/alternative;
 boundary="----sinikael-?=_1-17230038801180.23133196497826436"
From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <parthka.2005@gmail.com>
Subject: Hello Test Message From Parth!
Date: Wed, 07 Aug 2024 04:11:20 +0000
Message-Id: <1723003880120-2db85b81-708ee516-1ca6f483@cockatielone.biz>
MIME-Version: 1.0

------sinikael-?=_1-17230038801180.23133196497826436
Content-Type: text/plain
Content-Transfer-Encoding: 7bit

Hello World!, my name is parth degama and your name
------sinikael-?=_1-17230038801180.23133196497826436--

`