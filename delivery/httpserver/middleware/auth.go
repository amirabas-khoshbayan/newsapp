package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"newsapp/pkg/customcontext"
	"newsapp/service/authenticationservice"
	"newsapp/service/authorizationservice"
)

func Auth(authSvc authenticationservice.Service) echo.MiddlewareFunc {

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(authSvc.Config.SignKey),

		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			customContext := c.(*customcontext.ApiContext)
			headerToken := customContext.GetHeaderToken()
			claims, err := authSvc.ParseToken(headerToken)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}
func CheckPermissions(authSvc authenticationservice.Service, authorizeSvc authorizationservice.Service, permissionRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			customContext := c.(*customcontext.ApiContext)
			headerToken := customContext.GetHeaderToken()
			claims, err := authSvc.ParseToken(headerToken)

			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}

			hasAccess := authorizeSvc.CheckAccess(permissionRole, string(claims.Role))
			if !hasAccess {
				return c.JSON(http.StatusBadRequest, echo.Map{"error": "no access"})
			}

			return next(c)
		}
	}

}
