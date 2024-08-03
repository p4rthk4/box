package server

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/p4rthk4/u2smtp/pkg/config"
	"github.com/p4rthk4/u2smtp/pkg/logx"
)

var greetReplyMessage []string = []string{} // greet reply message for esmtp/ehlo
var tlsConfig	 *tls.Config

// smtp server pre-process like EHLO
// greet reply, tls config etc
func smtpServerPreProcess(l *logx.Log) {
	// greet reply
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

	// tls config
	if config.ConfOpts.ESMTP.Enable && config.ConfOpts.ESMTP.Tls {
		cert, err := tls.LoadX509KeyPair(config.ConfOpts.Tls.Cert, config.ConfOpts.Tls.Key)
		if err != nil {
			l.Error("Error loading certificates: %v", err)
			os.Exit(1)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS13,
		}
	}
}
