// U2SMTP - client cmd
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package client

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
)

func ClientTmp() {

	fmt.Println("Client run...")
	serverAddr := getMXRecord("gmail.com")

	fmt.Println(serverAddr)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:25", serverAddr))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := make([]byte, 1024)
	l, err := conn.Read(b)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	_ = l
	fmt.Println("Start: \n", string(b[:l]))

	_, err = conn.Write([]byte("EHLO localhost\r\n"))
	if (err != nil) {
		fmt.Println(err)
		return
	}

	b = make([]byte, 1024)
	l, err = conn.Read(b)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	_ = l
	fmt.Println("EHLO: \n", string(b[:l]))

	_, err = conn.Write([]byte("STARTTLS\r\n"))
	if (err != nil) {
		fmt.Println(err)
		return
	}

	b = make([]byte, 1024)
	l, err = conn.Read(b)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	_ = l
	fmt.Println("STARTTLS: \n", string(b[:l]))

	tlsConfig := &tls.Config{
        // InsecureSkipVerify: true, // WARNING: Only for testing. Use proper verification in production.
        ServerName: serverAddr,
    }

	tlsConn := tls.Client(conn, tlsConfig)

	err = tlsConn.Handshake()
    if err != nil {
        fmt.Println("TLS handshake failed:", err)
        return
    }

    // Now tlsConn can be used as a secure connection
    fmt.Println("TLS connection established")

	_, err = tlsConn.Write([]byte("MAIL FROM: <hello@parthka.com>\r\n"))
	if (err != nil) {
		fmt.Println(err)
		return
	}

	b = make([]byte, 1024)
	l, err = tlsConn.Read(b)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	_ = l
	fmt.Println("HANDSHACK: \n", string(b[:l]))



}

func getMXRecord(domain string) string {

	mxs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("loock up err:", err)
		os.Exit(1)
	}

	mxr := getRand(mxs)

	return mxr.Host
}

func getRand[V any](arr []V) V {
	randomIndex := rand.Intn(len(arr))
	randomElement := arr[randomIndex]
	return randomElement
}
