package main

import (
	"fmt"
	"web-api/config"
	"web-api/delivery/httpserver"
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
