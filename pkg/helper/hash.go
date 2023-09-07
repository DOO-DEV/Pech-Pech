package helper

import (
	"github.com/doo-dev/pech-pech/pkg/constants"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", constants.ErrHashPassword
	}

	return string(bytes), nil
}

func Decrypt(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return constants.ErrComparePassword
	}

	return nil
}
