package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type repository struct {
}

func NewChatRepository() ChatRepository {
	return &repository{}
}

func (r repository) CreateChat(ctx context.Context, chat *models.Chat) (string, error) {
	return "", nil
}
