package api

import (
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/middlewares"
	"github.com/doo-dev/pech-pech/internal/modules/auth/delivery"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepository "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	authService "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	userRepository "github.com/doo-dev/pech-pech/internal/modules/users/repository"
)

func (a Api) HttpApi() error {
	maiAdt := mail.NewMail(a.mailConf)

	userRepo := userRepository.NewUserRepository(a.pgDB)
	authRepo := authRepository.NewAuthRepository(a.pgDB)

	authSvc := authService.NewAuthService(a.authConf, userRepo, authRepo, maiAdt)

	authValidator := presenter.NewAuthValidator()
	authHandler := delivery.NewAuthHandler(authSvc, authValidator)

	authMw := middlewares.NewAuthMiddleware(authSvc)

	g := a.Echo.Group("/api/v1")
	delivery.SetRoutes(g, authHandler, authMw)

	return nil
}
