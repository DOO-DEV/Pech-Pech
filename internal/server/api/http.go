package api

import (
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/doo-dev/pech-pech/internal/modules/auth/delivery"
	authRepository "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	authService "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	userRepository "github.com/doo-dev/pech-pech/internal/modules/users/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a Api) HttpApi() error {
	userRepo := userRepository.NewUserRepository(a.pgDB)
	authRepo := authRepository.NewAuthRepository(a.pgDB)

	authSvc := authService.NewAuthService(userRepo, authRepo)

	authHandler := delivery.NewAuthHandler(authSvc)

	authMw := middlewares.NewAuthMiddleware(authSvc)

	g := a.Echo.Group("/api/v1")
	g.GET("/helth-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"messagge": "server is working",
		})
	})
	delivery.SetRoutes(g, authHandler, authMw)

	return nil
}
