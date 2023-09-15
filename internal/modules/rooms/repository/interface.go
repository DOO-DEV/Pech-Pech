package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *models.Room) error
	GetUserRooms(ctx context.Context, userID string) ([]*models.Room, error)
	DeleteRoom(ctx context.Context, roomName, userID string) error
	UpdateRoom(ctx context.Context, room *models.Room, oldRoomName, userID string) (*models.Room, error)
}
