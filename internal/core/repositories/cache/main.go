package cache

import "context"

type CacheRepositoryInterface interface {
	Get(ctx context.Context, key string, obj any) error
	Set(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) error
}
