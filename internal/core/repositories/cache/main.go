package cache

import (
	"books/internal/pkg/env"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.GetVariable("CACHE_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	status := rdb.Ping(context.Background())

	if status.Err() != nil {
		log.Fatal("Error init cache")
	}
	return &CacheRepository{rdb}
}

func (c CacheRepository) Get(key string) (string, error) {

	var ctx = context.Background()
	val, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key dos not exists")
	} else if err != nil {
		return "", err
	}

	return val, nil
}
func (c CacheRepository) GetObject(key string, obj any) error {

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
	err := c.cache.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
