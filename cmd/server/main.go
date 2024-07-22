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
