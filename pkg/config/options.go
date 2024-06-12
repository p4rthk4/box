// U2SMTP - config
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package config

type ConfigsOptions struct {
	Name     string        `yaml:"name"`
	HostName string        `yaml:"hostname"`
	Server   ServerOptions `yaml:"server"`
}

type ServerOptions struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
