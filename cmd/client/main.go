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
 q=dns/txt; s=draft; bh=sIAi0xXPHrEtJmW97Q5q9AZTwKC+l1Iy+0m8vQIc/DY=;
 h=from:subject:to:mime-version:content-type:content-transfer-encoding;
 b=V8c/X9JzyFXV1jaRSM8ShWz8a9rVkyj8i83U43g3/yRYPooaApbym5dj53JH9YLSJoZfFHIqa
 CgSi9gPw7rhBRDSnPuqyV9guasN6mtB0oj8F6tjALDats1mmECe7Cu72xO0KNvQ+Nb3J3OXO5Ys
 RvWThpuDChPMFJcUYvS7TnSEFb2t+zbnpVtd8K03B0f5CoOvsHvPoAi9a1SY92GIN/c0+MT7jvR
 SbKS0dOFAF33lJ7yKpGAxtvqZFlAi93qDF9gE6wV5m6t7CIm10eIt8ft+PPfOyYYQyh5nOwWnRM
 FdtQZylmfvN8nqQZ1pBugEuumoZ2R+h4QIYwzNNK0w9w==
From: lopel@cockatielone.biz
To: pthreeh@outlook.com
Subject: Hello Test Message From Parth!
Message-ID: <b3ba2989-06f7-aa15-3093-a7a75fb99400@cockatielone.biz>
Content-Transfer-Encoding: 7bit
Date: Wed, 07 Aug 2024 03:26:22 +0000
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8

Hello World
	
`
