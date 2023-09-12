package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *models.Room) error
	GetUserRooms(ctx context.Context, userID string) ([]*models.Room, error)
}
