package httperr

import (
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"net/http"
)

func Error(err error) (int, string) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()
		code := mapKindToHttpStatusCode(re.Kind())

		// we don't want to show end user what is the real problem of our system
		if code >= 500 {
			msg = constants.ErrMsgSomethingWentWrong
		}

		return code, msg
	default:
		return http.StatusBadRequest, err.Error()

	}
}

func mapKindToHttpStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	default:
		return http.StatusBadRequest
	}
}
