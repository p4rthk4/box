package serverapp

import (
	"sync"

	"github.com/p4rthk4/u2smtp/config"
	"github.com/p4rthk4/u2smtp/pkg/server"
)

func StartServer() {

	server.SetConfig(server.SMTPConfig{
		Name:     config.ConfOpts.Name,
		HostName: config.ConfOpts.HostName,
		ESMTP: server.ESMTPOptions{
			Enable:      true,
			Tls:         config.ConfOpts.Tls.StartTls,
			Utf8:        true,
			BinaryMime:  true,
			MessageSize: config.ConfOpts.MessageSize,
		},

		ClientGreet: "Hello, Client!",
		ClientByyy:  "Ok, Byyy!",

		SpfCheck:      config.ConfOpts.SpfCheck,
		MaxRecipients: config.ConfOpts.MaxRecipients,
		MaxClients:    config.ConfOpts.MaxClients,

		CheckMailBoxExist: config.ConfOpts.CheckMailBoxExist,
		TlsConfig: server.TlsKeyCert{
			Key:  config.ConfOpts.Tls.Key,
			Cert: config.ConfOpts.Tls.Cert,
		},

		LogDirPath:  config.ConfOpts.LogDirPath,
		LogFilePath: config.ConfOpts.LogFilePath,
		Dev:         config.ConfOpts.Dev,
	})

	server.SetMailFwdMethod(&MailFwdBackendAmqp{})

	serverWait := sync.WaitGroup{}

	// IPv4 listen
	serverWait.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		smtpServer := server.SMTPServer{
			Host:            "0.0.0.0",
			Port:            config.ConfOpts.Port,
			ServerWaitGroup: wg,
		}

		// fmt.Println(smtpServer)

		smtpServer.SetLogger()
		smtpServer.Listen()
		smtpServer.AcceptConnections()

	}(&serverWait)

	// IPv6 listen
	serverWait.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		smtpServer := server.SMTPServer{
			Host:            "::",
			Port:            config.ConfOpts.Port,
			ServerWaitGroup: wg,
			IsIPv6:          true,
		}

		// fmt.Println(smtpServer)

		smtpServer.SetLogger()
		smtpServer.Listen()
		smtpServer.AcceptConnections()

	}(&serverWait)

	serverWait.Wait()
}
