package main

import (
	clientapp "github.com/rellitelink/box/client"
	"github.com/rellitelink/box/config"
)

func main() {
	config.LoadConfig()
	clientapp.ClientClient()
}
