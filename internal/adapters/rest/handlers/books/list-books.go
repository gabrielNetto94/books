package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) ListBooks(ctx *gin.Context) {

	b.service.ListAll()

	books, err := b.service.ListAll()
	if err != nil {
		httpreponse.InternalServerError(ctx.Writer, httpreponse.InternalError{Message: "Error on create book", Error: err})
		return
	}

	ctx.JSON(200, books)
}
