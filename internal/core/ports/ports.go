package ports

import "books/internal/core/domain"

type MessengerService interface {
	CreateBook(userID string, message domain.Book) error
}

type MessengerRepository interface {
	CreateBook(userID string, message domain.Book) error
}
