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
	clinet.SetRcpt("parthka.2005@gmail.com")
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
 b=j57fHsYYj7BVRbZ++sxSmJyCR5d+XHsax5F6Dk1kYdapSOh3DjNUbLFMp8HDRELuz+LHgHlVc
 HFNPdOcO/lHZBis2DoI0+3etH1A1/QXhJGkgYv6ACdJpYVvDyKKs5uCMVDB9t/XyhJVzRLvFN1S
 9yN9CTCL3Mfv2EmGzWna9+wcxmWTNevSTeme9NgmUr9JClo1cpPVMoDmPCCN4/2QpvNepMuCeR3
 o732Zce2FKgN7nOXVwBHt7JjQTCskwYTsTsPDu/bEv+gSv0eQgc063lcklxQ+ginBpeXV3fUrt0
 rF6YaKg9GVxHA5ApHTWAjmDv8Ykw68k/b6TE52b3Yr6w==
From: parthdegama@cockatielone.biz
To: parthka.2005@gmail.com
Subject: Hello Test Message From Parth!
Message-ID: <4e1ec5df-75b9-a6ac-055a-c9c9d4933b1b@cockatielone.biz>
Content-Transfer-Encoding: 7bit
Date: Wed, 07 Aug 2024 03:44:08 +0000
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8

Hello World!, my name is parth degama and your name






`