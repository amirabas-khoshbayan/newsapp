package main

import (
	"fmt"
	"newsapp/adapter/redis"
	"newsapp/config"
	"newsapp/delivery/httpserver"
	"newsapp/logger"
	"newsapp/repository/mysql"
	"newsapp/repository/mysql/migrator"
	"newsapp/repository/mysql/mysqlnews"
	"newsapp/repository/mysql/mysqluser"
	"newsapp/repository/redis/redispublish"
	"newsapp/scheduler"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/newsservice"
	"newsapp/service/publishservice"
	"newsapp/service/userservice"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	appLabel := fmt.Sprintf("time = %s , Branch =  %s, AppVersion  %s", time.Now().String(), "master", config.AppVersion)
	fmt.Println(appLabel)

	cfg := config.GetConfig()
	fmt.Println(cfg)

	mgr := migrator.New(cfg.MySQL)
	mgr.Up()

	// init zap logger // TODO - Replace zap with new logger
	logger.Init(cfg.ZapLogger)

	//mongoConn := mongodb.New(cfg.MongoDB)
	mySqlConn := mysql.New(cfg.MySQL)
	redisAdapter := redis.New(cfg.Redis)
	redisRepo := redispublish.New(redisAdapter)

	//userMongo := mongodbuser.New(mongoConn)
	userMySql := mysqluser.New(mySqlConn)
	newsMySql := mysqlnews.New(mySqlConn)
	authSvc := authenticationservice.New(cfg.Auth)
	authorizeSvc := authorizationservice.New(mySqlConn)
	userSvc := userservice.New(userMySql, authSvc)
	newsSvc := newsservice.New(newsMySql)
	publishSvc := publishservice.New(cfg.PublishService, redisRepo, redisAdapter)

	done := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		sch := scheduler.New(cfg.Scheduler, publishSvc)

		wg.Add(1)
		sch.Start(done, &wg)
	}()

	server := httpserver.New(cfg, userSvc, newsSvc, publishSvc, authSvc, authorizeSvc)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
