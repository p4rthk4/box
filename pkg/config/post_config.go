// U2SMTP - config
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package config

import (
	"os"

	"github.com/p4rthk4/u2smtp/pkg/logx"
)

func postConfigAction() {

	// make log dir if not exists
	_, err := os.Stat(ConfOpts.LogDirPath)
	// fmt.Println(err)
	if err != nil {

		if os.IsNotExist(err) {
			err := os.MkdirAll(ConfOpts.LogDirPath, 0755)
			if err != nil {
				logx.LogError(err.Error(), err)
				os.Exit(1)
			}
		} else {
			logx.LogError(err.Error(), err)
			os.Exit(1)
		}
	}

}
