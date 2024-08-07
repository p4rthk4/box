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
 q=dns/txt; s=draft; bh=MmGHooEuTk74lX8yQJXyZbJE4JwzASjnQOkoeF2bb2k=;
 h=from:subject:to:mime-version:content-type;
 b=OHBmNdo4zEY7bU9DT/uP9/3B17Y0lFdZ/10G/5f7T/AEPr6wIsmdhstGRs9ff/5Skpfq5ugqQ
 atVbFoWYHCENMhaHoIqVqRXtOjWDjL1WplgJS+HurEgR4A1ztGORwbG+d3WeCz3taf5z+vx9tGz
 cmk5/Z+LQwLAOdRpnUQMWZUfUVkMAPtRlHXvoRlw/fX4R5GOFcMEl4HGljQnxmWmX0HmDNeajqh
 yWKazRdnWw46q3KPsdSN/OCzmsnnKqCNU9bdHw+AjRjO/1izjM4QcE7Yo2GOWr313BVPNdmk/2k
 rzdGaAKV7xwbT1TOXO0xVReigePyRL9uZBoU7JJZJv3Q==
Content-Type: multipart/alternative;
 boundary="----sinikael-?=_1-17230042791050.9117167673958415"
From: PARTH <parthdegama@cockatielone.biz>
To: Parthka <pthreeh@outlook.com>
Subject: Hello Test Message From Parth!
Date: Wed, 07 Aug 2024 04:17:59 +0000
Message-Id: <1723004279108-5f956586-43ce2ed3-04b9ea31@cockatielone.biz>
MIME-Version: 1.0

------sinikael-?=_1-17230042791050.9117167673958415
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

U0dWc2JHOGdWMjl5YkdRaExDQnRlU0J1WVcxbElHbHpJSEJoY25Sb0lHUmxaMkZ0WVNCaGJtUWdl
VzkxY2lCdVlXMWw=
------sinikael-?=_1-17230042791050.9117167673958415--

`