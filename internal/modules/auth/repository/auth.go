package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"gorm.io/gorm"
)

type authRepository struct {
	pgDB *gorm.DB
}

func NewAuthRepository(pgDB *gorm.DB) *authRepository {
	return &authRepository{pgDB: pgDB}
}

func (a authRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := a.pgDB.WithContext(ctx).Create(&user).Error; err != nil {
		// TODO - log error
		return err
	}

	return nil
}
