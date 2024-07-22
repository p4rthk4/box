package config

type ConfigsOptions struct {
	Name       string        `yaml:"name"`
	HostName   string        `yaml:"hostname"`
	Server     ServerOptions `yaml:"server"`
	MaxClients int           `yaml:"max_clients"`

	LogDirPath  string `yaml:"log_dir"`
	LogFilePath string `yaml:"log_file"`

	ClientGreet string `yaml:"client_greet"`
	ClientByyy  string `yaml:"client_byyy"`

	MaxRecipients     int  `yaml:"max_recipients"`
	CheckMailBoxExist bool `yaml:"check_mailbox"`

	RedisConfig RedisConfig `yaml:"redis_conf"`
	Amqp        AmqpConfig  `yaml:"amqp_conf"`

	Client ClientConfig `yaml:"client"`

	Dev bool `yaml:"dev"`
}

type ServerOptions struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	HostIPv6    string `yaml:"host_ipv6"`
	PortIPv6    int    `yaml:"port_ipv6"`
	IPv6Disable bool   `yaml:"IPv6_disable"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type AmqpConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Queue    string `yaml:"queue"`
}

type ClientConfig struct {
	HostName string `yaml:"hostname"`

	LogDirPath  string `yaml:"log_dir"`
	LogFilePath string `yaml:"log_file"`

	Amqp AmqpConfig `yaml:"amqp_conf"`
}
