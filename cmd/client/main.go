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
	clinet.SetRcpt("parthka.2005@cockatielone.biz")
	
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
 q=dns/txt; s=draft; bh=qsJaBTlVweNE/meK3+bX5jaRZpKJ8Ym17OllQAAQWmg=;
 h=from:subject:to:mime-version:content-type;
 b=Ei3hPdKWQT26APqldl5VUJ9fFW/KJqzSbSYqVt4LVyE6qZx5gLcC7YCbmXDjydKWvXKfEkxfK
 ruhHnJUhz/Gn9EzJM2JhoPgGG6sAt7gG2FccFlwUseyZvbTN+dF1S2pUe2hn96pv1zFX/0D8ptg
 dkMc5IMi4sBp/GCuMxC3vUes12anZG67vBEgSgMEmah/Ubiw6CEmVx09KHtRK+U3Hyl50r46xys
 KFLE16L+lopQewp6EvInExd35IE+4KMbyN2zNYjAP+viXUotJFkDLWYHXg67xQdWyT86QcmoWLe
 RgPgSuoaa2XdeNFOEIuKBMRd6OKguUms+b/OiwgSG+3A==
From: PARTH <parthdegama@cockatielone.biz>
To: pthreeh@outlook.com
Subject: Hello Test Message From India
Message-ID: <1cd003df-76dd-eb31-e228-efaa49e71d94@cockatielone.biz>
Date: Fri, 09 Aug 2024 11:50:30 +0000
MIME-Version: 1.0
Content-Type: multipart/alternative;
 boundary="--_NmP-4da7f6a060b1093c-Part_1"

----_NmP-4da7f6a060b1093c-Part_1
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: base64

SGkgUGFydGghIGkgYW0gamhvbgoKIGFrYSBsb2wga2luZw==
----_NmP-4da7f6a060b1093c-Part_1
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: base64

PGgxPkhpIFBhcnRoITwvaDE+PGgzPkFrYSBMb2xraW5nPC9oMz4=
----_NmP-4da7f6a060b1093c-Part_1--


`
