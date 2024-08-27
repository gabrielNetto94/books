package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) GetBookById(ctx *gin.Context) {

	id := ctx.Param("id")
	book, serviceErr := b.service.FindById(id)
	if serviceErr.Error != nil {
		httpreponse.ErrorResponse(ctx, serviceErr)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
