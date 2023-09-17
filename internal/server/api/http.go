package api

import (
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/middlewares"
	authDelivery "github.com/doo-dev/pech-pech/internal/modules/auth/delivery"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepository "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	authService "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	chatDelivery "github.com/doo-dev/pech-pech/internal/modules/chat/delivery"
	"github.com/doo-dev/pech-pech/internal/modules/chat/hub"
	chatRepository "github.com/doo-dev/pech-pech/internal/modules/chat/repository"
	roomDelivery "github.com/doo-dev/pech-pech/internal/modules/rooms/delivery"
	roomPresenter "github.com/doo-dev/pech-pech/internal/modules/rooms/presenter"
	roomRepository "github.com/doo-dev/pech-pech/internal/modules/rooms/repository"
	roomService "github.com/doo-dev/pech-pech/internal/modules/rooms/usecase"
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

	roomRepo := roomRepository.NewRoomRepository(a.pgDB)
	roomSvc := roomService.NewRoomSvc(roomRepo)
	roomValidator := roomPresenter.NewRoomValidator()
	roomHandler := roomDelivery.NewRoomHandler(roomSvc, roomValidator)

	// websocket initialization
	chatRepo := chatRepository.NewChatRepository()
	wsChatHub := hub.NewHub(chatRepo)
	go wsChatHub.Broadcast()
	chatHandler := chatDelivery.NewChatHandler(wsChatHub)
	chatDelivery.SetRoutes(a.Echo, chatHandler, authMw)

	p := a.Echo.Group("/api/v1")

	authGroup := p.Group("/auth")
	authDelivery.SetRoutes(authGroup, authHandler, authMw)

	userGroup := p.Group("/users")
	userDelivery.SetRoutes(userGroup, userHandler, authMw)

	roomGroup := p.Group("/rooms", authMw.JwtValidate)
	roomDelivery.SetRoutes(roomGroup, &roomHandler, authMw)

	return nil
}
