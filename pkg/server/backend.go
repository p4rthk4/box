// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

var mailFwd ForwardBackend = nil // it is hold MailForward struct by config

type ForwardBackend interface {
	Init()
	ForwardMail(Email)
	ExistMailBox(string) bool
}

func SetMailFwdMethod(backend ForwardBackend) {
	mailFwd = backend
	mailFwd.Init()
}
