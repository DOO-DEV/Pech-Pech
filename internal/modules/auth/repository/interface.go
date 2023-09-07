package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}
