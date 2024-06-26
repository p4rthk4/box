// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package app

import (
	"strings"
)

var tmpExistMailBox []string = []string{
	"hello@parthka.dev",
	"iamlolg@fs.com",
	"pka@ds.com",
	"hello@raviitaliya.info",
}

func checkMaiboxFromRcpt(rcpt string) bool {
	rcpt = strings.ToLower(rcpt)
	return contains(tmpExistMailBox, rcpt)
}

func contains(slice []string, element string) bool {
    for _, value := range slice {
        if value == element {
            return true
        }
    }
    return false
}
