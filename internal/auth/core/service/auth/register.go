package auth

import (
	"context"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
	"github.com/hesam-khorshidi/eagle-user-service/pkg/encrypt"
)

func (s Service) Register(ctx context.Context, user domain.User) (*valueobject.Token, error) {
	emailExists, err := s.userSrv.EmailExists(ctx, user.Email)
	if err != nil {
		return nil, s.errSrv.NewReportableError(ctx, err, "auth service")
	}
	if emailExists {
		return nil, s.errSrv.NewRawError(ctx, err, "auth service")
	}

	hashedPassword, err := encrypt.HashPassword(user.Password)
	if err != nil {
		return nil, s.errSrv.NewReportableError(ctx, err, "auth service")
	}
	user.Password = hashedPassword

	err = s.userSrv.Create(ctx, user.ToUserDomain())
	if err != nil {
		return nil, s.errSrv.NewReportableError(ctx, err, "auth service")
	}
	token, err := s.Login(ctx, user.Email, user.Password)
	if err != nil {
		return nil, s.errSrv.NewReportableError(ctx, err, "auth service")
	}
	return token, nil
}
