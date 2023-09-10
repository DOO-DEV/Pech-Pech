package richerror

type Kind int

const (
	KindInvalid = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type Op string

type RichError struct {
	operation    Op
	kind         Kind
	message      string
	meta         map[string]interface{}
	wrappedError error
}

func (r RichError) Error() string {
	if r.message == "" {
		return r.wrappedError.Error()
	}

	return r.message
}

func New(op Op) RichError {
	return RichError{
		operation: op,
	}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op

	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WithError(wrapErr error) RichError {
	r.wrappedError = wrapErr

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) WithMessage(msg string) RichError {
	r.message = msg

	return r
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.kind
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return r.wrappedError.Error()
	}

	return re.Message()
}
