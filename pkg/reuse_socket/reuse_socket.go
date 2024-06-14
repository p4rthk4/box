// U2SMTP - reuse socket (set socket option, it help to reuse address)
// user this socket or listner for server.
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package reusesocket

import (
	"fmt"
	"net"
)

func Listen(network string, address string) (l net.Listener, err error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return tpcListner(address)
	default:
		return nil, fmt.Errorf("reusesocket: unsupported network %q", network)
	}
}

func ListenPacket(network string, address string) (l net.PacketConn, err error) {
	switch network {
	case "udp", "udp4", "udp6":
		return udpListner(address)
	default:
		return nil, fmt.Errorf("reusesocket: unsupported network %q", network)
	}
}
