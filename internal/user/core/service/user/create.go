package user

import (
	"context"

	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
)

func (s Service) Create(ctx context.Context, user domain.User) error {
	err := s.userRepo.Create(ctx, user)
	return s.errorSrv.NewReportableError(ctx, err, "user service")
}
