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
	clinet.SetFrom("parthdegama@cockatielone.biz")
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
 q=dns/txt; s=draft; bh=/W9YZjqmeW2SEYNlLZjjGzd//yNcXYtg/3+BADllX+Q=;
 h=from:subject:to:mime-version:content-type;
 b=T4Dr0uoR5IFYWKgwoK4wjXlr0BcQHxgeAdh1j0XGOeyOg6fVPhLHGmUYV1miSGNvBFEqugtaC
 w7wJEoBe1NUgQs2/D+18bRHD/Amv7pUVf8KWvbIHkt26qhYZKzWv2sF9upxIVPiD9rdFtFYAhqq
 4LbgbZm7vXdTIpXgdBUoD1/TYMhyxi2m3PP3KXLsWY7+6qJB3obQZhtZk8tY3EDAqFfEzJ+9pkC
 LjIsfI9B86iuKG0IldGGW+GBRlB87alpAclZugmvfg2g6LZbZcWt9rfaN1KUtzeljUzWD1zknLn
 qlin2AUjvNH1BCQXWe3NiG+iA7naDC2irj1UrPIlDUhg==
From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From India!
Message-ID: <652f6881-662e-ce06-9f3f-1bddd2c8af3a@cockatielone.biz>
Date: Wed, 07 Aug 2024 05:04:12 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-a192fa3e5098ee4e-Part_1"

----_NmP-a192fa3e5098ee4e-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-a192fa3e5098ee4e-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-a192fa3e5098ee4e-Part_1--

`
