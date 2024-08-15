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
	// FindById(id string) (*domain.Book, error)
	// ListAll() ([]*domain.Book, error)
}

func (s *UserRepositoryImpl) Save(book domain.Book) error {
	return s.db.Create(book).Error
}

// func (s *UserRepositoryImpl) FindById(id string) (*domain.Book, error) {

// 	return &domain.Book{Id: "asd", Title: "asdf"}, nil
// }

// func (s *UserRepositoryImpl) ListAll() ([]*domain.Book, error) {

// 	return []*domain.Book{
// 		{Id: "ASF13asda", Title: "asdf 1"},
// 		{Id: "ASF13asda", Title: "asdf 2"},
// 	}, nil
// }
