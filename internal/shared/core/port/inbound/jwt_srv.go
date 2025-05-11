package inbound

import (
	"context"
	"time"

	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
)

type JWTService interface {
	GenerateToken(userID sharedvo.ID, tokenType sharedvo.TokenType) (string, error)
	ValidateAccessToken(ctx context.Context, tokenString string) error
	ValidateRefreshToken(ctx context.Context, tokenString string) error
	ExtractTokenID(ctx context.Context, token string) (string, error)
	ExtractUserID(ctx context.Context, tokenType sharedvo.TokenType, token string) (sharedvo.ID, error)
	ExtractExpirationDate(ctx context.Context, token string) (time.Time, error)
}
