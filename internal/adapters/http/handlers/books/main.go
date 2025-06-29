package bookhandler

import (
	httputils "books/internal/adapters/http/http-utils"
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

	b.log.Info("CreateBook called")
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		httputils.HandleError(w, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if err := b.service.CreateBook(book); err != nil {
		b.log.Error("Error creating book: ", err.Error.Error())
		httputils.HandleError(w, *err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (b BookHandlers) ListBooks(w http.ResponseWriter, r *http.Request) {

	b.log.Info("ListBooks called")
	books, err := b.service.ListAll(r.Context())
	if err != nil {
		b.log.Error("Error listing books: ", err.Error)
		httputils.HandleError(w, *err)
		return
	}
	httputils.JsonResponse(w, http.StatusOK, books)
}

func (b BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request) {

	b.log.Info("GetBookById called")

	id := r.URL.Query().Get("id")
	book, err := b.service.FindById(id)
	if err != nil {
		b.log.Error("Error getting book by ID: ", err.Error)
		httputils.HandleError(w, *err)
		return
	}

	httputils.JsonResponse(w, http.StatusOK, book)
}

func (b BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {

	b.log.Info("UpdateBook called")
	id := r.URL.Query().Get("id")
	var book domain.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		b.log.Error("Error decoding request body: ", err.Error())
		httputils.HandleError(w, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if bookError := b.service.UpdateBook(id, book); bookError != nil {
		b.log.Error("Error updating book: ", bookError.Error.Error())
		httputils.HandleError(w, *bookError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
