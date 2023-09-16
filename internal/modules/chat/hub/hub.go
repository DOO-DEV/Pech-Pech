package hub

import (
	"context"
	"encoding/json"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/chat/repository"
	"log"
	"time"
)

type Hub struct {
	Clients    map[*Client]bool
	ChatBucket chan *models.Chat
	chatRepo   repository.ChatRepository
}

func NewHub(repo repository.ChatRepository) *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		ChatBucket: make(chan *models.Chat),
		chatRepo:   repo,
	}
}

func (h *Hub) Broadcast() {
	for {
		msg := <-h.ChatBucket

		for cl := range h.Clients {
			if cl.username == msg.From || cl.username == msg.To {
				err := cl.socket.WriteJSON(msg)
				if err != nil {
					log.Printf("websocket error: %s", err)
					cl.socket.Close()
					delete(h.Clients, cl)
				}
			}
		}
	}
}

func (h *Hub) Receiver(ctx context.Context, client *Client) {
	for {
		_, p, err := client.socket.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := &models.Message{}
		err = json.Unmarshal(p, m)
		if err != nil {
			log.Println("error while unmarshalling: ", err)
			continue
		}
		if m.Type == "bootup" {
			client.username = m.User
		} else {
			c := m.Chat
			c.Timestamp = time.Now().Unix()

			id, err := h.chatRepo.CreateChat(ctx, &c)
			if err != nil {
				log.Println("error while saving chat in redis", err)
				return
			}
			c.ID = id
			h.ChatBucket <- &c
		}

	}
}
