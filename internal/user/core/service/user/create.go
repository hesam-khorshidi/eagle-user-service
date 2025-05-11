package user

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
)

func (s *Service) Create(ctx context.Context, user domain.User) error {
	err := s.userRepo.Create(ctx, user)
	return s.errorSrv.NewError(ctx, err, sharederr.ErrKindInternal, "user service")
}
