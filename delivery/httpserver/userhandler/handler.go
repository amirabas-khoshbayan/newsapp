package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/param/userparam"
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
func (h Handler) createNewUser(c echo.Context) error {
	var req userparam.CreateNewUserRequest
	var res userparam.CreateNewUserResponse
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res, err := h.userSvc.CreateNewUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
