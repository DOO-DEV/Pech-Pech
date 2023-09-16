package delivery

import (
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Echo, h chatHandler) {
	g.GET("/ws", h.ChatConnect)
}
