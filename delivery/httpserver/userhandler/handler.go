package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/service/userservice"
)

type Handler struct {
	userSvc userservice.Service
}

func New(userSvc userservice.Service) Handler {
	return Handler{userSvc: userSvc}
}
func (h Handler) getUserList(c echo.Context) error {

	userList, err := h.userSvc.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users": userList,
	})
}
