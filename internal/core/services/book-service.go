package services

import (
	errorscode "books/internal/consts/errors-code"
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

func (s *BookService) CreateBook(book domain.Book) domain.BookError {

	if err := book.Validate(); err != nil {
		return domain.BookError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	book.Id = uuid.New().String()
	if err := s.repo.Save(book); err != nil {
		return domain.BookError{
			Message: "failed to save book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}

	return domain.BookError{}
}

func (s *BookService) UpdateBook(book domain.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) ListAll() ([]domain.Book, error) {
	return s.repo.ListAll()
}
