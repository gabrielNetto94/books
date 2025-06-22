package userrepository

import (
	"books/internal/core/domain"
	"books/internal/core/repositories/cache"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db    *gorm.DB
	cache cache.CacheRepositoryInterface
}

func NewUserRepository(db *gorm.DB, cache cache.CacheRepositoryInterface) *UserRepositoryImpl {
	return &UserRepositoryImpl{db, cache}
}

type UserRepository interface {
	Save(user domain.User) error
	// FindById(id string) (domain.User, error)
	// ListAll() ([]domain.User, error)
	// Update(user domain.User) error
}

func (s *UserRepositoryImpl) Save(user domain.User) error {

	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	_ = s.cache.Set(user.Id, user)
	return nil
}
