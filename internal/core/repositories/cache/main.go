package cache

import (
	"books/internal/config/env"

	"context"
	"encoding/json"

	"errors"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	cache *redis.Client
}

var ctx = context.Background()

func NewCacheInstance(cache *redis.Client) *CacheRepository {
	return &CacheRepository{cache}
}
func ConnectCache() *CacheRepository {

	opt, err := redis.ParseURL(env.GetVariable("CACHE_URL"))
	if err != nil {
		//logger.Log.Fatal("Error init cache", err.Error())
	}

	// Create client as usually.
	rdb := redis.NewClient(opt)
	status := rdb.Ping(ctx)

	if status.Err() != nil {
		//logger.Log.Fatal("Error init cache", status.Err().Error())
	}
	return &CacheRepository{rdb}
}

func (c CacheRepository) Get(key string, obj any) error {

	val, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("key dos not exists")
	} else if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), &obj)
}

func (c CacheRepository) Set(key string, value any) error {
	return c.cache.Set(ctx, key, value, 0).Err()
}
func (c CacheRepository) Del(key string) error {
	return c.cache.Del(ctx, key).Err()
}
