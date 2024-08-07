package server

import "time"

var config *SMTPConfig = nil

type SMTPConfig struct {
	Name       string
	HostName   string
	MaxClients int

	ESMTP ESMTPOptions

	SpfCheck bool

	LogDirPath  string
	LogFilePath string

	ClientGreet string
	ClientByyy  string

	MaxRecipients     int
	CheckMailBoxExist bool

	TlsConfig TlsKeyCert

	Timeout time.Duration // read timeout
	Dev     bool
}

type ESMTPOptions struct {
	Enable      bool
	Tls         bool
	Utf8        bool
	BinaryMime  bool
	MessageSize int
}

type TlsKeyCert struct {
	Key  string
	Cert string
}

// This is globle and overridable
// and give all option because not
// verify or default blank option
func SetConfig(conf SMTPConfig) {
	config = &conf
}
