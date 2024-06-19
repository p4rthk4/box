// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import "log"

type MailFwdAmqp struct {
	MailForward
}

func (mailFwd *MailFwdAmqp) Init() {
	log.Println("Init AMQP Forward method")
}

func (mailFwd *MailFwdAmqp) ForwardMail(email Email) {

}
