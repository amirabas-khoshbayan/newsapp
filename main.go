package main

import (
	"fmt"
	"newsapp/config"
	"newsapp/delivery/httpserver"
)

func main() {
	//get config
	cfg := config.GetConfig()
	fmt.Println(cfg)

	server := httpserver.New(cfg)
	go func() {
		server.Serve()
	}()
}
