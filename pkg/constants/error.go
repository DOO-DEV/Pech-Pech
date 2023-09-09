package constants

import "errors"

var (
	ErrUserExisted     = errors.New("user existed")
	ErrHashPassword    = errors.New("error when encrypt password")
	ErrComparePassword = errors.New("password not match")
	ErrNoRecord        = errors.New("record not found")
	ErrSetupHttpRouter = errors.New("error while setting up http router")
	ErrStartHttp       = errors.New("start http server failed")
	ErrOtpExpired      = errors.New("otp is expired")
	ErrOtpInvalid      = errors.New("otp is invalid")
)
