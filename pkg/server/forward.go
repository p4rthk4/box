// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"github.com/p4rthk4/u2smtp/pkg/config"
)

var mailFwd MailForward = nil // it is hold MailForward struct by config

type MailForward interface {
	Init()
	ForwardMail(Email)
}

func SetMailFwdMethod() {
	switch config.ConfOpts.Forward {
	case "webhook", "http":
		mailFwd = &MailFwdWebhook{}
		mailFwd.Init()
	case "amqp":
		mailFwd = &MailFwdAmqp{}
		mailFwd.Init()
	case "redis":
		mailFwd = &MailFwdRedis{}
		mailFwd.Init()
	default:
		mailFwd = &MailFwdNone{}
		mailFwd.Init()
	}
}

type MailFwdNone struct {
	MailForward
}

func (mailFwd *MailFwdNone) Init()                   {}
func (mailFwd *MailFwdNone) ForwardMail(email Email) {}
