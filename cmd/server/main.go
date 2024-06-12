// U2SMTP - server cmd
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package main

import (
	"fmt"

	conf "github.com/p4rthk4/u2smtp/pkg/config"
)

func main() {
	fmt.Println(conf.ConfOpts)
	conf.LoadConfig()
	fmt.Println(conf.ConfOpts)
}
