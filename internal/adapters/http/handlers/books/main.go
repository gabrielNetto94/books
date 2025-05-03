package bookhandler

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"books/internal/infra/log"
	"books/internal/ports"
	"encoding/json"
	"net/http"
)

type BookHTTPHandler interface {
	ListBooks(w http.ResponseWriter, r *http.Request)
	GetBookById(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type BookHandlers struct {
	service ports.BookServiceInterface
	log     log.Logger
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service ports.BookServiceInterface, log log.Logger) BookHTTPHandler {
	return &BookHandlers{service, log}
}

func (b BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if bookError := b.service.CreateBook(book); bookError.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(bookError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (b BookHandlers) ListBooks(w http.ResponseWriter, r *http.Request) {

	b.log.Info("ListBooks called 2")
	books, err := b.service.ListAll()
	if err != nil {
		b.log.Error("Error listing books: ", err.Error)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
func (b BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request) {
	b.log.Info("GetBookById called")
	id := r.URL.Query().Get("id")
	book, err := b.service.FindById(id)
	if err != nil {
		b.log.Error("Error getting book by ID: ", err.Error)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (b BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	b.log.Info("UpdateBook called")
	id := r.URL.Query().Get("id")
	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if bookError := b.service.UpdateBook(id, book); bookError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(bookError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
