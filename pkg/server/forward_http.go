// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import "log"

type MailFwdHttp struct {
	MailForward
}

func (mailFwd *MailFwdHttp) Init() {
	log.Println("Init HTTP Forward method")
}

func (mailFwd *MailFwdHttp) ForwardMail(email Email) {
	
}
