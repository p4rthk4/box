// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"fmt"
	"net"
)

type Connection struct {
	conn net.Conn
}

func HandleNewConnection(conn net.Conn) {
	connection := Connection{
		conn: conn,
	}

	connection.handle()
}

func (connection *Connection) handle() {
	fmt.Printf("%s	 %s\n", connection.conn.LocalAddr().String(), connection.conn.RemoteAddr().String())
	for {
		n, err := fmt.Fprintf(connection.conn, "Hello!\n")
		if n < 1 {
			fmt.Printf("%d %s\n", n, err)
			connection.conn.Close()
			break
		}
	}
}
