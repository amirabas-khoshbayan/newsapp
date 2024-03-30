package userhandler

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"newsapp/entity"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	jwtConfig := echojwt.Config{
		SigningKey: []byte(h.authSvc.Config.SignKey),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &entity.JwtClaims{}
		},
	}
	userGroup.POST("/login", h.loginUser)
	userGroup.GET("/get/list", h.getUserList).Name = "userList"
	userGroup.POST("/create/new", h.createNewUser, echojwt.WithConfig(jwtConfig)).Name = "createNewUser"
}
