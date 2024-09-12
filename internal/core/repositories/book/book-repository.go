package bookrepository

import (
	"books/internal/core/domain"
	"books/internal/core/repositories/cache"
	"encoding/json"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db    *gorm.DB
	cache *cache.CacheRepository
}

func NewBookRepository(db *gorm.DB, cache *cache.CacheRepository) *UserRepositoryImpl {
	return &UserRepositoryImpl{db, cache}
}

type Bookrepository interface {
	Save(book domain.Book) error
	FindById(id string) (domain.Book, error)
	ListAll() ([]domain.Book, error)
	Update(book domain.Book) error
}

func (s *UserRepositoryImpl) Save(book domain.Book) error {

	if err := s.db.Create(book).Error; err != nil {
		return err
	}
	_ = s.cache.Set(book.Id, book)
	return nil
}

func (s *UserRepositoryImpl) Update(book domain.Book) error {

	err := s.db.Save(book).Error
	if err == nil {
		_ = s.cache.Del(book.Id)
		return nil
	}

	return err
}

func (s *UserRepositoryImpl) FindById(id string) (domain.Book, error) {
	var book = domain.Book{Id: id}
	err := s.cache.Get(id, &book)
	if err == nil {
		return book, nil
	}
	resp := s.db.Find(&book)
	bookBytes, err := json.Marshal(book)
	if err == nil {
		_ = s.cache.Set(book.Id, bookBytes)
	}

	return book, resp.Error
}

func (s *UserRepositoryImpl) ListAll() ([]domain.Book, error) {

	var books []domain.Book

	return books, s.db.Find(&books).Error
}
