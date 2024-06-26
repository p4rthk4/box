// U2SMTP - config
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package config

func defaultConfig() ConfigsOptions {
	return ConfigsOptions{
		Name:     "S2SMTP",
		HostName: "localhost",
		Server: ServerOptions{
			Host:        "",
			Port:        25,
			HostIPv6:    "",
			PortIPv6:    25,
			IPv6Disable: false,
		},
		MaxClients: 0,

		LogDirPath:  "./logs",
		LogFilePath: "email.log",
		Dev:         false,

		ClientGreet: "`Hello, Client!`",
		ClientByyy:  "Ok, Byyy!",

		MaxRecipients: 1024,
		CheckMailBoxExist: false,

		RedisConfig: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},

		Amqp: AmqpConfig{
			Host:     "localhost",
			Port:     5672,
			Username: "",
			Password: "",
			Queue:    "",
		},

		Client: ClientConfig{
			HostName: "",
			
			LogDirPath: "",
			LogFilePath: "",

			Amqp: AmqpConfig{
				Host: "",
				Port: 0,
				Username: "",
				Password: "",
				Queue: "",
			},
		},


	}
}
