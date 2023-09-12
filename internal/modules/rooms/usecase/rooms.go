package usecase

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/repository"
	"github.com/doo-dev/pech-pech/pkg/richerror"
)

type RoomsSvc struct {
	roomRepo repository.RoomRepository
}

func NewRoomSvc(roomRepo repository.RoomRepository) RoomsSvc {
	return RoomsSvc{roomRepo: roomRepo}
}

func (r RoomsSvc) CreateRoom(ctx context.Context, req *presenter.CreateRoomRequest, userID string) error {
	const op = "roomservice.CreateRoom"

	room := &models.Room{
		ID:          uuid.New().String(),
		Description: req.Description,
		Category:    req.Category,
		CreatedBy:   userID,
	}

	if err := r.roomRepo.CreateRoom(ctx, room); err != nil {
		return richerror.New(op).WithError(err)
	}
}
func (r RoomsSvc) GetRooms()          {}
func (r RoomsSvc) DeleteRoom()        {}
func (r RoomsSvc) UpdateDescription() {}
