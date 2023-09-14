package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h *RoomHandler, mw *middlewares.AuthMiddleware) {
	g.POST("/create", h.CreateRoom)
	g.GET("", h.GetUserRooms)
}
