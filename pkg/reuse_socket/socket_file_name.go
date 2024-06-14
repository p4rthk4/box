// U2SMTP - reuse socket (set socket option, it help to reuse address)
// user this socket or listner for server.
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package reusesocket

import (
	"fmt"
	"os"
)

func getSocketFileName(network, addr string) string {
	return fmt.Sprintf("u2smtp_socket_%d_%s_%s", os.Getpid(), network, addr)
}
