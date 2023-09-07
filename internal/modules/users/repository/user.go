package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/constants"
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
	var user *models.User

	_, err := strconv.ParseUint(idOrUsername, 10, 64)
	if err != nil {
		user.Username = idOrUsername
	} else {
		user.ID = idOrUsername
	}

	if err := a.pgDB.WithContext(ctx).Where(user).First(user).Error; err == nil {
		return nil, constants.ErrNoRecord
	}

	return user, nil
}
