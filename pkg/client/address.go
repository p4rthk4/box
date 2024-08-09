// NOTE: This is client address file not server address file,
// 		 this address utils code only use in client package,
//       if use this code in server than write new code in
//		 server address file...

package client

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"unicode"
)

func getIPFromString(addr string) ([]net.IP, error) {
	ip := net.ParseIP(addr)
	if ip != nil {
		return []net.IP{ip}, nil
	}

	ips, err := net.LookupIP(addr)
	if err != nil {

		switch e := err.(type) {
		case *net.DNSError:
			if e.IsNotFound {
				return []net.IP{}, fmt.Errorf("no any ip records (A/AAAA) found of this %s domain, may be domain dose not exist", addr)
			}

		default:
			return nil, fmt.Errorf("Domain Name System (DNS) error when looking A or AAAA records for %s domain", addr)
		}

		return []net.IP{}, err
	}

	return ips, nil
}

func isIPv6(ip net.IP) bool {
	if ip != nil && ip.To4() == nil && ip.To16() != nil {
		return true
	}
	return false
}

func isValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	emailRegex := regexp.MustCompile(`^[\p{L}\p{N}\p{M}\p{S}\p{P}._%+\-]+@[\p{L}\p{N}.\-]+\.[\p{L}]{2,}$`)	
	return emailRegex.MatchString(email)
}

func getDomainFromEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if !isValidEmail(email) {
		return "", fmt.Errorf("invalid email address")
	}
	parts := strings.Split(email, "@")
	return parts[1], nil
}

func IsSMTPUTF8(email string) bool {
	for _, r := range email {
		if r > unicode.MaxASCII {
			return true // Contains Unicode characters
		}
	}
	return false // Contains only ASCII characters
}
