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
 q=dns/txt; s=draft; bh=syf40hXwgZLxjCerTMBtjRy78eOxi8/4c07t9KzYT70=;
 h=from:subject:to:mime-version:content-type;
 b=gHvhZj2wr/YrTbjqnhCJAqlZrFG3lNZRkMfmYPJQMQ5BquVU+UZA58XP2Z0ORPdlMeSM5tBOY
 0k/ARX41nZvqhydFAczVbOAu2okA5lje1A9vqyVYw3/0lIaxVvhGUfRWf7j9grBSBrNDehC6sXD
 aUtPyLEpHxReSNygV8ZJXDRhbAnp2COWPBEqDU7EURUqGwWM9CPgYsOJ1aap4J8zX1P7EDJtKu8
 8PtexWm4/dFUPw+716YbLcOI3n2xPZMinw7I9O+QG54KHuqRGAUU4dTWv4hx95y8KZktGxBG5Ph
 q7RfLuFLuDeAQ7DZqCNx40sM5DuieVZ9leAzdBiI36nQ==
From: PARTH <aly@cockatielone.biz>
To: Parthka <hello@parthka.dev>
Subject: Final Helo Mail
Message-ID: <49625f91-b609-6619-0164-1177052f5cb9@cockatielone.biz>
Date: Wed, 07 Aug 2024 06:24:53 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-b2260b694d6fd9b8-Part_1"

----_NmP-b2260b694d6fd9b8-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-b2260b694d6fd9b8-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-b2260b694d6fd9b8-Part_1--

`
