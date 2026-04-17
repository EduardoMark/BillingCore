package rediscache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisCache() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return client
}
