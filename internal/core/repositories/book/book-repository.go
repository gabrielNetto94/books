package bookrepository

import (
	"books/internal/core/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

type Bookrepository interface {
	Save(book domain.Book) error
	FindById(id string) (domain.Book, error)
	ListAll() ([]domain.Book, error)
}

func (s *UserRepositoryImpl) Save(book domain.Book) error {
	return s.db.Create(book).Error
}

func (s *UserRepositoryImpl) FindById(id string) (domain.Book, error) {
	var book = domain.Book{Id: id}
	return book, s.db.Find(&book).Error
}

func (s *UserRepositoryImpl) ListAll() ([]domain.Book, error) {

	var books []domain.Book

	return books, s.db.Find(&books).Error
}
