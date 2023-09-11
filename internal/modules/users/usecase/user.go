package usecase

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/modules/users/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/users/repository"
	"github.com/doo-dev/pech-pech/pkg/abstract"
	"github.com/doo-dev/pech-pech/pkg/richerror"
)

type UserSvc struct {
	userRepo repository.UserRepository
}

func NewUserSvc(repo repository.UserRepository) UserSvc {
	return UserSvc{userRepo: repo}
}

func (u UserSvc) SearchUser(ctx context.Context, name string, pagination *abstract.Pagination) (*presenter.SearchResponse, error) {
	const op = "userservice.SearchUser"

	users, err := u.userRepo.Search(ctx, name, pagination)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	usersInfo := presenter.SearchResponse{}
	for _, u := range users {
		usersInfo.Users = append(usersInfo.Users, u.Username)
	}

	return &usersInfo, nil
}
