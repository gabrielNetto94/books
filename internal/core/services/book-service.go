package services

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/infra/log"
	"books/internal/ports"

	"github.com/google/uuid"
)

type BookService struct {
	repo bookrepository.BookRepository
	log  log.Logger
}

func NewBookService(repo bookrepository.BookRepository, log log.Logger) ports.BookServiceInterface {
	return &BookService{repo, log}
}

func (s *BookService) FindById(bookId string) (domain.Book, *domain.DomainError) {
	book, err := s.repo.FindById(bookId)
	if err != nil {
		s.log.Error("Failed to find book by ID: ", err)
		return domain.Book{}, &domain.DomainError{
			Message: "book not found",
			Code:    errorscode.ErrNotFound,
			Error:   err,
		}
	}
	return book, nil
}

func (s *BookService) CreateBook(book domain.Book) *domain.DomainError {

	s.log.Info("Creating book: ", book)
	book.Id = uuid.New().String()
	if err := book.Validate(); err != nil {
		s.log.Error("Book validation failed: ", err)
		return &domain.DomainError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	if err := s.repo.Save(book); err != nil {
		s.log.Error("Failed to save book: ", err)
		return &domain.DomainError{
			Message: "failed to save book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}

	return nil
}

func (s *BookService) UpdateBook(bookId string, book domain.Book) *domain.DomainError {

	book.Id = bookId
	if err := book.Validate(); err != nil {
		s.log.Error("Book validation failed: ", err)
		return &domain.DomainError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	err := s.repo.Update(book)
	if err != nil {
		s.log.Error("Failed to update book: ", err)
		return &domain.DomainError{
			Message: "failed to update book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return nil
}

func (s *BookService) ListAll() ([]domain.Book, *domain.DomainError) {

	books, err := s.repo.ListAll()
	if err != nil {
		s.log.Error("Failed to list all books: ", err)
		return nil, &domain.DomainError{
			Message: "failed to list all books",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return books, nil
}
