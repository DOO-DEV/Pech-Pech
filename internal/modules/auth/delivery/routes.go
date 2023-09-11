package delivery

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetRoutes(g *echo.Group, h AuthHandler, mw *middlewares.AuthMiddleware) {

	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.POST("/forget-password", h.ForgetPassword)
	g.PATCH("/reset-password", h.ResetPassword)
	g.PATCH("/update-password", h.UpdatePassword, mw.JwtValidate)
}
