package repository

import (
	"context"
	"fmt"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/abstract"
	"github.com/doo-dev/pech-pech/pkg/constants"
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

func (a userRepository) Search(ctx context.Context, name string, pagination *abstract.Pagination) ([]*models.User, error) {
	const op = "psql.Search"

	fmt.Println(pagination.GetSize(), pagination.GetOffset(), name)
	var users []*models.User
	if err := a.pgDB.WithContext(ctx).Where("username LIKE ?", "%"+name+"%").
		Limit(pagination.GetPage()).Offset(pagination.GetSize()).Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, richerror.New(op).WithError(err).
				WithKind(richerror.KindNotFound).WithMessage(constants.ErrMsgNoRecord)
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return users, nil
}
