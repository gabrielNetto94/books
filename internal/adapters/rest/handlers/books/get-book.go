package bookhandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b BookHandlers) GetBookById(ctx *gin.Context) {

	id := ctx.Param("id")
	book, err := b.service.FindById(id)
	if err != nil {
		httpreponse.InternalServerError(ctx.Writer, httpreponse.InternalError{Message: "Error on create book", Error: err})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
