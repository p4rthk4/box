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
 q=dns/txt; s=draft; bh=olxGR7wvqmpGCBHsxWlk5wgwYwId+jmy2Zz3NqRgoFo=;
 h=from:subject:to:mime-version:content-type;
 b=cNy3zuVXparY6apEr7xIMWfhO3gTIXMIQiS5upB+kjEi4RD7o+dAO/Mu6GiyKzLRqhu3etzu2
 D3MkVNd92yZ9GI+gMUajvvWPJycYMdD4S0pDEuczdshdVJ7PL2Fju+8eu8QX1PTreMmJUeJcxWw
 k8D7Rw0t7+DeujdZOOzmjOGhzpX7qsdyOksltxVEvw4l3e5WERpc6wtPFWtlY2AXdiiR+eRgY2L
 X6uTqDu5Wu+NLoqzrXswopTPYqG7sAjncjdTAM2lsVPdoC52YpN3HtqM64cagLFkta1tCqX7FK8
 w+avn4FuRyJym9gGHkBki8miRYvHzMnEa3/s4sIF5eUw==
Content-Type: multipart/alternative;
 boundary="----sinikael-?=_1-17230040309390.004235149441492503"
From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From Parth!
Date: Wed, 07 Aug 2024 04:13:50 +0000
Message-Id: <1723004030941-22df59a5-8ce462be-8f469442@cockatielone.biz>
MIME-Version: 1.0

------sinikael-?=_1-17230040309390.004235149441492503
Content-Type: text/plain
Content-Transfer-Encoding: 7bit

Hello World!, my name is parth degama and your name
------sinikael-?=_1-17230040309390.004235149441492503--

`