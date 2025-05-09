package errors

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (s Service) NewReportableError(ctx context.Context, err error, source string) error {
	if err == nil {
		return nil
	}
	// Will add report to log service here later
	return errors.Wrap(err, fmt.Sprintf("error on %s", source))
}
