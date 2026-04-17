package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(ctx context.Context, key string, dest any) (bool, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		zap.L().Error("RedisCache.Get error on find data:", zap.Error(err))
		return false, fmt.Errorf("RedisCache.Get error on find data: %w", err)
	}

	if err := json.Unmarshal([]byte(value), dest); err != nil {
		zap.L().Error("RedisCache.Get error on unmarshal data:", zap.Error(err))
		return false, fmt.Errorf("RedisCache.Get error on unmarshal data: %w", err)
	}

	return true, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	var payload any = value

	if _, ok := value.(string); !ok {
		body, err := json.Marshal(value)
		if err != nil {
			zap.L().Error("RedisCache.Set error on marshal data:", zap.Error(err))
			return fmt.Errorf("RedisCache.Set error on marshal data: %w", err)
		}

		payload = body
	}

	err := r.client.Set(ctx, key, payload, ttl).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}

		zap.L().Error("RedisCache.Set error on set data:", zap.Error(err))
		return fmt.Errorf("RedisCache.Set error on set data: %w", err)
	}

	return nil
}

func (r *RedisCache) SetNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	b, err := json.Marshal(value)
	if err != nil {
		zap.L().Error("Error on marshaling data", zap.Error(err))
		return false, err
	}

	resp, err := r.client.SetArgs(ctx, key, b, redis.SetArgs{
		Mode: "NX",
		TTL:  ttl,
	}).Result()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		zap.L().Error("RedisCache.SetNX error on set data:", zap.Error(err))
		return false, fmt.Errorf("RedisCache.SetNX error on set data: %w", err)
	}

	if resp == "" {
		return false, nil
	}

	return true, nil
}
