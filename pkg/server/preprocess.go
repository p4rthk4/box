package server

import (
	"fmt"

	"github.com/p4rthk4/u2smtp/pkg/config"
)

var greetReplyMessage []string = []string{} // greet reply message for esmtp/ehlo

// smtp server pre-process like EHLO
// greet reply etc
func smtpServerPreProcess() {
	if config.ConfOpts.ESMTP.Enable {
		reply := []string{
			"PIPELINING",
			"8BITMIME",
			// "ENHANCEDSTATUSCODES", // TODO
			"CHUNKING",
		}

		if config.ConfOpts.ESMTP.Tls {
			reply = append(reply, "STARTTLS")
		}
		if config.ConfOpts.ESMTP.Tls && config.ConfOpts.ESMTP.RequireTLS {
			reply = append(reply, "REQUIRETLS")
		}
		if config.ConfOpts.ESMTP.Utf8 {
			reply = append(reply, "SMTPUTF8")
		}
		if config.ConfOpts.ESMTP.BinaryMime {
			reply = append(reply, "BINARYMIME")
		}
		// TODO: DSN
		// if c.server.EnableDSN {
		// 	reply = append(reply, "DSN")
		// }
		if config.ConfOpts.ESMTP.MessageSize > 0 {
			reply = append(reply, fmt.Sprintf("SIZE %v", config.ConfOpts.ESMTP.MessageSize))
		} else {
			reply = append(reply, "SIZE")
		}
		if config.ConfOpts.MaxRecipients > 0 {
			reply = append(reply, fmt.Sprintf("LIMITS RCPTMAX=%v", config.ConfOpts.MaxRecipients))
		}
		
		greetReplyMessage = reply
	}
}
