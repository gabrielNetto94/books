package bookhandler

import (
	"books/internal/core/services"
	"books/internal/infra/log"
)

type BookHandlers struct {
	service *services.BookService
	log     log.Logger
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service *services.BookService, log log.Logger) *BookHandlers {
	return &BookHandlers{service, log}
}
