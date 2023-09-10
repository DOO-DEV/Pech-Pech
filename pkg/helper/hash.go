package helper

import (
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	const op = "helper.Encrypt"

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", richerror.New(op).WithError(err).
			WithKind(richerror.KindUnexpected).WithMessage(constants.ErrMsgSomethingWentWrong)
	}

	return string(bytes), nil
}

func Decrypt(password, hashedPassword string) error {
	const op = "helper.Decrypt"

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return richerror.New(op).WithError(err).
			WithKind(richerror.KindInvalid).WithMessage(constants.ErrMsgNoRecord)
	}

	return nil
}
