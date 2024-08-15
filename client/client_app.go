package clientapp

import (
	"github.com/rellitelink/box/config"
)

func StartClient() {
	logger := makeLogger()
	emails := getAmqpConsume(logger, config.ConfOpts.Client.Amqp.Queue)
	forever := make(chan bool)

	for i := 0; i < config.ConfOpts.Client.Worker; i++ {
		go func(workerID int) {
			for d := range emails {
				emailHandler := EmailHandler{
					logger:    logger,
					emailYaml: d.Body,
					wid:       workerID,
				}
				emailHandler.handleClient()
			}
		}(i)
	}

	<-forever
}
