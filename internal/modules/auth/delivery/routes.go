package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h AuthHandler, mw *middlewares.AuthMiddleware) {
	g.Group("/auth")
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
}
