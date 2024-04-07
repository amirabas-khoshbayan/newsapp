package main

import (
	"fmt"
	"newsapp/config"
	"newsapp/delivery/httpserver"
	"newsapp/logger"
	"newsapp/repository/mysql"
	"newsapp/repository/mysql/migrator"
	"newsapp/repository/mysql/mysqlnews"
	"newsapp/repository/mysql/mysqluser"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/newsservice"
	"newsapp/service/userservice"
	"os"
	"os/signal"
	"time"
)

func main() {
	appLabel := fmt.Sprintf("time = %s , Branch =  %s, AppVersion  %s", time.Now().String(), "master", config.AppVersion)
	fmt.Println(appLabel)

	cfg := config.GetConfig()
	fmt.Println(cfg)

	mgr := migrator.New(cfg.MySQL)
	mgr.Up()

	// init zap logger
	logger.Init(cfg.ZapLogger)

	//mongoConn := mongodb.New(cfg.MongoDB)
	mySqlConn := mysql.New(cfg.MySQL)

	//userMongo := mongodbuser.New(mongoConn)
	userMySql := mysqluser.New(mySqlConn)
	newsMySql := mysqlnews.New(mySqlConn)
	authSvc := authenticationservice.New(cfg.Auth)
	userSvc := userservice.New(userMySql, authSvc)
	newsSvc := newsservice.New(newsMySql)
	authorizeSvc := authorizationservice.New(mySqlConn)

	server := httpserver.New(cfg, userSvc, newsSvc, authSvc, authorizeSvc)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
