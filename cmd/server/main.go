// U2SMTP - server cmd
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package main

import (
	"fmt"
	"sync"

	conf "github.com/p4rthk4/u2smtp/pkg/config"
	"github.com/p4rthk4/u2smtp/pkg/server"
)

func main() {
	conf.LoadConfig()
	fmt.Println(conf.ConfOpts)

	serverWait := sync.WaitGroup{}

	// IPv4 listen
	serverWait.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		smtpServer := server.SMTPServer{
			Host:            conf.ConfOpts.Server.Host,
			Port:            conf.ConfOpts.Server.Port,
			ServerWaitGroup: wg,
		}

		fmt.Println(smtpServer)

		smtpServer.SetLogfile()
		smtpServer.Listen()
		smtpServer.AcceptConnections()

	}(&serverWait)

	// IPv6 listen
	if !conf.ConfOpts.Server.IPv6Disable {
		serverWait.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			smtpServer := server.SMTPServer{
				Host:            conf.ConfOpts.Server.HostIPv6,
				Port:            conf.ConfOpts.Server.PortIPv6,
				ServerWaitGroup: wg,
				IsIPv6:          true,
			}

			fmt.Println(smtpServer)

			smtpServer.SetLogfile()
			smtpServer.Listen()
			smtpServer.AcceptConnections()

		}(&serverWait)
	}

	serverWait.Wait()

}
