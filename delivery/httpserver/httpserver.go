package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"newsapp/config"
	"newsapp/service/userservice"
)

type Server struct {
	config  config.Config
	userSvc userservice.Service
	Echo    *echo.Echo
}

func New(config config.Config, userSvc userservice.Service) Server {
	return Server{Echo: echo.New(), config: config, userSvc: userSvc}
}

func (s Server) Serve() {
	//Middleware
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(middleware.Recover())

	//Routs
	s.Echo.GET("/health-check", s.healthCheck)

	//start server
	address := fmt.Sprintf(":%d", s.config.HttpServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Echo.Start(address); err != nil {
		fmt.Println("router start error", err)
	}

}
