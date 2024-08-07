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
	clinet.SetRcpt("parthka.2005@proton.me")
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
 q=dns/txt; s=draft; bh=Y2eUb0vlqw1429h0ECF512DoA5SOHdeGf0Qy9i0G8Zw=;
 h=from:subject:to:mime-version:content-type;
 b=aZrfOVutIURRrm1X8UZjaclcvWwC34YxojJfgwvmaGC2BmRezAdHqqhh1HCjbcmhCjgry0ArZ
 IsVJGHUubJkQa492fMkPIVPRrM7cGU0eVdoAATzCM2R+Zl/0M0ckb1X0ecqc6ANhLMI8muSC1xI
 t+ysNsHYu2BmnUjH4OVyBYsarLiY6RU1nZMVUbt6LPEyLDu7+L1USP10VYSFadXIvU0IHWD5lPd
 A2ht343J5oN27cURqYSuNcC5D33wBE+sowDtduwfVImSiH1XMcMFqolvSGdENJmYQTmqvB3msMX
 pGN0EsYWuC7z9bNrrs9He6Jc/FQCyZo2o+qcKsEi3Uug==
From: PARTH <aly@cockatielone.biz>
To: Parthka <parthka.2005@proton.me>
Subject: Final Helo Mail
Message-ID: <5020211e-4481-cca3-968f-5b7ad80f24a9@cockatielone.biz>
Date: Wed, 07 Aug 2024 06:31:54 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-a816edaefe471680-Part_1"

----_NmP-a816edaefe471680-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-a816edaefe471680-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-a816edaefe471680-Part_1--

`
