package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")
	userGroup.GET("/get/list", nil).Name = "userList"
	e.Reverse("userList")
}
