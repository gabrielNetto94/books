package ports

import (
	"books/internal/core/domain"
	"context"
)

type BookServiceInterface interface {
	FindById(ctx context.Context, bookId string) (domain.Book, *domain.DomainError)
	CreateBook(ctx context.Context, book domain.Book) *domain.DomainError
	UpdateBook(ctx context.Context, bookId string, book domain.Book) *domain.DomainError
	ListAll(ctx context.Context) ([]domain.Book, *domain.DomainError)
}

type UserServiceInterface interface {
	CreateUser(user domain.User) (string, *domain.DomainError)
}
