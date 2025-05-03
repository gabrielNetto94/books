package cache

type CacheRepositoryInterface interface {
	Get(key string, obj any) error
	Set(key string, value any) error
	Del(key string) error
}
