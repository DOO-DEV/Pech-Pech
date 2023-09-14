package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type roomRepository struct {
	pgDB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{pgDB: db}
}

func (r roomRepository) CreateRoom(ctx context.Context, room *models.Room) error {
	const op = "roomrepository.CreateRoom"

	if err := r.pgDB.WithContext(ctx).Create(room).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// duplicate key
			if pgErr.Code == "23505" {
				return richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).
					WithMessage(constants.ErrMsgUsernameExisted)
			}
		}

		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (r roomRepository) GetUserRooms(ctx context.Context, userID string) ([]*models.Room, error) {
	const op = "roomrepository.GetUserRooms"
	var rooms []*models.Room

	if err := r.pgDB.WithContext(ctx).Where(`create_by = ?`, userID).Find(&rooms).Error; err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return rooms, nil
}
