package jwt

import (
	"context"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func (s *Service) ExtractUserID(ctx context.Context, tokenType sharedvo.TokenType, token string) (sharedvo.ID, error) {
	var signingKey string
	claims := &jwt.RegisteredClaims{}

	switch tokenType {
	case sharedvo.TokenTypeAccessToken:
		signingKey = s.config.AccessTokenSecret
	case sharedvo.TokenTypeRefreshToken:
		signingKey = s.config.RefreshTokenSecret
	default:
		return 0, sharederr.ErrInvalidTokenType
	}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, sharederr.ErrInvalidTokenSignature
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, sharederr.ErrInvalidToken
	}
	if !parsedToken.Valid {
		return 0, sharederr.ErrInvalidToken
	}

	userID, err := parsedToken.Claims.GetSubject()
	if err != nil {
		return 0, sharederr.ErrInvalidToken
	}

	parsedUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return 0, sharederr.ErrInvalidToken
	}
	return sharedvo.ID(parsedUserID), nil
}
