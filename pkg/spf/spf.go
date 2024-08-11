package spf

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// Result of SPF check
type Result string

// SPF results
const (
	None      = Result("NONE")
	Neutral   = Result("NEUTRAL")
	Pass      = Result("PASS")
	Fail      = Result("FAIL")
	Softfail  = Result("SOFTFAIL")
	TempError = Result("TEMPERROR")
	PermError = Result("PERMERROR")
)

// String representation of Result type
func (r Result) String() string {
	return string(r)
}

// IsSet returns true if Result var is set to some value
func (r Result) IsSet() bool {
	return string(r) != ""
}

type check struct {
	cnt int
}

func CheckHost(ip net.IP, domain, sender, helo string) Result {

	if sender == "" {
		sender = "postmaster@" + helo
	}

	c := check{
		cnt: 0,
	}

	return c.checkHost(ip, domain, sender)
}

func (c *check) checkHost(ip net.IP, domain, sender string) Result {
	defRes := None

	spf, r := LookupSPF(domain)
	if r.IsSet() {
		return r
	}
	// log.Println("\n\n", spf, "\n------------------------")
	terms, err := parseSPF(spf)
	if err != nil {
		return TempError
	}

	for _, term := range terms {
		switch t := term.(type) {
		case directive:
			d := t
			var r Result
			switch d.mechanism {
			case "a":
				dom := d.domain(domain)
				r = c.check(ip, dom, d.cidr(), d.qualifier)

			case "mx":
				dom := d.domain(domain)
				r = c.checkMX(ip, dom, d.cidr(), d.qualifier)

			case "include":
				dom := d.domain(domain)
				r = c.checkHost(ip, dom, sender)
				// TODO: r is not as returned, see page 22

			case "ptr":
				// not recommended for use, but must be implemented
				dom := d.domain(domain)
				r = c.checkPTR(ip, dom, d.qualifier)

			case "ip4":
				if ip.To4() != nil { // check only if ip is IPv4 address
					r = checkIP(ip, d.param, d.qualifier)
				}

			case "ip6":
				if ip.To4() == nil { // check only if ip is IPv6 address
					r = checkIP(ip, d.param, d.qualifier)
				}

			case "all":
				return evalQualifier(d.qualifier)

			case "exists":
				dom, res := c.macro(d.param, ip, domain, sender, "")
				if res == PermError {
					return PermError
				}
				ips, _ := lookupA(dom)
				if len(ips) > 0 {
					return Pass
				}
			}

			// check result from mechanism
			switch r {
			case Pass, PermError:
				return r
			case TempError:
				defRes = r
			default:
				defRes = Neutral
			}

		case modifier:
			switch t.name {
			case "redirect":
				return c.checkHost(ip, t.value, sender)
			case "exp":
			default:
			}
			//	something to do with modifiers
		}
	}

	return defRes
}

// check record of specific domain for IP, return true if match
func (c *check) check(ip net.IP, domain, cidr, qualifier string) Result {
	if c.cnt == 10 {
		return PermError
	}

	c.cnt++
	var ips []net.IP
	var err error

	if ip.To4() == nil {
		ips, err = lookupAAAA(domain)
	} else {
		ips, err = lookupA(domain)
	}

	if err != nil {
		return TempError
	}

	for _, a := range ips {
		if r := checkIP(ip, a.String()+cidr, qualifier); r != Neutral {
			return r
		}
	}
	return Neutral
}

func checkIP(ip net.IP, ipstr, qualifier string) Result {
	_, ips, err := net.ParseCIDR(ipstr)
	if err == nil {
		//log.Println("Check range", ips.String())
		if ips.Contains(ip) {
			return evalQualifier(qualifier)
		}
	} else {
		ipaddr := net.ParseIP(ipstr)
		if ip.Equal(ipaddr) {
			return evalQualifier(qualifier)
		}
	}
	return Neutral
}

// evalQualifier returns Pass if qualifier is + or "" or other spf results accordingly
func evalQualifier(q string) Result {
	switch q {
	case "~":
		return Softfail
	case "-":
		return Fail
	case "?":
		return Neutral
	default:
		return Pass
	}
}

// checkA record of specific domain for IP, return true if match
func (c *check) checkMX(ip net.IP, domain, cidr, qualifier string) Result {
	defRes := None

	mxs, err := lookupMX(domain)
	if err != nil {
		return TempError
	}

	for _, mx := range mxs {
		r := c.check(ip, mx, cidr, qualifier)
		switch r {
		case Pass, PermError, Fail:
			return r
		case TempError:
			// on TempError continue to check other mx records, but remember temperror (dns unavailable)
			defRes = r
		}
	}
	return defRes
}

