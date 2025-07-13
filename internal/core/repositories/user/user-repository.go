package userrepository

import (
	"books/internal/core/domain"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"context"
)

type UserRepositoryImpl struct {
	db    db.DatabaseRepositoryInterface
	cache cache.CacheRepositoryInterface
}

func NewUserRepository(db db.DatabaseRepositoryInterface, cache cache.CacheRepositoryInterface) *UserRepositoryImpl {
	return &UserRepositoryImpl{db, cache}
}

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	// FindById(id string) (domain.User, error)
	// ListAll() ([]domain.User, error)
	// Update(user domain.User) error
}

func (s *UserRepositoryImpl) Save(ctx context.Context, user domain.User) error {

	if err := s.db.Create(ctx, user); err != nil {
		return err
	}
	_ = s.cache.Set(context.Background(), user.Id, user) //@TODO AJUSTAR CONTEXT
	return nil
}
