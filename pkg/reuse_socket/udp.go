// U2SMTP - reuse socket (set socket option, it help to reuse address)
// user this socket or listner for server.
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package reusesocket

import (
	"net"
	"os"
	"syscall"
)

func udpListner(address string) (l net.PacketConn, err error) {
	var sockType int
	var sockFd int
	var sockAddr syscall.Sockaddr
	var file *os.File

	sockAddr, sockType = getSocketAddress("tcp", address)

	syscall.ForkLock.RLock()
	sockFd, err = syscall.Socket(sockType, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err == nil {
		syscall.CloseOnExec(sockFd)
	}
	syscall.ForkLock.RUnlock()
	if err != nil {
		syscall.Close(sockFd)
		return nil, err
	}

	defer func() {
		if err != nil {
			syscall.Close(sockFd)
		}
	}()

	if err = syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		return nil, err
	}

	if err = syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, 0x0F, 1); err != nil { // 0x0F or 15 is reuse port option value
		return nil, err
	}

	if err = syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1); err != nil {
		return nil, err
	}

	if err = syscall.Bind(sockFd, sockAddr); err != nil {
		return nil, err
	}

	socketFileName := getSocketFileName("udp", address)
	file = os.NewFile(uintptr(sockFd), socketFileName)
	l, err = net.FilePacketConn(file)
	if err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return l, err

}
