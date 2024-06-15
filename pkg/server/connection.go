// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"fmt"
	"net"

	"github.com/p4rthk4/u2smtp/pkg/logx"
	"github.com/p4rthk4/u2smtp/pkg/uid"
)

type Connection struct {
	conn          net.Conn
	uid           string
	serverLogger  *logx.Log
	logger        *logx.Log
	remoteAddress RemoteAddress
}

type RemoteAddress struct {
	ip         net.IP
	port       int
	isIPv6     bool
	ptrRecords []string
}

// handle new client connection
func HandleNewConnection(conn net.Conn, serverLogger *logx.Log) {
	connection := Connection{
		conn:         conn,
		serverLogger: serverLogger,
	}

	err := connection.init()
	if err {
		return
	}

	connection.handle()
}

// init client connection
// return true if error
func (conn *Connection) init() bool {

	uid, err := uid.GetNewId()
	if err != nil {
		conn.serverLogger.Error("generate email uid error: %v", err)
		return true
	}

	conn.uid = uid
	conn.logger = conn.serverLogger.GetNewWithPrefix(uid)
	conn.parseRemoteAddress()

	conn.logger.Info("client %s[%s]", conn.remoteAddress.GetPTR(), conn.remoteAddress.ip.String())

	return false // err!
}

// handle client connection
func (conn *Connection) handle() {
	for {
		n, err := fmt.Fprintf(conn.conn, "Hello!\n")
		if n < 1 {
			fmt.Printf("%d %s\n", n, err)
			conn.conn.Close()
			break
		}
	}
}

// parse client address(remote address)
func (conn *Connection) parseRemoteAddress() {

	ipAddr, err := net.ResolveTCPAddr(conn.conn.RemoteAddr().Network(), conn.conn.RemoteAddr().String()) // resolve / parse ip address
	if err != nil {
		fmt.Println(err)
	}

	conn.remoteAddress.port = ipAddr.Port

	if ipAddr.IP != nil && ipAddr.IP.To4() == nil && ipAddr.IP.To16() != nil { // if IPv4
		conn.remoteAddress.ip = ipAddr.IP.To16()
		conn.remoteAddress.isIPv6 = true
	} else if ipAddr.IP != nil && ipAddr.IP.To4() != nil { // if IPv4
		conn.remoteAddress.ip = ipAddr.IP.To4()
	} else {
		conn.remoteAddress.ip = ipAddr.IP
	}

	// get/lookup ptr records
	adds, err := net.LookupAddr(conn.remoteAddress.String())
	if err != nil {
		conn.logger.Warn("no PTR record or faild to find PTR records of %s", conn.remoteAddress.String())
		conn.remoteAddress.ptrRecords = nil
	} else {
		conn.remoteAddress.ptrRecords = adds
	}

}

// return first PTR records
func (ra *RemoteAddress) GetPTR() string {
	if len(ra.ptrRecords) > 0 {
		return ra.ptrRecords[0]
	} else {
		return "unknow"
	}
}

// get string of ip
func (ra *RemoteAddress) String() string {
	return ra.ip.String()
}
