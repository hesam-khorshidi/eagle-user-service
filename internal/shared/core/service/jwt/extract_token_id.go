package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) ExtractTokenID(ctx context.Context, token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.RefreshTokenSecret), nil
	})
	if err != nil {
		return "", sharederr.ErrInvalidToken
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if jti, ok := claims["jti"].(string); ok {
			return jti, nil
		}
		return "", sharederr.ErrInvalidToken
	}
	return "", sharederr.ErrInvalidToken
}
