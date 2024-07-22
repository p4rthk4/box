package reusesocket

import (
	"fmt"
	"os"
)

func getSocketFileName(network, addr string) string {
	return fmt.Sprintf("u2smtp_socket_%d_%s_%s", os.Getpid(), network, addr)
}
