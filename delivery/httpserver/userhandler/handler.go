package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/param/userparam"
	"newsapp/service/authservice"
	"newsapp/service/userservice"
)

type Handler struct {
	userSvc userservice.Service
	authSvc authservice.Service
}

func New(userSvc userservice.Service, authSvc authservice.Service) Handler {
	return Handler{userSvc: userSvc, authSvc: authSvc}
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

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res, err := h.userSvc.CreateNewUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h Handler) loginUser(c echo.Context) error {
	var req userparam.LoginRequest
	var res userparam.LoginResponse
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res, err := h.userSvc.Login(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
