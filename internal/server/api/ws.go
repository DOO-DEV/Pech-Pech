package api

import (
	"github.com/doo-dev/pech-pech/internal/modules/chat/delivery"
	"github.com/doo-dev/pech-pech/internal/modules/chat/hub"
	"github.com/doo-dev/pech-pech/internal/modules/chat/repository"
)

func (a Api) Websocket() error {
	chatRepo := repository.NewChatRepository()
	wsChatHub := hub.NewHub(chatRepo)
	go wsChatHub.Broadcast()
	chatHandler := delivery.NewChatHandler(wsChatHub)

	delivery.SetRoutes(a.Echo, chatHandler)

	return nil
}
