// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"net"
)

// remote or local address
type RoLAddress struct {
	ip         net.IP
	port       int
	isIPv6     bool
	ptrRecords []string
}

func (ra *RoLAddress) SetAddress(network string, address string) bool {
	ipAddr, err := net.ResolveTCPAddr(network, address) // resolve / parse ip address
	if err != nil {
		ra.ptrRecords = nil
		return false
	}

	ra.port = ipAddr.Port

	if ipAddr.IP != nil && ipAddr.IP.To4() == nil && ipAddr.IP.To16() != nil { // if IPv6
		ra.ip = ipAddr.IP.To16()
		ra.isIPv6 = true
	} else if ipAddr.IP != nil && ipAddr.IP.To4() != nil { // if IPv4
		ra.ip = ipAddr.IP.To4()
	} else {
		ra.ip = ipAddr.IP
	}

	// get/lookup ptr records
	adds, err := net.LookupAddr(ra.String())
	if err != nil {
		ra.ptrRecords = nil
		return false
	} else {
		ra.ptrRecords = adds
	}

	return true
}

// return first PTR records
func (ra *RoLAddress) GetPTR() string {
	if len(ra.ptrRecords) > 0 {
		ptr := ra.ptrRecords[0]

		// if name end with . than remove this dot
		lastChar := ptr[len(ptr)-1:]
		if lastChar == "." {
			ptr = ptr[:len(ptr)-1]
		}

		return ptr
	} else {
		return "unknow"
	}
}

func (ra *RoLAddress) String() string {
	return ra.ip.String()
}
