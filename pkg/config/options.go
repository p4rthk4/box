// U2SMTP - config
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package config

type ConfigsOptions struct {
	Name       string        `yaml:"name"`
	HostName   string        `yaml:"hostname"`
	Server     ServerOptions `yaml:"server"`
	MaxClients int           `yaml:"max_clients"`

	LogDirPath  string `yaml:"log_dir"`
	LogFilePath string `yaml:"log_file"`
	Dev         bool   `yaml:"dev"`

	ClientGreet string `yaml:"client_greet"`
	ClientByyy  string `yaml:"client_byyy"`

	MaxRecipients int `yaml:"max_recipients"`

	Forward string `yaml:"forward"` // 'http' or 'amqp'
}

type ServerOptions struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	HostIPv6    string `yaml:"host_ipv6"`
	PortIPv6    int    `yaml:"port_ipv6"`
	IPv6Disable bool   `yaml:"IPv6_disable"`
}
