package outbound

import (
	"context"
	"time"
)

type TokenCache interface {
	Set(ctx context.Context, key string, token string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
