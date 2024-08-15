package main

import (
	"fmt"
	"os"

	clientapp "github.com/rellitelink/box/client"
	"github.com/rellitelink/box/config"
	serverapp "github.com/rellitelink/box/server"
)


func main() {
	mode := os.Getenv("MODE")
	verbose := os.Getenv("VERBOSE")

	config.LoadConfig()
	

	if verbose == "true" {
		fmt.Println(config.ConfOpts)
	}

	if mode != "client" {
		serverapp.StartServer()
	} else {
		clientapp.StartClient()
	}
}