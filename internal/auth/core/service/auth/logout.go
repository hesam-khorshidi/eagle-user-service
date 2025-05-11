package auth

import (
	"context"
	"time"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) Logout(ctx context.Context, token string) error {
	tokenID, err := s.jwtService.ExtractTokenID(ctx, token)
	if err != nil {
		return s.errSrv.NewError(ctx, err, sharederr.ErrKindUnauthorized, "auth service")
	}

	expirationDate, err := s.jwtService.ExtractExpirationDate(ctx, token)
	if err != nil {
		return s.errSrv.NewError(ctx, err, sharederr.ErrKindUnauthorized, "auth service")
	}

	if time.Now().After(expirationDate) {
		return s.errSrv.NewError(ctx, err, sharederr.ErrKindUnauthorized, "auth service")
	}

	err = s.tokenCache.Set(ctx, tokenID, token, time.Until(expirationDate))
	if err != nil {
		return s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "error on auth service")
	}
	return nil
}
