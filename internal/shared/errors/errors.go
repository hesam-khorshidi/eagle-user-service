package errors

import "github.com/pkg/errors"

type Error error

var (
	ErrEntityNotFound = errors.New("entity not found")
)
