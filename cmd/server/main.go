package main

import (
	"fmt"

	"github.com/p4rthk4/u2smtp/server"
	"github.com/p4rthk4/u2smtp/config"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.ConfOpts)

	serverapp.StartServer()
}
