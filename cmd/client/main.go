package main

import (
	"fmt"
	"strings"
	"time"

	smtpclient "github.com/p4rthk4/u2smtp/pkg/client"
	"github.com/p4rthk4/u2smtp/config"
)

func main() {

	config.LoadConfig() // load conifg in config file

	clinet := smtpclient.NewClinet()
	// try cockatielone.biz
	// clinet.SetHost("alt3.gmail-smtp-in.l.google.com")
	clinet.SetTimeout(5 * time.Second)
	clinet.SetFrom("aly@cockatielone.biz")
	clinet.SetRcpt("pthreeh@outlook.com")
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
 q=dns/txt; s=draft; bh=awFkpYpIVVN9jx6j/meesqwnbMSMS2rJVGnxh3jV+x8=;
 h=from:subject:to:mime-version:content-type;
 b=ANkOW1I/tTsUpEPClU/OlDE5xPAXxvSU0KmSOp8sW3kxzMUN4a0CG6bafGlgIv83KgmQZLKSF
 9ZvCwrQGNcsMn0SBi1tW0HLzHjZ+uAa5IuyR6oSX7Ijdp4ep/LXPetjxL6LOpsif7r0gJxVYhv6
 5cpepDoJ/GEbfAUZMsBoDHudUmlzfbSSN3IdntF7fsLRQaQ7271NzgRdsujueI1Jzl0FkDmv5vd
 1HJIxx3HA17Aod7ztSnj2CySNmr1eRh0y9v0qwaSD2Ksiqsaus7qsoo3JkSmOswjHQrMSeQKxX/
 3asWLZyfd0iiTJWrov0Vzb9ADyQPyDZdoRVZZwwm7RyQ==
From: PARTH <aly@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Final Helo Mail
Message-ID: <c838c781-af32-149b-640b-cb85611c1f12@cockatielone.biz>
Date: Wed, 07 Aug 2024 06:33:48 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-73cafc4cba2668a6-Part_1"

----_NmP-73cafc4cba2668a6-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-73cafc4cba2668a6-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-73cafc4cba2668a6-Part_1--

`
