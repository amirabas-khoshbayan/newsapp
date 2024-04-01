package userhandler

import (
	"github.com/labstack/echo/v4"
	"newsapp/delivery/httpserver/middleware"
	"newsapp/entity"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.POST("/login", h.loginUser)
	userGroup.GET("/get/list", h.getUserList).Name = "userList"
	userGroup.POST("/create", h.createNewUser).Name = "createNewUser"
	userGroup.PUT("/give-admin-role/:id", h.giveAdminRole, middleware.Auth(h.authSvc), middleware.CheckPermissions(h.authSvc, h.authorizeSvc, entity.OwnerRole)).Name = "giveAdminRole"
}
