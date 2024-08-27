package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) UpdateBook(ctx *gin.Context) {

	var book domain.Book
	id := ctx.Param("id")
	if err := ctx.BindJSON(&book); err != nil {
		httpreponse.ErrorResponse(ctx, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	if serviceErr := b.service.UpdateBook(id, book); serviceErr.Error != nil {
		httpreponse.ErrorResponse(ctx, serviceErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
