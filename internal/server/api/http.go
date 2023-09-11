package api

import (
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/middlewares"
	authDelivery "github.com/doo-dev/pech-pech/internal/modules/auth/delivery"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepository "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	authService "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	userDelivery "github.com/doo-dev/pech-pech/internal/modules/users/delivery"
	userRepository "github.com/doo-dev/pech-pech/internal/modules/users/repository"
	userService "github.com/doo-dev/pech-pech/internal/modules/users/usecase"
)

func (a Api) HttpApi() error {
	maiAdt := mail.NewMail(a.mailConf)

	userRepo := userRepository.NewUserRepository(a.pgDB)
	userSvc := userService.NewUserSvc(userRepo)
	userHandler := userDelivery.NewUserHandler(userSvc)

	authRepo := authRepository.NewAuthRepository(a.pgDB)
	authSvc := authService.NewAuthService(a.authConf, userRepo, authRepo, maiAdt)
	authValidator := presenter.NewAuthValidator()
	authHandler := authDelivery.NewAuthHandler(authSvc, authValidator)

	authMw := middlewares.NewAuthMiddleware(authSvc)

	p := a.Echo.Group("/api/v1")

	authGroup := p.Group("/auth")
	authDelivery.SetRoutes(authGroup, authHandler, authMw)

	userGroup := p.Group("/users")
	userDelivery.SetRoutes(userGroup, userHandler, authMw)

	return nil
}
