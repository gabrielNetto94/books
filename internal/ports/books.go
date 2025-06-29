package ports

import (
	"books/internal/core/domain"
	"context"
)

type BookServiceInterface interface {
	FindById(bookId string) (domain.Book, *domain.DomainError)
	CreateBook(book domain.Book) *domain.DomainError
	UpdateBook(bookId string, book domain.Book) *domain.DomainError
	ListAll(ctx context.Context) ([]domain.Book, *domain.DomainError)
}

type UserServiceInterface interface {
	CreateUser(user domain.User) (string, *domain.DomainError)
}
