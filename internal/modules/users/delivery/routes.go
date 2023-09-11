package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(userGroup *echo.Group, h UserHandler, authMw *middlewares.AuthMiddleware) {
	userGroup.GET("/search", h.Search, authMw.JwtValidate)
}
