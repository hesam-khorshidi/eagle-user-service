package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, sharederr.ErrInvalidTokenSignature
		}
		return []byte(s.config.RefreshTokenSecret), nil
	})

	if err != nil {
		return sharederr.ErrInvalidTokenSignature
	}

	now := time.Now()
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		if claims.NotBefore != nil && claims.NotBefore.Time.After(now) {
			return sharederr.ErrTokenNotValidYet
		}
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(now) {
			return sharederr.ErrTokenExpired
		}
		return nil
	}

	return sharederr.ErrInvalidToken
}
