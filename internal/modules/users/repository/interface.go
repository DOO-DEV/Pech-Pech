package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/abstract"
)

type UserRepository interface {
	GetUserByIdOrUsername(ctx context.Context, idOrUsername string) (*models.User, error)
	Search(ctx context.Context, name string, pagination *abstract.Pagination) ([]*models.User, error)
}
