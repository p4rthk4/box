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
	clinet.SetFrom("lopel@cockatielone.biz")
	clinet.SetRcpt("deg4m4@gmail.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=8xP40pIb4CiJd05NKONcInIjO4M8/Z9DwA9Xj5NCMDo=;
 h=from:subject:to:mime-version:content-type;
 b=OJYKmbqJTt9PDJoSDICNClPfi4bOev3lhhCWSgZAZZyd7QojHzByDGA4sdv276EV2wHFDQJJS
 /iW8pH1qVwRzdx3OzxsmhEfTNqEW41cH9naZYatddRue9WT07nH3HVZNUu3OpHvw+7vE0Znr2/L
 VCH+FC80InKumiCDU9NqSnQFfypIBXuX+mBoLO16RHzXbP4HXygXq6LOBEtxbXJeWaxlFlwHtOW
 RcJt2ZW17f+rGaQgoujNhWxSV5mlTNIOCdJ755bLv00zSl4SaWFKuyvZePcjljbeKS9QO9DqOY0
 UYD8GcL1l1wFzOGI3db10z9I3rJotydFoSW3/CKcvZyA==
From: lopel@cockatielone.biz
To: deg4m4@gmail.com
Subject: Test Message!
Message-ID: <a7500a67-33d3-9f4d-1456-085daaf70273@cockatielone.biz>
Date: Tue, 06 Aug 2024 18:36:32 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-545163bbd7588cf1-Part_1"

----_NmP-545163bbd7588cf1-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit

Hello World
----_NmP-545163bbd7588cf1-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: 7bit

<h1>Hello!</h1>
----_NmP-545163bbd7588cf1-Part_1--


`
