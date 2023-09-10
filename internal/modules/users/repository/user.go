package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"gorm.io/gorm"
	"strconv"
)

type userRepository struct {
	pgDB *gorm.DB
}

func NewUserRepository(pgDB *gorm.DB) *userRepository {
	return &userRepository{pgDB: pgDB}
}

func (a userRepository) GetUserByIdOrUsername(ctx context.Context, idOrUsername string) (*models.User, error) {
	const op = "psql.GetUserByIdOrUsername"

	user := &models.User{}

	_, err := strconv.Atoi(idOrUsername)
	if err != nil {
		user.Username = idOrUsername
	} else {
		user.ID = idOrUsername
	}

	if err := a.pgDB.WithContext(ctx).Where(user).First(user).Error; err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
	}

	return user, nil
}
