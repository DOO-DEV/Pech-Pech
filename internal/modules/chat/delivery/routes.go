package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Echo, h chatHandler, mw *middlewares.AuthMiddleware) {
	g.GET("/ws", h.ChatConnect, mw.JwtValidate)
}
