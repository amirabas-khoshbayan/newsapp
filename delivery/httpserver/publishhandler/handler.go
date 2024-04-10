package publishhandler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/publishservice"
)

type Handler struct {
	publishSvc   publishservice.Service
	authorizeSvc authorizationservice.Service
	authSvc      authenticationservice.Service
}

func New(publishSvc publishservice.Service, authorizeSvc authorizationservice.Service, authenSvc authenticationservice.Service) Handler {
	return Handler{
		publishSvc:   publishSvc,
		authorizeSvc: authorizeSvc,
		authSvc:      authenSvc,
	}
}

func (h Handler) publishNews(c echo.Context) error {

	return nil
}
func (h Handler) addToWaitingNewsList(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("invalid id")
	}

	return nil
}
