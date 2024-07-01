// U2SMTP - client cmd
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package main

import smtpclient "github.com/p4rthk4/u2smtp/pkg/client"

func main() {
	clinet := smtpclient.NewClinet()
	clinet.SetHost("lolkongd.com")

	mailContent := []byte{}

	clinet.SendMail(mailContent)
}
