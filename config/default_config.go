package config

func defaultConfig() ConfigsOptions {
	return ConfigsOptions{
		Port: 25,

		Name:       "Box - SMTP",
		HostName:   "localhost",
		MaxClients: 0,

		Tls: TlsConfig{
			StartTls: false,
			Key:      "",
			Cert:     "",
		},

		SpfCheck:          true,
		MessageSize:       1024000,
		MaxRecipients:     -1,
		CheckMailBoxExist: false,

		LogDirPath:  "./logs",
		LogFilePath: "email.log",
		Dev:         false,

		Amqp: AmqpConfig{
			Host:        "localhost",
			Port:        5672,
			Username:    "",
			Password:    "",
			Queue:       "",
			StatusQueue: "",
		},

		Client: ClientConfig{
			HostName: "",
			Worker:   2,

			LogDirPath:  "",
			LogFilePath: "",

			Amqp: AmqpConfig{
				Host:        "",
				Port:        0,
				Username:    "",
				Password:    "",
				Queue:       "",
				StatusQueue: "",
			},
		},
	}
}
