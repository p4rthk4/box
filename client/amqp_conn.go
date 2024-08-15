package clientapp

import (
	"fmt"
	"net/url"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rellitelink/box/config"
	"github.com/rellitelink/box/pkg/logx"
)

func getAmqpConn(logger *logx.Log) *amqp.Channel {
	encodedPassword := url.QueryEscape(config.ConfOpts.Amqp.Password)
	amqpUri := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.ConfOpts.Amqp.Username, encodedPassword, config.ConfOpts.Amqp.Host, config.ConfOpts.Amqp.Port)
	conn, err := amqp.Dial(amqpUri)
	if err != nil {
		logger.Error("error on amqp conn: %s", err)
		os.Exit(1)
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("error on amqp channel: %s", err)
		os.Exit(1)
	}

	return ch
}

func getAmqpConsume(logger *logx.Log, queueName string) <-chan amqp.Delivery {

	ch := getAmqpConn(logger)

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Error("error on amqp queue declare: %s", err)
		os.Exit(1)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Error("error on amqp consume: %s", err)
		os.Exit(1)
	}

	return msgs
}
