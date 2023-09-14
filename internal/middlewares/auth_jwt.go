package middlewares

import (
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/pkg/httperr"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authSvc usecase.AuthService
}

func NewAuthMiddleware(authSvc usecase.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authSvc: authSvc}
}

func (a AuthMiddleware) JwtValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		// TODO - add prefix to config
		if len(token) < 1 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims, err := a.authSvc.ParseToken(token)
		if err != nil {
			code, msg := httperr.Error(err)
			return echo.NewHTTPError(code, msg)
		}

		c.Set("user", claims)

		return next(c)
	}
}
