package services

import (
	"books/internal/core/domain"
	bookrepository "books/internal/core/repositories/book"

	"github.com/google/uuid"
)

type BookService struct {
	repo bookrepository.Bookrepository
}

func NewBookService(repo bookrepository.Bookrepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) FindById(bookId string) (domain.Book, error) {
	return s.repo.FindById(bookId)
}

func (s *BookService) CreateBook(book domain.Book) error {
	book.Id = uuid.New().String()
	return s.repo.Save(book)
}

func (s *BookService) ListAll() ([]domain.Book, error) {
	return s.repo.ListAll()
}
