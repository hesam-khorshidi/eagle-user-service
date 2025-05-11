package tokencache

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/outbound"
	"github.com/redis/go-redis/v9"
)

const CacheKey = "blacklist:token:jti"

var _ outbound.TokenCache = (*Repository)(nil)

type Repository struct {
	redisClient *redis.Client
}

func New(rc *redis.Client) *Repository {
	return &Repository{redisClient: rc}
}
