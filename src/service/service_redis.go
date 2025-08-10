package service

import (
	"context"
	"insider/src/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService(cfg *config.ConfigType) *RedisService {
	opt, err := redis.ParseURL(cfg.REDIS_URL)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return &RedisService{
		client: client,
		ctx:    context.Background(),
	}
}

func (rs *RedisService) Get(key string) (string, error) {
	return rs.client.Get(rs.ctx, key).Result()
}

func (rs *RedisService) Set(key string, value interface{}, expiration time.Duration) error {
	return rs.client.Set(rs.ctx, key, value, expiration).Err()
}

func (rs *RedisService) Delete(key string) error {
	return rs.client.Del(rs.ctx, key).Err()
}
