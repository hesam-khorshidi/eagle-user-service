package errors

import (
	"context"
	"fmt"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/pkg/errors"
)

func (s *Service) NewError(ctx context.Context, err error, kind sharederr.ErrorKind, source string) error {
	if err == nil {
		return nil
	}
	reportable := false
	if kind == sharederr.ErrKindInternal {
		//TODO add logic to report error
		reportable = true
	}
	return &sharederr.AppError{
		Kind:       kind,
		Reportable: reportable,
		Err:        errors.Wrap(err, fmt.Sprintf("error on %s", source)),
	}
}
