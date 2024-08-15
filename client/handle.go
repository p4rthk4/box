package clientapp

import (
	"fmt"
	"time"

	"github.com/rellitelink/box/config"
	smtpClient "github.com/rellitelink/box/pkg/client"
	"github.com/rellitelink/box/pkg/logx"
	"gopkg.in/yaml.v3"
)

type EmailHandler struct {
	em        *EmailYAML
	emailYaml []byte

	wid    int
	logger *logx.Log
}

type EmailYAML struct {
	Id        string `yaml:"id"`
	From      string `yaml:"from"`
	Recipient string `yaml:"recipient"`
	Data      string `yaml:"data"`
}

func (eh *EmailHandler) handleClient() {
	em, err := openMail(eh.emailYaml)
	if err != nil {
		eh.logger.Error("error on mail parser, worker-%d: %s", eh.wid, err.Error())
		return
	}

	eh.em = em
	eh.logger.Info("Worker %d received a message: %s", eh.wid, eh.em.Id)

	res := eh.sendMail()
	errStr, err := yaml.Marshal(res)
	if err != nil {
		eh.logger.Error("error on make status yaml, worker-%d: %s", eh.wid, err.Error())
		return
	}

	fmt.Println(string(errStr))
}

func (eh *EmailHandler) sendMail() smtpClient.ClientResponse {
	client := smtpClient.NewClinet()
	client.Logger = eh.logger.GetNewWithPrefix(eh.em.Id)
	client.Timeout = time.Second * 10

	client.SetHostname(config.ConfOpts.Client.HostName)
	client.SetFrom(eh.em.From)
	client.SetRcpt(eh.em.Recipient)

	client.SetData([]byte(eh.em.Data))

	client.StartTls = config.ConfOpts.Tls.StartTls
	client.TlsCert = config.ConfOpts.Tls.Cert
	client.TlsKey = config.ConfOpts.Tls.Key

	client.SendMail()
	res := client.GetResponse()
	return res
}
