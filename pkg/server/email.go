// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"gopkg.in/yaml.v3"
)

type Email struct {
	Uid string `yaml:"uid"`

	Domain     string   `yaml:"domain"`
	From       string   `yaml:"from"`
	Recipients []string `yaml:"recipients"`
	Data       string   `yaml:"data"`
}

func (e *Email) ToDocument() ([]byte, error) {
	data, err := yaml.Marshal(&e)
	if err != nil {
		return nil, err
	}

	return data, nil
}
