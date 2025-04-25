package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) ListBooks(ctx *gin.Context) {

	books, serviceErr := b.service.ListAll()
	if serviceErr != nil {
		httpreponse.ErrorResponse(ctx, *serviceErr)
		return
	}

	ctx.JSON(200, books)
}
