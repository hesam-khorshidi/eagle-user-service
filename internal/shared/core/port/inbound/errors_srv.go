package inbound

import (
	"context"
)

type ErrorService interface {
	NewReportableError(ctx context.Context, err error, source string) error
	NewRawError(ctx context.Context, err error, source string) error
}
