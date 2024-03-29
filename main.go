package main

import (
	"fmt"
	"newsapp/config"
	"newsapp/delivery/httpserver"
	"newsapp/repository/mongodb"
	"newsapp/repository/mongodb/mongodbuser"
	"newsapp/service/authservice"
	"newsapp/service/userservice"
	"os"
	"os/signal"
)

func main() {

	cfg := config.GetConfig()
	fmt.Println(cfg)

	mongoConn := mongodb.New(cfg.MongoDB)

	userMongo := mongodbuser.New(mongoConn)
	authSvc := authservice.New(cfg.Auth)
	userSvc := userservice.New(userMongo, authSvc)

	server := httpserver.New(cfg, userSvc, authSvc)

	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
