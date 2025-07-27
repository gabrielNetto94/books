package services

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/infra/log"
	"books/internal/ports"
	"books/pkg/observability"
	"context"

	"github.com/google/uuid"
)

type BookService struct {
	repo   bookrepository.BookRepository
	log    log.Logger
	tracer observability.Observability
}

func NewBookService(repo bookrepository.BookRepository, log log.Logger, tracer observability.Observability) ports.BookServiceInterface {
	return &BookService{repo, log, tracer}
}

func (s *BookService) FindById(ctx context.Context, bookId string) (domain.Book, *domain.DomainError) {
	s.log.Info("Finding book by ID: ", bookId)
	book, err := s.repo.FindById(ctx, bookId)
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

func (s *BookService) CreateBook(ctx context.Context, book domain.Book) *domain.DomainError {

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

	if err := s.repo.Save(ctx, book); err != nil {
		s.log.Error("Failed to save book: ", err)
		return &domain.DomainError{
			Message: "failed to save book",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}

	return nil
}

func (s *BookService) UpdateBook(ctx context.Context, bookId string, book domain.Book) *domain.DomainError {

	book.Id = bookId
	if err := book.Validate(); err != nil {
		s.log.Error("Book validation failed: ", err)
		return &domain.DomainError{
			Message: "book validation failed",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		}
	}

	err := s.repo.Update(ctx, book)
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

func (s *BookService) ListAll(ctx context.Context) ([]domain.Book, *domain.DomainError) {
	s.log.Info("[SERVICE] Listing all books")
	books, err := s.repo.ListAll(ctx)
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
