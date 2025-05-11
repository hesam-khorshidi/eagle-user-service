package errors

import "fmt"

type ErrorKind string

const (
	ErrKindInvalidInput       ErrorKind = "INVALID_INPUT"
	ErrKindUnauthorized       ErrorKind = "UNAUTHORIZED"
	ErrKindNotFound           ErrorKind = "NOT_FOUND"
	ErrKindConflict           ErrorKind = "CONFLICT"
	ErrKindBadRequest         ErrorKind = "BAD_REQUEST"
	ErrKindInternal           ErrorKind = "INTERNAL_SERVER_ERROR"
	ErrKindInvalidCredentials ErrorKind = "INVALID_CREDENTIALS"
)

type AppError struct {
	Kind       ErrorKind
	Message    string
	Reportable bool
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) IsReportable() bool {
	return e.Reportable
}

func (e *AppError) Code() string {
	return string(e.Kind)
}
