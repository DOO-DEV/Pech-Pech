package presenter

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Validator struct {
}

func NewAuthValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateLoginRequest(req LoginRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required, validation.Length(8, 8)),
		validation.Field(&req.Username, validation.Required),
	)
}
func (v Validator) ValidateRegister(req RegisterRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Password, validation.Required),
	)
}

func (v Validator) ValidateVerifyOtpRequest(req VerifyResetPasswordOtpRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Code, validation.Required, validation.Length(5, 5)),
	)
}

func (v Validator) ValidateForgetPasswordRequest(req ForgetPasswodRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
	)
}

func (v Validator) ValidateUpdatePasswordRequest(req UpdatePasswordRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 8)),
	)
}
