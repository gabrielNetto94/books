package db

import (
	"context"
)

type DatabaseRepositoryInterface interface {
	Create(ctx context.Context, data any) error
	Find(ctx context.Context, data any) error
}
