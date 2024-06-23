// U2SMTP - server cmd
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package main

import (
	"fmt"

	"github.com/p4rthk4/u2smtp/pkg/app"
	"github.com/p4rthk4/u2smtp/pkg/config"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.ConfOpts)

	app.StartApp()
}
