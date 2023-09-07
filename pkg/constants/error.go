package constants

import "errors"

var (
	ErrUserExisted     = errors.New("user existed")
	ErrHashPassword    = errors.New("error when encrypt password")
	ErrComparePassword = errors.New("password not match")
)
