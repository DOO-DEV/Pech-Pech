package presenter

import (
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation"
)

type RoomValidator struct {
}

func NewRoomValidator() RoomValidator {
	return RoomValidator{}
}

func (v RoomValidator) ValidateCreateRoomRequest(req *CreateRoomRequest) (map[string]string, error) {
	const op = "validator.ValidateCreateRoomRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 30)),
		validation.Field(&req.Description, validation.Length(4, 200)),
		validation.Field(&req.Category, validation.Required, validation.Length(4, 20)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v RoomValidator) convertErrorsToMap(err error, op richerror.Op) (map[string]string, error) {
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
