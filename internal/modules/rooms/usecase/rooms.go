package usecase

import (
	"context"
	"fmt"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/repository"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"github.com/google/uuid"
)

type RoomsSvc struct {
	roomRepo repository.RoomRepository
}

func NewRoomSvc(roomRepo repository.RoomRepository) RoomsSvc {
	return RoomsSvc{roomRepo: roomRepo}
}

func (r RoomsSvc) CreateRoom(ctx context.Context, req *presenter.CreateRoomRequest, userID string) error {
	const op = "roomservice.CreateRoom"

	fmt.Print(req, userID)
	room := &models.Room{
		ID:          uuid.New().String(),
		Description: req.Description,
		Category:    req.Category,
		CreatedBy:   userID,
	}

	if err := r.roomRepo.CreateRoom(ctx, room); err != nil {
		return richerror.New(op).WithError(err)
	}

	return nil
}

func (r RoomsSvc) GetRooms(ctx context.Context, userID string) ([]*models.Room, error) {
	const op = "roomservice.GetRooms"

	rooms, err := r.roomRepo.GetUserRooms(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return rooms, nil
}

func (r RoomsSvc) DeleteRoom()        {}
func (r RoomsSvc) UpdateDescription() {}
