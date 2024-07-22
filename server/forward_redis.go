package serverapp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/p4rthk4/u2smtp/pkg/config"
	"github.com/p4rthk4/u2smtp/pkg/server"
	"github.com/redis/go-redis/v9"
)

type MailFwdBackendRedis struct {
	server.ForwardBackend

	host     string
	port     int
	username string
	password string
	db       int

	ctx    context.Context
	client *redis.Client
}

func (mailFwd *MailFwdBackendRedis) Init() {
	// load config
	mailFwd.host = config.ConfOpts.RedisConfig.Host
	mailFwd.port = config.ConfOpts.RedisConfig.Port
	mailFwd.username = config.ConfOpts.RedisConfig.Username
	mailFwd.password = config.ConfOpts.RedisConfig.Password
	mailFwd.db = config.ConfOpts.RedisConfig.DB

	mailFwd.ctx = context.Background()

	mailFwd.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", mailFwd.host, mailFwd.port),
		Password: mailFwd.password, // no password set
		DB:       mailFwd.db,       // use default DB
	})

	_, err := mailFwd.client.Ping(mailFwd.ctx).Result()
	if err != nil {
		log.Println("Redis Connection Faild...")
		if config.ConfOpts.Dev {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	log.Println("Init Redis Forward method")
}

func (mailFwd *MailFwdBackendRedis) ForwardMail(email server.Email) {
	fmt.Println("Mail Recive in Redis")

	fmt.Println("Redis Host:", mailFwd.host)
	fmt.Println("Redis Port:", mailFwd.port)
	fmt.Println("Redis Password:", mailFwd.password)
	fmt.Println("Redis DB:", mailFwd.db)

	data, err := email.ToDocument()
	if err != nil {
		log.Println("error to gen document")
	}

	fmt.Println(string(data))

	err = mailFwd.client.Set(mailFwd.ctx, email.Uid, string(data), 0).Err()
	if err != nil {
		log.Println("email add error into redis")
		if config.ConfOpts.Dev {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("email add successful")

	emailS, ok := email.ParseMail()
	if !ok {
		fmt.Println("Error In mail Parse")
		return
	}

	fmt.Println("email: ")
	fmt.Println(emailS)
}

func (mailFwd *MailFwdBackendRedis) ExistMailBox(rcpt string) bool {
	return checkMaiboxFromRcpt(rcpt)
}