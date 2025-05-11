package auth

import (
	"context"
	"database/sql"

	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	uservo "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/hesam-khorshidi/eagle-user-service/pkg/encrypt"
	"github.com/pkg/errors"
)

func (s *Service) Login(ctx context.Context, email, password string) (*valueobject.Token, error) {
	user, err := s.userSrv.FindBy(ctx, uservo.UserFieldEmail, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindNotFound, "auth service")
		}
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "auth service")
	}

	err = encrypt.VerifyPassword(password, user.Password)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInvalidCredentials, "auth service")
	}
	accessToken, err := s.jwtService.GenerateToken(user.ID, sharedvo.TokenTypeAccessToken)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "auth service")
	}
	refreshToken, err := s.jwtService.GenerateToken(user.ID, sharedvo.TokenTypeRefreshToken)
	if err != nil {
		return nil, s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "error on auth service")
	}

	return &valueobject.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
