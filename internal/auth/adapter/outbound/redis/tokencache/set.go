package tokencache

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func (r *Repository) Set(ctx context.Context, key string, token string, ttl time.Duration) error {
	status := r.redisClient.Set(ctx, r.GenerateCacheKey(key), token, ttl)
	if status.Err() != nil {
		return errors.Wrap(status.Err(), "error on auth cache repository")
	}
	return nil
}
