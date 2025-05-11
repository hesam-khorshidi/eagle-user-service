package user

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/pkg/errors"
)

func (s *Service) Update(ctx context.Context, user domain.User, fields ...valueobject.UserField) error {
	err := s.userRepo.Update(ctx, user, fields...)
	if err != nil {
		if errors.Is(err, sharederr.ErrEntityNotFound) {
			return s.errorSrv.NewError(ctx, err, sharederr.ErrKindNotFound, "user service")
		}
		return s.errorSrv.NewError(ctx, err, sharederr.ErrKindInternal, "user service")
	}
	return nil
}
