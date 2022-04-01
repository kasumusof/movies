package utils

import (
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	Host     string
	ID       int
	Password string
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: cache.Password,
		DB:       cache.ID,
	})
}

func (cache *redisCache) Get(key string) (string, error) {
	client := cache.getClient()
	result, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (cache *redisCache) Set(key string, value interface{}, exp time.Duration) error {
	client := cache.getClient()
	err := client.Set(key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cache *redisCache) Del(key string) error {
	client := cache.getClient()
	err := client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewCache(host, password string, id int) *redisCache {
	return &redisCache{
		Host:     host,
		ID:       id,
		Password: password,
	}
}
