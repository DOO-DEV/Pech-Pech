package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/helper"
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

func (a authRepository) UpdatePassword(ctx context.Context, email, password string) error {
	hashedPassword, err := helper.Encrypt(password)
	if err != nil {
		return err
	}
	if err := a.pgDB.WithContext(ctx).Where(`email = ?`, email).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}
