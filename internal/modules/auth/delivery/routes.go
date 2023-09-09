package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h AuthHandler, mw *middlewares.AuthMiddleware) {
	authGroup := g.Group("/auth")

	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
	authGroup.POST("/forget-password", h.ForgetPassword)
	authGroup.PATCH("/reset-password", h.ResetPassword)
	authGroup.PATCH("/update-password", h.UpdatePassword, mw.JwtValidate)
}
