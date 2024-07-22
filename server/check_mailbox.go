package serverapp

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
