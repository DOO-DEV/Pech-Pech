package repository

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type authRepository struct {
	pgDB *gorm.DB
}

func NewAuthRepository(pgDB *gorm.DB) *authRepository {
	return &authRepository{pgDB: pgDB}
}

func (a authRepository) CreateUser(ctx context.Context, user *models.User) error {
	const op = "psql.CreateUser"

	if err := a.pgDB.WithContext(ctx).Create(&user).Error; err != nil {
		// TODO - log error
		// TODO - find a better solution for this error
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// duplicate key
			if pgErr.Code == "23505" {
				return richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).
					WithMessage(constants.ErrMsgUsernameExisted)
			}
		}

		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (a authRepository) UpdatePassword(ctx context.Context, email, password string) error {
	const op = "psql.UpdatePassword"

	if err := a.pgDB.WithContext(ctx).Model(&models.User{}).Where(`email = ?`, email).Update("password", password).Error; err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
