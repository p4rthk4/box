// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"bytes"
	"log"

	"github.com/mnako/letters"
	"github.com/p4rthk4/u2smtp/pkg/config"
	"gopkg.in/yaml.v3"
)

// recive mail structure
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

func (e *Email) ParseMail() (letters.Email, bool) {
	emailReader := bytes.NewReader([]byte(e.Data))
	email, err := letters.ParseEmail(emailReader)
	if err != nil {
		if config.ConfOpts.Dev {
			log.Println(err)
		}
		return  letters.Email{}, false
	}

	return email, true
}
