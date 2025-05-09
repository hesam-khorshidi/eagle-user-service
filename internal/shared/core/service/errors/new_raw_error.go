package errors

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (s Service) NewRawError(_ context.Context, err error, source string) error {
	if err == nil {
		return nil
	}
	return errors.Wrap(err, fmt.Sprintf("error on %s", source))
}
