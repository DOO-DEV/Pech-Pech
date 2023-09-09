package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdatePassword(ctx context.Context, email, password string) error
}
