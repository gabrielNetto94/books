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
	if err := book.Validate(); err != nil {

		httpreponse.BadRequest(ctx.Writer, err.Error())
		return

	}
	err := b.service.CreateBook(book)
	if err != nil {
		httpreponse.InternalServerError(ctx.Writer, httpreponse.InternalError{Message: "Error on create book", Error: err})
		return
	}

	ctx.Status(http.StatusNoContent)
}
