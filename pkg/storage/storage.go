package storage

import (
	"context"
	"io"
	"time"
)

type Storage interface {
	Put(ctx context.Context, key string, tags map[string]string, r io.Reader, size int64, contentType string) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	URL(ctx context.Context, key string, expiry time.Duration) (string, error)
}
