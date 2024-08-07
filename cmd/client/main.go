package main

import (
	"fmt"
	"strings"
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
	clinet.SetFrom("aly@cockatielone.biz")
	clinet.SetRcpt("hello@parthka.dev")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	if !strings.Contains(mail, "\r\n") {
		fmt.Println("Not Any \\r\\n found!!!")
		mail = strings.ReplaceAll(mail, "\n", "\r\n")
	}
	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=gNV3UrXsTQZbe7uc/s9LXtlsoIcIdnFRHD8qR38qDWg=;
 h=from:subject:to:mime-version:content-type;
 b=MbJkBsa3oRCryWhWdrVJpmpJODaJm/fKDnPlC8jXLfrAB433vKTu4UB9huMVMphkqrixgKsfe
 V0LvhsLy8eA+SkiA0iBZfbvokR/tMDoahXgibcpmmTPJZhWkO4tkG1DxE0OwHsz7B989F7nYk5m
 dW/0NVpPcIweyBSwJtSkv4IfYDQeW7DltEdgaJ8DheuxNccokGdxYCK+uqtwt1Zh1s1YOvhH1BN
 um8/fd7cAbowfFi3x3C12VKEyQaapDjf1b8n8AunKnHsLpzdgLHZnsIrBF0125Km0WqE3TUkVis
 uuG9ig0/kyyccFyedT6EZuUUl595e+EaAnTVLVCd8Fyw==
From: PARTH <aly@cockatielone.biz>
To: Parthka <hello@parthka.dev>
Subject: Final Helo Mail
Message-ID: <23286281-77a3-f0f1-103b-572eff747d22@cockatielone.biz>
Date: Wed, 07 Aug 2024 06:27:49 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-44fd71683cf1b1f6-Part_1"

----_NmP-44fd71683cf1b1f6-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-44fd71683cf1b1f6-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-44fd71683cf1b1f6-Part_1--

`
