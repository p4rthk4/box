// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"log"
)

type MailFwdWebhook struct {
	MailForward
}

func (mailFwd *MailFwdWebhook) Init() {
	log.Println("Init Webhook/HTTP Forward method")
}

func (mailFwd *MailFwdWebhook) ForwardMail(email Email) {
	log.Println("Mail Recive in HTTP")
}
