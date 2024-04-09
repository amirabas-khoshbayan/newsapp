package userhandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io"
	"math/rand"
	"net/http"
	"newsapp/entity"
	"newsapp/logger"
	"newsapp/param/userparam"
	"newsapp/pkg/customcontext"
	"newsapp/pkg/httpresponse"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
	"newsapp/service/userservice"
	"os"
	"path/filepath"
	"strconv"
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

	userList, err := h.userSvc.GetUsers(c.Request().Context())

	if err != nil {
		logger.ZapLogger.Named("userHandler").Error("getUserList", zap.Any("userSvc.GetUsers error", err.Error()))
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

	res, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h Handler) giveAdminRole(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	err := h.userSvc.GiveAdminRole(c.Request().Context(), int(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success"})

}
func (h Handler) deleteUser(c echo.Context) error {
	id := c.Param("id")
	err := h.userSvc.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success"})
}
func (h Handler) uploadAvatar(c echo.Context) error {
	customContext := c.(*customcontext.ApiContext)

	file, err := customContext.FormFile("file")
	if err != nil {
		logger.ZapLogger.Named("userHandler").Error("uploadAvatar", zap.Any("customContext.FormFile(\"file\") error", err.Error()))
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	src, err := file.Open()
	if err != nil {
		logger.ZapLogger.Named("userHandler").Error("uploadAvatar", zap.Any("file.Open() error", err.Error()))
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	wd, err := os.Getwd()
	// TODO - add path to config
	if _, err := os.Stat(wd + "/wwwroot/images/useravatar"); os.IsNotExist(err) {
		_ = os.MkdirAll(wd+"/wwwroot/images/useravatar", os.ModePerm)
	}
	randomInt := rand.Int()
	fileID := fmt.Sprintf("%s", randomInt)
	imageServerPath := filepath.Join(wd, "wwwroot", "images", "useravatar", file.Filename+fileID)

	destination, err := os.Create(imageServerPath)
	if err != nil {
		logger.ZapLogger.Named("userHandler").Error("uploadAvatar", zap.Any("os.Create(file.Filename) error", err.Error()))
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	defer destination.Close()

	_, ioErr := io.Copy(destination, src)
	if ioErr != nil {
		logger.ZapLogger.Named("userHandler").Error("uploadAvatar", zap.Any("io.Copy(destination, src) error", ioErr.Error()))
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: ioErr.Error(),
		}))
	}

	return c.JSON(http.StatusOK, httpresponse.New(httpresponse.HttpResponse{
		Code:     http.StatusOK,
		Message:  "success",
		MetaData: map[string]interface{}{"avatar_file_name": file.Filename + fileID},
	}))
}
func (h Handler) editUser(c echo.Context) error {
	id := c.Param("id")

	var req userparam.EditUserRequest
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	req.ID = int(parseInt)

	var res userparam.EditUserResponse
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	res, err := h.userSvc.EditUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.New(httpresponse.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	return c.JSON(http.StatusOK, httpresponse.New(httpresponse.HttpResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	}))
}
