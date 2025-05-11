package auth

import (
	"context"

	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/pkg/encrypt"
)

func (s *Service) Register(ctx context.Context, user domain.User) (*valueobject.Token, error) {
	emailExists, err := s.userSrv.EmailExists(ctx, user.Email)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "failed to check email existence")
	}
	if emailExists {
		return nil, s.errSrv.NewError(ctx, sharederr.ErrEntityAlreadyExists, sharederr.ErrKindConflict, "email already in use")
	}

	plainPassword := user.Password
	hashedPassword, err := encrypt.HashPassword(plainPassword)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "failed to hash password")
	}
	user.Password = hashedPassword

	user.ID = s.idGen.ID()
	if err := s.userSrv.Create(ctx, user.ToUserDomain()); err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "failed to create user")
	}

	token, err := s.Login(ctx, user.Email, plainPassword)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "login after registration failed")
	}

	return token, nil
}
