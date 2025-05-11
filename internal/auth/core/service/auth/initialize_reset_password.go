package auth

import (
	"context"

	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/pkg/errors"
)

func (s *Service) InitializeResetPassword(ctx context.Context, id sharedvo.ID) error {
	//TODO implement me
	return s.errSrv.NewError(ctx, errors.New("unimplemented service"), sharederr.ErrKindInternal, "auth service")
}
