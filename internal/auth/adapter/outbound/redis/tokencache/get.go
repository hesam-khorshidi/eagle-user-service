package tokencache

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	result, err := r.redisClient.Get(ctx, r.GenerateCacheKey(key)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", sharederr.ErrCacheNotFound
		}
		return "", errors.Wrap(err, "error on auth cache repository")
	}

	return result, nil
}
