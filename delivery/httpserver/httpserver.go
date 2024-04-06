package httpserver

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"newsapp/config"
	"newsapp/delivery/httpserver/userhandler"
	"newsapp/logger"
	"newsapp/pkg/customcontext"
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
	s.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &customcontext.ApiContext{Context: c}

			return next(apiContext)
		}
	})
	s.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogRequestID:     true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var errMsg string
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.ZapLogger.Named("http-server").Info("request",
				zap.String("uri", v.URI),
				zap.String("referer", v.Referer),
				zap.Duration("latency", v.Latency),
				zap.String("user_agent", v.UserAgent),
				zap.Int("response_size", int(v.ResponseSize)),
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content_length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("error", errMsg),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

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
