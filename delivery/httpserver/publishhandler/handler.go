package publishhandler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/param/newsparam"
	"newsapp/pkg/httpresponse"
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
	param := newsparam.AddToWaitingListRequest{}
	if id == "" {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: errors.New("invalid id").Error(),
		}))
	}

	resWaitingList, err := h.publishSvc.AddNewsToWaitingList(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	return c.JSON(http.StatusOK, httpresponse.New(httpresponse.HttpResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resWaitingList,
	}))
}
