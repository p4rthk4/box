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
	clinet.SetFrom("lolk@myworkspacel.ink")
	clinet.SetRcpt("deg4m4@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=myworkspacel.ink;
        q=dns/txt; s=draft; bh=ujNdhHpDGWw+POIbrDZgZUuaWfr351c6LCnwhuqZx+c=;
        h=from:subject:to:mime-version:content-type;
        b=JikMxbuwqYh0ZVsc9IQesnqsW9ZDKaOoC4ibTITqTVGqvo3QZepcox8NoiU8K1rY9l2uIXxnj
         JfZmoTZFeHKg2WcbNlE7w2LdpXdn6YfGn/eej10GXN9wt5M0USVrWcJDRrK/RTDLXCZZs0f8ATq
         0nbRVsl3mNUiRWQP4X0dNHZwqIXYxV+C42gGzgBk/UppFZfdkhcLw0XwxEU3GJBVtEwhqHZ8crx
         L/mBOt51osBePtiePN7BaKts6Wj/Va66Iv7uNn1rEql7Y5aZ9iFicZlozfdFo94xJHnWdq9Oa3O
         LteVs7V1hYj+OTtipOhYSJayJ3PQJp0SXyQxnnq6dCSw==
From: lolk@myworkspacel.ink
To: deg4m4@gmail.com
Subject: Test Message!
Message-ID: <d1cd7af0-42e8-dcda-5415-c831ee005e20@myworkspacel.ink>
Date: Tue, 06 Aug 2024 16:55:50 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-fba039d6714fd8a7-Part_1"

----_NmP-fba039d6714fd8a7-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-fba039d6714fd8a7-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-fba039d6714fd8a7-Part_1--


`
