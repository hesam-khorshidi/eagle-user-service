package http

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	jwtsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/jwt"
	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
)

func AuthorizationMiddleware(config jwtsrv.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(config.AccessTokenSecret)},
			ContextKey: "user",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return Unauthorized(c)
			},
			TokenProcessorFunc: func(token string) (string, error) {
				parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, sharederr.ErrInvalidTokenSignature
					}
					return []byte(config.AccessTokenSecret), nil
				})

				if err != nil {
					return "", sharederr.ErrInvalidTokenSignature
				}

				now := time.Now()
				if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok && parsedToken.Valid {
					if claims.NotBefore != nil && claims.NotBefore.Time.After(now) {
						return "", sharederr.ErrTokenNotValidYet
					}
					if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(now) {
						return "", sharederr.ErrTokenExpired
					}
					return token, nil
				}

				return "", sharederr.ErrInvalidToken
			},
		})(c)
	}
}
