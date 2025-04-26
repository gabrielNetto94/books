package bookhandler

import (
	"books/internal/config/logger"
	"books/internal/core/services"
)

type BookHandlers struct {
	service *services.BookService
	log     logger.Logger
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service *services.BookService, log logger.Logger) *BookHandlers {
	return &BookHandlers{service, log}
}
