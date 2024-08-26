package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	"books/internal/core/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) CreateBook(ctx *gin.Context) {

	var book domain.Book
	if err := ctx.BindJSON(&book); err != nil {
		httpreponse.BadRequest(ctx.Writer, "Invalid request")
		return
	}
	if bookError := b.service.CreateBook(book); bookError.Error != nil {
		httpreponse.ErrorResponse(ctx.Writer, httpreponse.ErrorResponseModel{
			Message: "Error on create book",
			Error:   bookError.Error,
			Code:    bookError.Code,
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
