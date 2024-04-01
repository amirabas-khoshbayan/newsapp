package main

import (
	"fmt"
	"newsapp/config"
	"newsapp/delivery/httpserver"
	"newsapp/repository/mongodb"
	"newsapp/repository/mongodb/mongodbuser"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/userservice"
	"os"
	"os/signal"
)

func main() {

	cfg := config.GetConfig()
	fmt.Println(cfg)

	mongoConn := mongodb.New(cfg.MongoDB)

	userMongo := mongodbuser.New(mongoConn)
	authSvc := authenticationservice.New(cfg.Auth)
	userSvc := userservice.New(userMongo, authSvc)
	authorizeSvc := authorizationservice.New(mongoConn)

	server := httpserver.New(cfg, userSvc, authSvc, authorizeSvc)

	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
