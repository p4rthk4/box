package clientapp

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rellitelink/box/config"
	"github.com/rellitelink/box/pkg/client"
	"github.com/rellitelink/box/pkg/logx"
	"gopkg.in/yaml.v3"
)

type Email struct {
	Id        string `yaml:"id"`
	From      string `yaml:"from"`
	Recipient string `yaml:"recipient"`
	Data      string `yaml:"data"`
}

func ClientClient() {

	em, err := openMail([]byte(mailYaml))	
	if err != nil {
		panic(err)
	}

	fmt.Println(em)

	clinet := client.NewClinet()

	// logget
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

	clinet.SetFrom(em.From)
	clinet.SetRcpt(em.Recipient)

	if !strings.Contains(em.Data, "\r\n") {
		panic("CRLF no found in mail")
	}
	
	clinet.SetData([]byte(em.Data))

	clinet.CheckTlsHost = false
	clinet.StartTls = false
	clinet.TlsKey = config.ConfOpts.Tls.Key
	clinet.TlsCert = config.ConfOpts.Tls.Cert

	clinet.Timeout = 2 * time.Minute

	clinet.SendMail()
	fmt.Println(clinet.GetResponse())
}

func openMail(data []byte) (*Email, error) {
	e := new(Email)
	err := yaml.Unmarshal(data, e)
	if err != nil {
		return nil, err
	}

	if e.Id == "" {
		return nil, fmt.Errorf("email id not found")
	}

	if e.From == "" {
		return nil, fmt.Errorf("email from field not found")
	}

	if e.Recipient == "" {
		return nil, fmt.Errorf("email recipient field not found")
	}

	if e.Data == "" {
		return nil, fmt.Errorf("email data not found")
	}

	return e, nil
}

const mailYaml string = `
---
id: dsad
from: hello@pa.com
recipient: parthka.2005@gmail.com
data: "parth\r\n"
`
