package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) ExtractExpirationDate(ctx context.Context, token string) (time.Time, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.RefreshTokenSecret), nil
	})
	if err != nil {
		return time.Time{}, sharederr.ErrInvalidToken
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			return time.Unix(int64(exp), 0), nil
		}
		return time.Time{}, sharederr.ErrInvalidToken
	}
	return time.Time{}, sharederr.ErrInvalidToken
}
