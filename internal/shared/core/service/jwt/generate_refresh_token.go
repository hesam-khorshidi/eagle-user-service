package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	"time"
)

func (s Service) GenerateRefreshToken(claims map[string]interface{}) (sharedvo.ID, string, error) {
	now := time.Now()
	expiresAt := now.Add(s.config.RefreshTokenExpiry)

	jwtClaims := jwt.MapClaims{
		"iat": now.Unix(),
		"exp": expiresAt.Unix(),
	}

	for key, value := range claims {
		jwtClaims[key] = value
	}

	tokenID := s.idGen.ID()
	if _, ok := jwtClaims["jti"]; !ok {
		jti := tokenID
		jwtClaims["jti"] = jti
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims) // Using HS256, configure as needed

	tokenString, err := token.SignedString([]byte(s.config.RefreshTokenSecret))
	if err != nil {
		return 0, "", err
	}

	return tokenID, tokenString, nil
}
