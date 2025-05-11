package inbound

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

type ErrorService interface {
	NewError(ctx context.Context, err error, kind sharederr.ErrorKind, source string) error
}
