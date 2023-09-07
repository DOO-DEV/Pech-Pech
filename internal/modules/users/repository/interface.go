package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
)

type UserRepository interface {
	GetUserByIdOrUsername(ctx context.Context, idOrUsername string) (models.User, error)
}
