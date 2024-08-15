package config

type ConfigsOptions struct {
	Name       string    `yaml:"name"`
	HostName   string    `yaml:"hostname"` // disply host name not listen host
	MaxClients int       `yaml:"max_clients"`
	Tls        TlsConfig `yaml:"tls"`

	Port int `yaml:"port"`

	SpfCheck          bool `yaml:"spf_check"`
	MessageSize       int  `yaml:"message_size"`
	MaxRecipients     int  `yaml:"max_recipients"`
	CheckMailBoxExist bool `yaml:"check_mailbox"`

	Amqp AmqpConfig `yaml:"amqp_conf"`

	Client ClientConfig `yaml:"client"`

	LogDirPath  string `yaml:"log_dir"`
	LogFilePath string `yaml:"log_file"`
	Dev         bool   `yaml:"dev"`
}

type TlsConfig struct {
	StartTls bool   `yaml:"starttls"`
	Key      string `yaml:"key"`
	Cert     string `yaml:"cert"`
}

type AmqpConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Queue       string `yaml:"queue"`
	StatusQueue string `yaml:"status_queue"`
}

type ClientConfig struct {
	HostName string `yaml:"hostname"`
	Worker   int    `yaml:"worker"`

	LogDirPath  string `yaml:"log_dir"`
	LogFilePath string `yaml:"log_file"`

	Amqp AmqpConfig `yaml:"amqp_conf"`
}
