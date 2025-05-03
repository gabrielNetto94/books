package bookhandler

import (
	"books/internal/infra/log"
	"books/internal/ports"
)

type BookHandlers struct {
	service ports.BookServiceInterface
	log     log.Logger
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service ports.BookServiceInterface, log log.Logger) *BookHandlers {
	return &BookHandlers{service, log}
}
