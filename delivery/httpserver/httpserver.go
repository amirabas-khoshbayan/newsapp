package httpserver

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"newsapp/config"
	"newsapp/delivery/httpserver/userhandler"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/userservice"
	"newsapp/validator/customvalidator"
)

type Server struct {
	config       config.Config
	userHandler  userhandler.Handler
	authorizeSvc authorizationservice.Service
	authSvc      authenticationservice.Service
	Echo         *echo.Echo
}

func New(config config.Config, userSvc userservice.Service, authSvc authenticationservice.Service, authorizeSvc authorizationservice.Service) Server {
	return Server{Echo: echo.New(), config: config, userHandler: userhandler.New(userSvc, authSvc, authorizeSvc)}
}

func (s Server) Serve() {
	//Middleware
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(middleware.Recover())

	//Routs
	s.Echo.GET("/health-check", s.healthCheck)
	s.userHandler.SetUserRoutes(s.Echo)

	if config.AppConfig.HttpServer.UseCustomValidator {
		s.Echo.Validator = &customvalidator.Custom{Validator: validator.New()}
	}

	//start server
	address := fmt.Sprintf(":%d", s.config.HttpServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Echo.Start(address); err != nil {
		fmt.Println("router start error", err)
	}

}
