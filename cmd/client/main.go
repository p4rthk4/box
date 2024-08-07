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
	clinet.SetFrom("parthdegama@cockatielone.biz")
	clinet.SetRcpt("pthreeh@outlook.com")
	clinet.SetHostname("mx.myworkspacel.ink")

	// clinet.DSNReturn = smtpclient.DSNReturnFull
	// clinet.UTF8 = true

	clinet.SetData([]byte(mail))
	clinet.CheckTlsHost = false
	err := clinet.SendMail()
	fmt.Println(err)
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=PKtoTjqHuudr64nYpizw9stW9CenBfl0sXHFI8KEfBA=;
 h=from:subject:to:mime-version:content-type:content-transfer-encoding;
 b=FyH1Vt0URXqWlRQz72cMnDs5U4De1YzU0CAVgVj1+dXpinlvXS5/UviNIg3tzBS7KP1gCjiK0
 xgaQ+TpBNezL+YKenR0wCC+YZDtgq/QVOrUlogVTzfkOQE+QPjs0YaVULN/t6y5VRjBz/JXIuBf
 v/ADu0T2allxJXh6yLtorP8S8dwjvov4zJIIN5oKwXylKIAbfYd3SQ76JLEzPGjQ4neBpEC1eD1
 cBbG2NnJIbe4B75CYOUJDv0jri5vvDRX7cAaQNuW4cUfaOyg02s9MsANWlIkZHfN7uDhJ86CsCV
 sUExIeN1ZKtAJsywB1UEeG0LbYBCOlIhVxMzV2/lua2A==
From: parthdegama@cockatielone.biz
To: pthreeh@outlook.com
Subject: Hello Test Message From Parth!
Message-ID: <7a20877a-ec53-0dd0-cd36-624af1c4d146@cockatielone.biz>
Content-Transfer-Encoding: 7bit
Date: Wed, 07 Aug 2024 03:46:06 +0000
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8

Hello World!, my name is parth degama and your name
`