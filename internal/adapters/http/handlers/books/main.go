package bookhandler

import (
	httputils "books/internal/adapters/http/http-utils"
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"books/internal/infra/log"
	"books/internal/ports"
	"books/pkg/observability"
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
	tracer  observability.Tracer
}

// NewBookHandlers cria uma nova instância de BookHandlers.
func NewBookHandlers(service ports.BookServiceInterface, log log.Logger, tracer observability.Tracer) BookHTTPHandler {
	return &BookHandlers{service, log, tracer}
}

func (b BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx, span := b.tracer.Span(r.Context(), "BookHandler.CreateBook")
	defer span.End()

	//@todo refactor to use DTO isntead domain model
	var book domain.Book

	b.log.Info("CreateBook called")
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		span.RecordError(err)
		span.SetAttribute("error", true)
		httputils.HandleError(w, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}
	span.SetAttribute("book", book)

	if err := b.service.CreateBook(ctx, book); err != nil {
		span.RecordError(err.Error)
		span.SetAttribute("error", true)
		b.log.Error("Error creating book: ", err.Error.Error())
		httputils.HandleError(w, *err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (b BookHandlers) ListBooks(w http.ResponseWriter, r *http.Request) {
	ctx, span := b.tracer.Span(r.Context(), "BookHandler.ListBooks")
	defer span.End()

	b.log.Info("ListBooks called")
	books, err := b.service.ListAll(ctx)
	if err != nil {
		span.RecordError(err.Error)
		span.SetAttribute("error", true)
		b.log.Error("Error listing books: ", err.Error)
		httputils.HandleError(w, *err)
		return
	}
	span.SetAttribute("books.count", len(books))
	httputils.JsonResponse(w, http.StatusOK, books)
}

func (b BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request) {
	ctx, span := b.tracer.Span(r.Context(), "BookHandler.GetBookById")
	defer span.End()

	b.log.Info("GetBookById called")

	id := r.URL.Query().Get("id")
	span.SetAttribute("book.id", id)

	book, err := b.service.FindById(ctx, id)
	if err != nil {
		span.RecordError(err.Error)
		span.SetAttribute("error", true)
		b.log.Error("Error getting book by ID: ", err.Error)
		httputils.HandleError(w, *err)
		return
	}

	httputils.JsonResponse(w, http.StatusOK, book)
}

func (b BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {

	ctx, span := b.tracer.Span(r.Context(), "BookHandler.UpdateBook")
	defer span.End()

	b.log.Info("UpdateBook called")
	id := r.URL.Query().Get("id")
	span.SetAttribute("book.id", id)
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		span.RecordError(err)
		span.SetAttribute("error", true)
		b.log.Error("Error decoding request body: ", err.Error())
		httputils.HandleError(w, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if bookError := b.service.UpdateBook(ctx, id, book); bookError != nil {
		span.RecordError(bookError.Error)
		span.SetAttribute("error", true)
		b.log.Error("Error updating book: ", bookError.Error.Error())
		httputils.HandleError(w, *bookError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
