package bookhandler

import (
	"books/internal/core/services"
	"books/internal/infra/logger"
)

type BookHandlers struct {
	service *services.BookService
	log     logger.Logger
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service *services.BookService, log logger.Logger) *BookHandlers {
	return &BookHandlers{service, log}
}