// checkPTR match
func (c *check) checkPTR(ip net.IP, domain, qualifier string) Result {
	defRes := None

	hosts, err := lookupPTR(ip)
	if err != nil {
		return TempError
	}

	for _, h := range hosts {
		h = strings.TrimRight(h, ".")
		domain = strings.TrimRight(domain, ".")
		fmt.Println(h, domain)
		if strings.Contains(h, domain) {
			s := strings.Split(h, domain)
			fmt.Println("split text", s, len(s))
		}
	}

	return defRes
}

type modifier struct {
	name  string
	value string
}

type directive struct {
	qualifier string
	mechanism string
	param     string
}

// domain returns default domain (param) or domain specified in spf record after : sign
func (d directive) domain(domain string) string {
	if d.param != "" {
		parts := strings.SplitN(d.param, "/", 2)
		return parts[0]
	}
	return domain
}

func (d directive) cidr() string {
	n := strings.Index(d.param, "/")
	if n != -1 {
		return d.param[n:]
	}
	return ""
}

// directive
// qualifier
// mechanism
// = *( 1*SP ( directive / modifier ) )
// = [ qualifier ] mechanism
// = "+" / "-" / "?" / "~"
// = ( all / include / a / mx / ptr / ip4 / ip6 / exists )

// ParseSPF record and return slice with directives and modifiers
func parseSPF(spf string) ([]interface{}, error) {
	spf = strings.TrimSpace(strings.TrimPrefix(spf, "v=spf1"))

	var terms []interface{}
	parts := strings.Fields(spf)
	for _, t := range parts {
		mDirective, err := regexp.Compile(`^(\+|\-|\?|\~)?(all|include|a|mx|ptr|ip4|ip6|exists):?(.*)$`)
		if err != nil {
			return nil, err
		}

		dirMatch := mDirective.FindStringSubmatch(t)
		if len(dirMatch) > 0 {
			terms = append(terms, directive{
				qualifier: dirMatch[1],
				mechanism: dirMatch[2],
				param:     dirMatch[3],
			})
			continue
		} else {
			mModifier, err := regexp.Compile(`^([a-z0-9\-\_\.]+)=(.*)$`)
			if err != nil {
				return nil, err
			}

			modMatch := mModifier.FindStringSubmatch(t)
			if len(modMatch) > 0 {
				terms = append(terms, modifier{
					name:  modMatch[1],
					value: modMatch[2],
				})
			}
		}
	}
	return terms, nil
}

func LookupSPF(domain string) (spf string, r Result) {
	txts, err := net.LookupTXT(domain)
	if err != nil {
		return "", TempError
	}

	var spfs []string
	for _, txt := range txts {
		txt = strings.ToLower(txt)
		if txt == "v=spf1" || strings.HasPrefix(txt, "v=spf1 ") {
			spfs = append(spfs, txt)
		}
	}

	switch len(spfs) {
	case 0:
		return "", None
	case 1:
		return spfs[0], Result("")
	default:
		return "", PermError
	}
}

func lookupA(d string) ([]net.IP, error) {
	var ips []net.IP

	r, err := net.LookupIP(d)
	if err != nil {
		return ips, err
	}

	for _, ip := range r {
		if ip == nil || ip.To4() != nil || ip.To16() == nil {
			ips = append(ips, ip)
		}
	}

	return ips, nil
}

func lookupAAAA(d string) ([]net.IP, error) {
	var ips []net.IP

	r, err := net.LookupIP(d)
	if err != nil {
		return ips, err
	}

	for _, ip := range r {
		if ip != nil && ip.To4() == nil && ip.To16() != nil {
			ips = append(ips, ip)
		}
	}

	return ips, nil
}

func lookupMX(d string) ([]string, error) {
	var mxs []string

	mxList, err := net.LookupMX(d)
	if err != nil {
		return mxs, err
	}

	for _, mx := range mxList {
		mxs = append(mxs, mx.Host)
	}

	return mxs, nil
}

func lookupPTR(ip net.IP) ([]string, error) {
	var hosts []string

	r, err := net.LookupAddr(ip.String())
	if err != nil {
		return hosts, err
	}

	return r, nil
}

//   v=spf1
//
//    550 5.7.1 SPF MAIL FROM check failed:
//    550 5.7.1 The domain example.com explains:
//    550 5.7.1 Please see http://www.example.com/mailpolicy.html

// Received-SPF:
// Authentication-Results:
