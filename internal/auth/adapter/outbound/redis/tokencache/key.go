package tokencache

import (
	"fmt"
)

func (r *Repository) GenerateCacheKey(key string) string {
	return fmt.Sprintf("%s:%s", CacheKey, key)
}
