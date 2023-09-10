package presenter

import (
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Validator struct {
}

func NewAuthValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateLoginRequest(req LoginRequest) (map[string]string, error) {
	const op = "validator.ValidateLoginRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required, validation.Length(8, 16)),
		validation.Field(&req.Username, validation.Required),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}
func (v Validator) ValidateRegister(req RegisterRequest) (map[string]string, error) {
	const op = "validator.ValidateRegister"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Password, validation.Required),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v Validator) ValidateResetPassword(req ResetPasswordRequest) (map[string]string, error) {
	const op = "validator.ValidateResetPassword"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Code, validation.Required, validation.Length(5, 5)),
		validation.Field(&req.ConfirmPassword, validation.Required, validation.Length(8, 16)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 16)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v Validator) ValidateForgetPasswordRequest(req ForgetPasswodRequest) (map[string]string, error) {
	const op = "validator.ValidateForgetPasswordRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v Validator) ValidateUpdatePasswordRequest(req UpdatePasswordRequest) (map[string]string, error) {
	const op = "validate.ValidateUpdatePasswordRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required, validation.Length(8, 16)),
		validation.Field(&req.ConfirmPassword, validation.Required, validation.Length(8, 16)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v Validator) convertErrorsToMap(err error, op richerror.Op) (map[string]string, error) {
	fields := make(map[string]string)

	errV, ok := err.(validation.Errors)
	if ok {
		for key, value := range errV {
			if value != nil {
				fields[key] = value.Error()
			}
		}
	}
	return fields, richerror.New(op).
		WithError(err).WithMessage(constants.ErrMsgInvalidInput).
		WithKind(richerror.KindInvalid)
}
