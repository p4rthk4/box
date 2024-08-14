package main

import (
	"fmt"

	"github.com/rellitelink/box/config"
	serverapp "github.com/rellitelink/box/server"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.ConfOpts)

	serverapp.StartServer()
}
