package newshandler

import (
	"github.com/labstack/echo/v4"
	"newsapp/delivery/httpserver/middleware"
	"newsapp/entity"
)

func (h Handler) SetNewsRoutes(e *echo.Echo) {
	newsGroup := e.Group("/news")

	newsGroup.POST("/create", h.createNews, middleware.Auth(h.authSvc),
		middleware.CheckPermissions(h.authSvc, h.authorizeSvc, entity.AdminRole)).Name = "createNews"
	newsGroup.GET("/view/:id", h.getNews).Name = "getNews"
	newsGroup.GET("/list", h.getNewsList).Name = "getNewsList"

	// TODO - add  Edit / Delete routes
}
