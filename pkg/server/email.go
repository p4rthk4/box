// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

type Email struct {
	Domain     string
	From       string
	Recipients []string
	Data       []byte
	Uid        string
}
