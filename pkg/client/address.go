// NOTE: This is client address file not server address file,
// 		 this address utils code only use in client package,
//       if use this code in server than write new code in
//		 server address file...

package client

import (
	"fmt"
	"net"
)

type IPAddress struct {
	ip     net.IP
	isIPv6 bool
}

type ServerAddress struct {
	name string
	ips  []IPAddress
}

func getServerAddress(addr string) (ServerAddress, error) {
	ip := net.ParseIP(addr)
	if ip != nil {
		ipAddr := getIPAddressFromIP(ip)
		return ServerAddress{
			ips:  []IPAddress{ipAddr},
			name: "",
		}, nil
	}

	ips, err := net.LookupIP(addr)
	if err != nil {
		return ServerAddress{}, fmt.Errorf("host IP loopup faild")
	}

	serverAddr := ServerAddress{
		name: addr,
		ips:  []IPAddress{},
	}
	for _, ip := range ips {
		serverAddr.ips = append(serverAddr.ips, getIPAddressFromIP(ip))
	}

	return serverAddr, nil
}

func getIPAddressFromIP(ip net.IP) IPAddress {
	if ip != nil && ip.To4() == nil && ip.To16() != nil {
		return IPAddress{
			ip:     ip,
			isIPv6: true,
		}
	} else {
		return IPAddress{
			ip:     ip,
			isIPv6: false,
		}
	}
}
