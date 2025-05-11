package auth

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/pkg/errors"
)

func (s *Service) ResetPassword(ctx context.Context, resetToken string) error {
	//TODO implement me
	return s.errSrv.NewError(ctx, errors.New("unimplemented service"), sharederr.ErrKindInternal, "auth service")
}
