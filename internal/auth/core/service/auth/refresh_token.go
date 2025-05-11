package auth

import (
	"context"

	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	tokenID, err := s.jwtService.ExtractTokenID(ctx, refreshToken)
	if err != nil {
		return "", s.errSrv.NewError(ctx, err, sharederr.ErrKindUnauthorized, "auth service")
	}

	userID, err := s.jwtService.ExtractUserID(ctx, valueobject.TokenTypeRefreshToken, refreshToken)
	if err != nil {
		return "", s.errSrv.NewError(ctx, err, sharederr.ErrKindUnauthorized, "auth service")
	}

	blacklistedToken, err := s.tokenCache.Get(ctx, tokenID)
	if err == nil || blacklistedToken != "" {
		return "", s.errSrv.NewError(ctx, sharederr.ErrTokenBlacklisted, sharederr.ErrKindInternal, "auth service")
	}

	accessToken, err := s.jwtService.GenerateToken(userID, valueobject.TokenTypeAccessToken)
	if err != nil {
		return "", s.errSrv.NewError(ctx, err, sharederr.ErrKindInternal, "auth service")
	}
	return accessToken, nil

}
