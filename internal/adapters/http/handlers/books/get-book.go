package bookhandler

import (
	httpreponse "books/internal/adapters/http/http-reponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) GetBookById(ctx *gin.Context) {
	b.log.Info("GetBookById called")
	id := ctx.Param("id")
	book, serviceErr := b.service.FindById(id)
	if serviceErr.Error != nil {
		b.log.Error("Error getting book by ID: ", serviceErr.Error)
		httpreponse.ErrorResponse(ctx, serviceErr)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
