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
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		CreatedBy:   userID,
	}

	if err := r.roomRepo.CreateRoom(ctx, room); err != nil {
		return richerror.New(op).WithError(err)
	}

	return nil
}

func (r RoomsSvc) GetRooms(ctx context.Context, userID string) ([]*presenter.GetRoomsResponse, error) {
	const op = "roomservice.GetRooms"

	rooms, err := r.roomRepo.GetUserRooms(ctx, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	var res []*presenter.GetRoomsResponse
	for _, v := range rooms {
		r := &presenter.GetRoomsResponse{
			Name:        v.Name,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   v.CreatedAt,
		}
		res = append(res, r)
	}

	return res, nil
}

func (r RoomsSvc) DeleteRoom(ctx context.Context, dto *presenter.DeleteRoomRequest, userID string) error {
	const op = "roomservice.DeleteRoom"

	if err := r.roomRepo.DeleteRoom(ctx, dto.Name, userID); err != nil {
		return richerror.New(op).WithError(err)
	}

	return nil
}
func (r RoomsSvc) UpdateRoomInfo(ctx context.Context, dto *presenter.UpdateRoomInfoRequest, userID string) (*presenter.UpdateRoomInfoResponse, error) {
	const op = "roomservice.UpdateRoomInfo"

	rm := &models.Room{
		Name:        dto.NewName,
		Description: dto.Description,
	}

	room, err := r.roomRepo.UpdateRoom(ctx, rm, dto.OldName, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	res := &presenter.UpdateRoomInfoResponse{
		Name:        room.Name,
		Description: room.Description,
	}

	return res, nil
}
