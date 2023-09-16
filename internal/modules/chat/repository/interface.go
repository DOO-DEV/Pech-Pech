package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *models.Chat) (string, error)
}
