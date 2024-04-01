package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/entity"
	"newsapp/param/userparam"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/userservice"
)

type Handler struct {
	userSvc      userservice.Service
	authorizeSvc authorizationservice.Service
	authSvc      authenticationservice.Service
}

func New(userSvc userservice.Service, authSvc authenticationservice.Service, authorizeSvc authorizationservice.Service) Handler {
	return Handler{userSvc: userSvc, authSvc: authSvc, authorizeSvc: authorizeSvc}
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

	req.Role = entity.UserRole

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
func (h Handler) giveAdminRole(c echo.Context) error {
	id := c.Param("id")
	err := h.userSvc.GiveAdminRole(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success"})

}
