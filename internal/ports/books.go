package ports

import "books/internal/core/domain"

type BookServiceInterface interface {
	FindById(bookId string) (domain.Book, *domain.DomainError)
	CreateBook(book domain.Book) *domain.DomainError
	UpdateBook(bookId string, book domain.Book) *domain.DomainError
	ListAll() ([]domain.Book, *domain.DomainError)
}
