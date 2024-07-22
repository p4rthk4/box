package app

import (
	"sync"

	"github.com/p4rthk4/u2smtp/pkg/config"
	"github.com/p4rthk4/u2smtp/pkg/server"
)

func StartApp() {

	server.SetMailFwdMethod(&MailFwdBackendAmqp{})

	serverWait := sync.WaitGroup{}

	// IPv4 listen
	serverWait.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		smtpServer := server.SMTPServer{
			Host:            config.ConfOpts.Server.Host,
			Port:            config.ConfOpts.Server.Port,
			ServerWaitGroup: wg,
		}

		// fmt.Println(smtpServer)

		smtpServer.SetLogger()
		smtpServer.Listen()
		smtpServer.AcceptConnections()

	}(&serverWait)

	// IPv6 listen
	if !config.ConfOpts.Server.IPv6Disable {
		serverWait.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			smtpServer := server.SMTPServer{
				Host:            config.ConfOpts.Server.HostIPv6,
				Port:            config.ConfOpts.Server.PortIPv6,
				ServerWaitGroup: wg,
				IsIPv6:          true,
			}

			// fmt.Println(smtpServer)

			smtpServer.SetLogger()
			smtpServer.Listen()
			smtpServer.AcceptConnections()

		}(&serverWait)
	}

	serverWait.Wait()
}
