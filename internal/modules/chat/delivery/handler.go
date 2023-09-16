package delivery

import (
	"fmt"
	"github.com/doo-dev/pech-pech/internal/modules/chat/hub"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

type chatHandler struct {
	hub *hub.Hub
}

func NewChatHandler(h *hub.Hub) chatHandler {
	return chatHandler{hub: h}
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h chatHandler) ChatConnect(c echo.Context) error {
	ws, err := upgrade.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Errorf("error while upgrading to ws: %w", err)
	}

	client := hub.NewClient(ws)
	h.hub.Clients[client] = true

	h.hub.Receiver(c.Request().Context(), client)

	defer delete(h.hub.Clients, client)

	return c.JSON(http.StatusOK, nil)
}
