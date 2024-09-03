package cache

import (
	"books/internal/config/env"
	"context"
	"encoding/json"
	"log"

	"errors"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	cache *redis.Client
}

func NewCacheInstance(cache *redis.Client) *CacheRepository {
	return &CacheRepository{cache}
}
func ConnectCache() *CacheRepository {

	opt, err := redis.ParseURL(env.GetVariable("CACHE_URL"))
	if err != nil {
		log.Fatal("Error init cache", err.Error())
	}

	// Create client as usually.
	rdb := redis.NewClient(opt)
	status := rdb.Ping(context.Background())

	if status.Err() != nil {
		log.Fatal("Error init cache")
	}
	return &CacheRepository{rdb}
}

func (c CacheRepository) Get(key string, obj any) error {

	var ctx = context.Background()
	val, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("key dos not exists")
	} else if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), &obj)
}

func (c CacheRepository) Set(key string, value any) error {

	var ctx = context.Background()
	return c.cache.Set(ctx, key, value, 0).Err()
}
func (c CacheRepository) Del(key string) error {

	var ctx = context.Background()
	return c.cache.Del(ctx, key).Err()
}
