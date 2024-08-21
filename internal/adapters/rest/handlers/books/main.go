package bookhandler

import "books/internal/core/services"

type BookHandlers struct {
	service *services.BookService
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service *services.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}
