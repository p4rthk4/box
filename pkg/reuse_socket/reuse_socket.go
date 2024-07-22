// reusable socket (set socket option, it help to reuse address)
// only for linux

package reusesocket

import (
	"fmt"
	"net"
)

// reusable listen
func Listen(network string, address string) (l net.Listener, err error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return tpcListner(address)
	default:
		return nil, fmt.Errorf("reusesocket: unsupported network %q", network)
	}
}

// reusable packet listen
func ListenPacket(network string, address string) (l net.PacketConn, err error) {
	switch network {
	case "udp", "udp4", "udp6":
		return udpListner(address)
	default:
		return nil, fmt.Errorf("reusesocket: unsupported network %q", network)
	}
}
