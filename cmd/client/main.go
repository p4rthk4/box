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
 q=dns/txt; s=draft; bh=IBv/jOiuxIzffvZsdUAl/YoXRgPE8s6FJKBLZMnWeV0=;
 h=from:subject:to:mime-version:content-type;
 b=Cl4EkEJjos5/KgGkjKApqGDbvcKuhMAb+LaEbjXDtHb5ymSx1QonLWBvVWum8mOc73OgM9FzP
 RdiMS6CJQ0DrLs7P+JY2fWdNCGKPFQehr3pz5Ih9feYRIOCwzQg5Ufrl5jS/M2UZRwGnkDOCjTE
 Vi8Mww3LVAGXPl2lqHVM/jc/PpRyXDykgrMyufN1tQkZAzjgXhIdyEwY8rMFxMUqFMbmQtTLREx
 +MWXbSAysr2ty4UexjPgP1M28HY/NZ0ihcuNcmmdmqT+3j68DiBc+yHXhRGagHk4rrvJwzQm7Sm
 V9HVBCmdWb5Gdf7h7k3ZMCY9mbnIF+ORGWnciVSaWRSg==
Content-Type: multipart/alternative;
 boundary="----sinikael-?=_1-17230030192850.5446859609884578"
From: parthdegama@cockatielone.biz
To: pthreeh@outlook.com
Subject: Hello Test Message From Parth!
Date: Wed, 07 Aug 2024 03:56:59 +0000
Message-Id: <1723003019287-32e460e8-9d67a2c4-2b162709@cockatielone.biz>
MIME-Version: 1.0

------sinikael-?=_1-17230030192850.5446859609884578
Content-Type: text/plain
Content-Transfer-Encoding: 7bit

Hello World!, my name is parth degama and your name
------sinikael-?=_1-17230030192850.5446859609884578--

`