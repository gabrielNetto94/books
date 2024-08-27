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

func (s *BookService) FindById(bookId string) (domain.Book, domain.DomainError) {
	book, err := s.repo.FindById(bookId)
	if err != nil {
		return domain.Book{}, domain.DomainError{
			Message: "book not found",
			Code:    errorscode.ErrNotFound,
			Error:   err,
		}
	}
	return book, domain.DomainError{}
}

func (s *BookService) CreateBook(book domain.Book) domain.DomainError {

	if err := book.Validate(); err != nil {
		return domain.DomainError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	book.Id = uuid.New().String()
	if err := s.repo.Save(book); err != nil {
		return domain.DomainError{
			Message: "failed to save book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}

	return domain.DomainError{}
}

func (s *BookService) UpdateBook(bookId string, book domain.Book) domain.DomainError {

	book.Id = bookId
	if err := book.Validate(); err != nil {
		return domain.DomainError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	err := s.repo.Update(book)
	if err != nil {
		return domain.DomainError{
			Message: "failed to update book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return domain.DomainError{}
}

func (s *BookService) ListAll() ([]domain.Book, domain.DomainError) {

	books, err := s.repo.ListAll()
	if err != nil {
		return nil, domain.DomainError{
			Message: "failed to list all books",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return books, domain.DomainError{}
}
