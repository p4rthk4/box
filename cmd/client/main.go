package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/p4rthk4/u2smtp/config"
	smtpclient "github.com/p4rthk4/u2smtp/pkg/client"
	"github.com/p4rthk4/u2smtp/pkg/logx"
)

func main() {

	config.LoadConfig() // load conifg in config file

	clinet := smtpclient.NewClinet()
	
	// logger
	logFile := os.Stdout
	if !config.ConfOpts.Dev {
		var err error
		logFile, err = os.OpenFile(config.ConfOpts.LogDirPath+"/"+config.ConfOpts.LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logx.LogError("Error opening file:", err)
			return
		}
	}
	logger := logx.NewLoggerWithPrefix(logFile, "EMAIL")
	clinet.Logger = logger
	
	clinet.SetHostname(config.ConfOpts.HostName)
	
	clinet.SetFrom("aly@cockatielone.biz")
	clinet.SetRcpt("parthka.2005@myworkspacel.ink")
	
	if !strings.Contains(mail, "\r\n") {
		fmt.Println("Not Any \\r\\n found!!!")
		mail = strings.ReplaceAll(mail, "\n", "\r\n")
	}
	clinet.SetData([]byte(mail))

	clinet.CheckTlsHost = false
	clinet.TlsKey = config.ConfOpts.Tls.Key
	clinet.TlsCert = config.ConfOpts.Tls.Cert
	
	clinet.Timeout = 5 * time.Second
	
	clinet.SendMail()
	fmt.Println(clinet.GetResponse())
}

var mail string = `DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=cockatielone.biz;
 q=dns/txt; s=draft; bh=TeIFrF74hDBEkjt9OeUkEQZM+ljjmJFn59ulnGw49JQ=;
 h=from:subject:to:mime-version:content-type;
 b=V0bD7B0eplOnGaL27lwIsuXgLBsig2ZbEmDJwEwdIfDtCOmyCcgNgOqD9rthM7/Epr9HyrSOy
 34oL1XQMfEcXF9pCGELQylaCo5q1L7FmixE3cNr7vkTF+ZVjwMgIfKIkObmbB+q5KBjaI9iGfNP
 2KUwvgpbNlJIBMCITQRF3jG7tHDRXOO0Uqp/Nc0d2fm8MGUAlWgD4jx9he09suzhgtpuT/adcEz
 SgaBXc3WxFepahEko+n/wuoLyCl+bZ6GP5GZtmEUhMkGEgZAQRdm+PmvyJ99QDo6k0TDJuuBxkv
 lPxEjCyA4/qdNnkQoQ+mFW+zj6fBTVn3SSjSxXGka+gw==
From: PARTH <parthdegama@cockatielone.biz>	
To: pthreeh@outlook.com
Subject: Hello, I'm From India!
Message-ID: <e560e7ed-7e22-9575-23a2-606026b6860a@cockatielone.biz>
Date: Sat, 10 Aug 2024 02:08:27 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-9a6384c3eb8355d1-Part_1"

----_NmP-9a6384c3eb8355d1-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGkgUGFydGghIGkgYW0gS1QKCiBha2EgRGhhcnZp
----_NmP-9a6384c3eb8355d1-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPkhpIFBhcnRoITwvaDE+PGgzPkFrYSBMb2xraW5nPC9oMz4=
----_NmP-9a6384c3eb8355d1-Part_1--

`
