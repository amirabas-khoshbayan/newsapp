package publishhandler

import (
	"github.com/labstack/echo/v4"
	"newsapp/delivery/httpserver/middleware"
	"newsapp/entity"
)

func (h Handler) SetPublishRoutes(e *echo.Echo) {
	publishGroup := e.Group("/publish")

	publishGroup.PUT("/submit", h.publishNews, middleware.Auth(h.authSvc),
		middleware.CheckPermissions(h.authSvc, h.authorizeSvc, entity.OwnerRole)).Name = "publishNews"
}
