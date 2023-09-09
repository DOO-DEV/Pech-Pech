package middlewares

import (
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
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

		headerParts := strings.Split(authorization, " ")
		// TODO - add prefix to config
		if len(headerParts) != 2 && headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims, err := a.authSvc.ParseToken(headerParts[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		c.Set("user", claims)

		return next(c)
	}
}
