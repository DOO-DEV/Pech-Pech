package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h *RoomHandler, _ *middlewares.AuthMiddleware) {
	g.POST("", h.CreateRoom)
	g.GET("", h.GetUserRooms)
	g.DELETE("", h.DeleteRoom)
	g.PATCH("", h.UpdateRoom)
}
