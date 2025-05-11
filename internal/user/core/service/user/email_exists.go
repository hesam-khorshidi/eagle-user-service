package user

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/pkg/errors"
)

func (s *Service) EmailExists(ctx context.Context, email string) (bool, error) {
	_, err := s.userRepo.FindBy(ctx, valueobject.UserFieldEmail, email)
	if err != nil {
		if errors.Is(err, sharederr.ErrEntityNotFound) {
			return false, nil
		}
		return false, s.errorSrv.NewError(ctx, err, sharederr.ErrKindInternal, "user service")
	}
	return true, nil
}
