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
	clinet.SetRcpt("parthka.2005@gmail.com")
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
 q=dns/txt; s=draft; bh=HEb4TOg3xzbFG7FxyrCf9B8kmU18dmMz8/VFoFL50TA=;
 h=from:subject:to:mime-version:content-type;
 b=IuIEdcKwe84iwv7J2tQxscxmfL5v97HlWMwVggUxYCzb9TkYp3qfzz3t/keZX74FnMd0oCSVE
 +jgOx017PZED/vQqp6LBGHbuLXJUXc0KXDko7Fgb7SnBazrIDGiQDqmOML6SVwQTCzogc7utThH
 erlRrWlhbts37iCVu41qpbjBn1OXX2guMqusfbB3StA3H5sPf8zlH2YIsXx7+lci+3owFqGJH5f
 t9D3UymhJIdH0zQBgh0AAD6R7AOyaoPfxQRtp+GT6crx6NK1FL2MV1LCnaYIAHN3EA6imvEhR+w
 ci6+5IEbhX5d+HgmZWnF91u6NI/fa12F6K8MP+/Y/SzQ==
From: PARTH <aly@cockatielone.biz>
To: Parthka <parthka.2005@gmail.com>
Subject: Final Helo Mail
Message-ID: <fd1a82e2-c59b-1021-c27b-d3b788bd7980@cockatielone.biz>
Date: Wed, 07 Aug 2024 06:30:31 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-36e3577c80e6410a-Part_1"

----_NmP-36e3577c80e6410a-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-36e3577c80e6410a-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-36e3577c80e6410a-Part_1--


`
