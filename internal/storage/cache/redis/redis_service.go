package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const CacheDuration = 6 * time.Minute

type RedisService struct {
	redisClient *redis.Client
}

var (
	ctx = context.Background()
)

func NewRedisService(address string, timeout, dialTimeout time.Duration) (*RedisService, error) {
	const op = "cache.redis.NewRedisService"
	redisClient := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     "",
		DB:           0,
		DialTimeout:  dialTimeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &RedisService{
		redisClient: redisClient,
	}, nil
}

// Save url and alias in cache
func (s *RedisService) SaveUrlMapping(url, alias string) error {
	const op = "cache.redis.SaveUrlMapping"

	if err := s.redisClient.Set(ctx, alias, url, CacheDuration).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Retrieve url by alias
func (s *RedisService) RetrieveUrl(alias string) (string, error) {
	const op = "storage.redis.RetrieveUrl"

	result, err := s.redisClient.Get(ctx, alias).Result()
	if err != nil {
		if err == redis.Nil {
			return "", err
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

// Delete url by alias
func (s *RedisService) DeleteUrl(alias string) error {
	const op = "storage.redis.DeleteUrl"

	if err := s.redisClient.Del(ctx, alias).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
