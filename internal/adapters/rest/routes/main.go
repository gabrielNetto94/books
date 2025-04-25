package routes

import (
	bookhandler "books/internal/adapters/rest/handlers/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(h *bookhandler.BookHandlers) *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/books", h.ListBooks)
	r.GET("/books/:id", h.GetBookById)
	r.POST("/books", h.CreateBook)
	r.PUT("/books/:id", h.UpdateBook)
	return r
}
