package routes

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(h *bookhandler.BookHandlers) *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "po2ng"})
	})
	r.GET("/v1/books", h.ListBooks)
	r.GET("/v1/books/:id", h.GetBookById)
	r.POST("/v1/books", h.CreateBook)
	r.PUT("/v1/books/:id", h.UpdateBook)
	return r
}
