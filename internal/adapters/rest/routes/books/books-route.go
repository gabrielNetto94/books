package booksroute

import (
	bookshandler "books/internal/adapters/rest/handlers/books"
	"books/internal/core/services"

	"github.com/gin-gonic/gin"
)

func LoadBooksRoute(router *gin.RouterGroup, service *services.BookService) {

	bookHandlers := bookshandler.NewBookHandlers(service)

	router.GET("/books", bookHandlers.ListBooks)
	router.GET("/books/:id", bookHandlers.GetBookById)
	router.POST("/books", bookHandlers.CreateBook)
	router.PUT("/books/:id", bookHandlers.UpdateBook)
}
