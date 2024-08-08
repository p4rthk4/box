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
	clinet.SetTimeout(5 * time.Second)
	clinet.SetFrom("aly@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@myworkspacel.ink")
	clinet.SetHostname("mx.myworkspacel.ink")

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
 q=dns/txt; s=draft; bh=28InEHXYZSFdNCv2dBpPHBwEr+J3uxCsYyQldgTfMOw=;
 h=from:subject:to:mime-version:content-type;
 b=A4AJt5quASgN7/9QF03XsLteS6KeyUYxoaoF7YktkI1yCFIf5D//ajzzLZPVsYNp7RSFWfKEo
 +FjVDoYu9lgQMQaCEjS/loPxme2BCOtzmaBbmfk2DfHHLwJ20P6GzdGo+9vBpxyLjuxCz0Jm0jp
 B6RbTOwjxLdYll79kilJm7gwbQ6Vjet2sCPzjkcha4oASkl5KHe1kM7/BAhQDp7O2krp9tkpG8f
 IvGPOqrRgXPBCNe6CCzW+pvO5JqQNwOsgo8+4zThTIW1sBEF5nizzD01zJnrkFFpKNM0x+spMai
 jnB0WRB/d2LmhVk32JIlV6iqLaiHzreEFOATxfIB+R8Q==
From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <parthka.2005@yandex.com>
Subject: Hello Test Message From India!
Message-ID: <860f0464-2805-45b1-c6a4-6624fa060a6b@cockatielone.biz>
Date: Wed, 07 Aug 2024 17:48:32 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-f66ddb6c316e9fc2-Part_1"

----_NmP-f66ddb6c316e9fc2-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGVsbG8gV29ybGQhLCBteSBuYW1lIGlzIHBhcnRoIGRlZ2FtYSBhbmQgeW91ciBuYW1l
----_NmP-f66ddb6c316e9fc2-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPjxiPkhlbGxvIFdvcmxkISwgbXkgbmFtZSBpcyBwYXJ0aCBkZWdhbWEgYW5kIHlvdXIgbmFt
ZTwvYj48L2gxPg==
----_NmP-f66ddb6c316e9fc2-Part_1--

`
