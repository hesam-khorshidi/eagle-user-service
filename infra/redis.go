package infra

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func NewRedisClient(cfg RedisConfig) (*redis.Client, func(), error) {
	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})

	return c, func() {
		if err := c.Close(); err != nil {
			fmt.Println(err)
		}
	}, nil
}
