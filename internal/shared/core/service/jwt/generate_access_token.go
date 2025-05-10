package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (s Service) GenerateAccessToken(claims map[string]interface{}) (string, error) {
	now := time.Now()
	expiresAt := now.Add(s.config.AccessTokenExpiry)

	jwtClaims := jwt.MapClaims{
		"iss": s.config.JWTIssuer,
		"aud": s.config.JWTAudience,
		"iat": now.Unix(),
		"exp": expiresAt.Unix(),
	}

	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := token.SignedString([]byte(s.config.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
