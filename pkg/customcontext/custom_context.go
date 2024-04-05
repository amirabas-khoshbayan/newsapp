package customcontext

import "github.com/labstack/echo/v4"

type ApiContext struct {
	echo.Context
}

func (c ApiContext) GetHeaderToken() string {
	return c.Request().Header.Get("authorization")
}
