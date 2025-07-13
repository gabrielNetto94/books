package cache

import (
	"books/internal/core/repositories/cache"
	"fmt"
	"log"

	"context"
	"encoding/json"

	"errors"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	cache *redis.Client
}

func NewCacheInstance(dsn string) cache.CacheRepositoryInterface {

	redisClient := connectCache(dsn)
	return &CacheRepository{redisClient}
}
func connectCache(url string) *redis.Client {

	opt, err := redis.ParseURL(url)
	if err != nil {
		fmt.Println("Error init cache", err.Error())
		//logger.Log.Fatal("Error init cache", err.Error())
	}

	// Create client as usually.
	rdb := redis.NewClient(opt)
	status := rdb.Ping(context.Background())

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	if status.Err() != nil {
		log.Fatal("Error init cache", status.Err().Error())
		//logger.Log.Fatal("Error init cache", status.Err().Error())
	}
	return rdb
}

func (c CacheRepository) Get(ctx context.Context, key string, obj any) error {

	val, err := c.cache.Get(ctx, key).Result()

	fmt.Println("val ->", val)
	if err != nil {
		if err == redis.Nil {
			return errors.New("key does not exist")
		}
		return err
	}

	return json.Unmarshal([]byte(val), &obj)
}

func (c CacheRepository) Set(ctx context.Context, key string, value any) error {
	return c.cache.Set(ctx, key, value, 0).Err()
}
func (c CacheRepository) Del(ctx context.Context, key string) error {
	return c.cache.Del(ctx, key).Err()
}
