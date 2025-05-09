package user

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/pkg/errors"
)

func (s Service) FindBy(ctx context.Context, field valueobject.UserField, value any) (*domain.User, error) {
	user, err := s.userRepo.FindBy(ctx, field, value)
	if err != nil {
		if errors.Is(err, sharederr.ErrEntityNotFound) {
			return nil, s.errorSrv.NewRawError(ctx, err, "user service")
		}
		return nil, s.errorSrv.NewReportableError(ctx, err, "user service")
	}
	return user, nil
}
