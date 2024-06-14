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
	}
}
