package main

import (
	"fmt"
	"newsapp/config"
	"newsapp/delivery/httpserver"
	"os"
	"os/signal"
)

func main() {
	//get config
	cfg := config.GetConfig()
	fmt.Println(cfg)

	server := httpserver.New(cfg)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
