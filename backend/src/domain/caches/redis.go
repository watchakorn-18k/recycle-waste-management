package caches

import (
	"context"
	"recycle-waste-management-backend/src/domain/datasources"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Context   context.Context
	RedisWR   *redis.Client
	RedisRead *redis.Client
}

type IRedisCache interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}

func NewRedisCache(redis *datasources.RedisConnection) IRedisCache {
	return &RedisCache{
		Context:   redis.Context,
		RedisWR:   redis.RedisWR,
		RedisRead: redis.RedisRead,
	}
}

func (cache *RedisCache) Set(key string, value interface{}) error {
	return cache.RedisWR.Set(cache.Context, key, value, 0).Err()
}

func (cache *RedisCache) Get(key string) (string, error) {
	return cache.RedisRead.Get(cache.Context, key).Result()
}
