package presenter

import (
	"errors"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RoomValidator struct {
}

func NewRoomValidator() RoomValidator {
	return RoomValidator{}
}

func (v RoomValidator) ValidateCreateRoomRequest(req CreateRoomRequest) (map[string]string, error) {
	const op = "validator.ValidateCreateRoomRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, is.ASCII, validation.Length(3, 20)),
		validation.Field(&req.Description, is.ASCII, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Category, validation.Required, is.ASCII, validation.Length(4, 20)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v RoomValidator) ValidateDeleteRequest(req DeleteRoomRequest) (map[string]string, error) {
	const op = "validator.DeleteRoom"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, is.ASCII, validation.Length(3, 20)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v RoomValidator) ValidateUpdateRoomRequest(req UpdateRoomInfoRequest) (map[string]string, error) {
	const op = "validator.UpdateRoom"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.OldName, validation.Required, is.ASCII, validation.Length(3, 20)),
		validation.Field(&req.NewName, validation.Required, is.ASCII, validation.Length(3, 20)),
		validation.Field(&req.Description, is.ASCII, validation.Length(3, 50)),
	)
	if err != nil {
		return v.convertErrorsToMap(err, op)
	}

	return nil, nil
}

func (v RoomValidator) convertErrorsToMap(err error, op richerror.Op) (map[string]string, error) {
	fields := make(map[string]string)

	var errV validation.Errors
	ok := errors.As(err, &errV)
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
