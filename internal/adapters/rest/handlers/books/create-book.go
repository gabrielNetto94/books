package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) CreateBook(ctx *gin.Context) {

	var book domain.Book
	if err := ctx.BindJSON(&book); err != nil {
		httpreponse.ErrorResponse(ctx, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}
	if bookError := b.service.CreateBook(book); bookError.Error != nil {
		httpreponse.ErrorResponse(ctx, bookError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
