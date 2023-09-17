package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h chatHandler, mw *middlewares.AuthMiddleware) {
	g.GET("", h.ChatConnect, mw.JwtValidate)
}
