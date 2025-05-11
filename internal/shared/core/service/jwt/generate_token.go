package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/pkg/errors"
)

func (s *Service) GenerateToken(userID sharedvo.ID, tokenType sharedvo.TokenType) (string, error) {
	now := time.Now()
	jti := s.idGen.ID()

	var tokenClaims jwt.Claims
	var signingKey string
	var expiryDuration time.Duration

	switch tokenType {
	case sharedvo.TokenTypeAccessToken:
		signingKey = s.config.AccessTokenSecret
		expiryDuration = s.config.AccessTokenExpiry

		tokenClaims = &jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(int64(userID), 10),
			Issuer:    s.config.JWTIssuer,
			Audience:  jwt.ClaimStrings{s.config.JWTAudience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiryDuration)),
			ID:        strconv.FormatInt(int64(jti), 10),
		}

	case sharedvo.TokenTypeRefreshToken:
		signingKey = s.config.RefreshTokenSecret
		expiryDuration = s.config.RefreshTokenExpiry

		tokenClaims = &jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(int64(userID), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiryDuration)),
			ID:        strconv.FormatInt(int64(jti), 10),
		}

	default:
		return "", sharederr.ErrInvalidTokenType
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.Wrap(err, "error on jwt service")
	}

	return tokenString, nil
}
