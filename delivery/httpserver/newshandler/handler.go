package newshandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/param/newsparam"
	"newsapp/pkg/customcontext"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/newsservice"
	"strconv"
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
	customContext := c.(*customcontext.ApiContext)
	headerToken := customContext.GetHeaderToken()
	claims, err := h.authSvc.ParseToken(headerToken)

	var req newsparam.CreateNewsRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res, err := h.newsSvc.CreateNewNews(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	parseUintID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res.CreatorUserID = uint(parseUintID)

	return c.JSON(http.StatusOK, res)
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
