// U2SMTP - reuse socket (set socket option, it help to reuse address)
// user this socket or listner for server.
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package reusesocket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func tpcListner(address string) (l net.Listener, err error) {
	var sockType int
	var sockFd int
	var sockAddr syscall.Sockaddr
	var file *os.File

	sockAddr, sockType = getSocketAddress("tcp", address)

	syscall.ForkLock.RLock()
	if sockFd, err = syscall.Socket(sockType, syscall.SOCK_STREAM, syscall.IPPROTO_TCP); err != nil {
		syscall.ForkLock.RUnlock()

		return nil, err
	}
	syscall.ForkLock.RUnlock()

	if err = syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		syscall.Close(sockFd)
		return nil, err
	}

	if err = syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, 0x0F, 1); err != nil { // 0x0F or 15 is reuse port option value
		syscall.Close(sockFd)
		return nil, err
	}

	if err = syscall.Bind(sockFd, sockAddr); err != nil {
		syscall.Close(sockFd)
		return nil, err
	}

	if err = syscall.Listen(sockFd, maxListenerBacklogFromSoMaxConn()); err != nil {
		syscall.Close(sockFd)
		return nil, err
	}

	socketFileName := getSocketFileName("tcp", address)
	fmt.Println(socketFileName)
	file = os.NewFile(uintptr(sockFd), socketFileName)
	if l, err = net.FileListener(file); err != nil {
		file.Close()
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return l, err

}

func maxListenerBacklogFromSoMaxConn() int {
	fd, err := os.Open("/proc/sys/net/core/somaxconn")
	if err != nil {
		return syscall.SOMAXCONN
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	line, err := rd.ReadString('\n')
	if err != nil {
		return syscall.SOMAXCONN
	}

	f := strings.Fields(line)
	if len(f) < 1 {
		return syscall.SOMAXCONN
	}

	n, err := strconv.Atoi(f[0])
	if err != nil || n == 0 {
		return syscall.SOMAXCONN
	}

	if n > 1<<16-1 {
		n = 1<<16 - 1
	}

	return n
}
