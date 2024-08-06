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
	clinet.SetFrom("parthka@myworkspacel.ink")
	clinet.SetRcpt("parthka.2005@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=myworkspacel.ink;
        q=dns/txt; s=draft; bh=qIkyo9AJwHq7QIGLlqRVwF5vnC+cxPXFIpgrPuOdSNI=;
        h=from:subject:to:mime-version:content-type;
        b=dm9qFzUPNua475y5GnLyhI1bLHgJicfRrSgXX7lzmbGCBMfQ1iWWE7ETxLyH4hxrm/85oXd9i
         DKK38P12DFUsquwVGlQSMUYJEBgDZgy5H5fyYWJF4pBk7146KNwX0zZ0KVPRD8p9Er7/BL+FNjq
         4YEeS6bLWl2gSCjPiFpV3Va5u6q+5YCjtizVKCSa3NSoveYMBsFVcO+sLGLIHaerRZ+lOz8I7TN
         UoafraL4SUCDyCne8esSuIUUgehcjRZkA0KEoaZqMD6wWQqjduBB44Ha0gUcR+M8BKNgvDFrRlq
         U5BP6dEUOT87PSw+U9Psf4S+VxhB0UXzejdOHXYRRodw==
From: parthka@myworkspacel.ink
To: Parth <parthka.2005@gmail.com>
Subject: Test Message!
Message-ID: <3123907b-9492-b26f-5b9f-900b75a9d864@myworkspacel.ink>
Date: Tue, 06 Aug 2024 16:19:43 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-aec7bcdd5a19b50d-Part_1"

----_NmP-aec7bcdd5a19b50d-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-aec7bcdd5a19b50d-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-aec7bcdd5a19b50d-Part_1--
`
