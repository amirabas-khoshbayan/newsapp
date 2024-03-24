package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"newsapp/config"
)

type Server struct {
	config config.Config
	Echo   *echo.Echo
}

func New(config config.Config) Server {
	return Server{Echo: echo.New(), config: config}
}

func (s Server) Serve() {
	//Middleware
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(middleware.Recover())

	//Routs
	s.Echo.GET("/health-check", s.healthCheck)

	//start server
	address := fmt.Sprintf(":%d", s.config.Server.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Echo.Start(address); err != nil {
		fmt.Println("router start error", err)
	}

}
