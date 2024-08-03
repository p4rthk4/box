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
	Uid     string `yaml:"uid"`
	Success bool   `yaml:"success"`

	Tls       bool             `yaml:"tls"`
	TlsVerify bool             `yaml:"tls_verify"`
	PtrIP     ServerClientInfo `yaml:"ptr_ip"`

	Domain   string `yaml:"domain"`
	PtrMatch bool   `yaml:"ptr_match"`

	From       string   `yaml:"from"`
	Recipients []string `yaml:"recipients"`
	UseBdat    bool     `yaml:"use_bdat"`
	Data       string   `yaml:"data"`
	dataByte   []byte
}

type ServerClientInfo struct {
	ServerPtr string `yaml:"server_ptr"`
	ServerIP  string `yaml:"server_ip"`
	ClinetPtr string `yaml:"client_ptr"`
	ClientIP  string `yaml:"client_ip"`
}

func (e *Email) ToDocument() ([]byte, error) {
	e.Data = string(e.dataByte)
	data, err := yaml.Marshal(&e)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *Email) GetBytes() []byte {
	return e.dataByte
}

func (e *Email) ParseMail() (letters.Email, bool) {
	emailReader := bytes.NewReader([]byte(e.Data))
	email, err := letters.ParseEmail(emailReader)
	if err != nil {
		if config.ConfOpts.Dev {
			log.Println(err)
		}
		return letters.Email{}, false
	}

	return email, true
}
