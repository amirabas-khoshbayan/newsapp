package newshandler

import (
	"github.com/labstack/echo/v4"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/newsservice"
)

type Handler struct {
	newsSvc      newsservice.Service
	authorizeSvc authorizationservice.Service
	authSvc      authenticationservice.Service
}

func New(newsSvc newsservice.Service, authorizeSvc authorizationservice.Service, authenSvc authenticationservice.Service) Handler {
	return Handler{
		newsSvc:      newsSvc,
		authorizeSvc: authorizeSvc,
		authSvc:      authenSvc,
	}
}

func (h Handler) createNews(c echo.Context) error {

	return nil
}
func (h Handler) getNews(c echo.Context) error {

	return nil
}
func (h Handler) getNewsList(c echo.Context) error {

	return nil
}
func (h Handler) addToWaitingList(c echo.Context) error {

	return nil
}
