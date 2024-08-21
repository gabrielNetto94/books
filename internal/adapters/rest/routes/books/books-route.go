package booksroute

import (
	bookshandler "books/internal/adapters/rest/handlers/books"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/db"
	"books/internal/core/services"

	"github.com/gin-gonic/gin"
)

func LoadBooksRoute(router *gin.RouterGroup) *gin.RouterGroup {

	db := db.ConnectDatabase()
	repo := bookrepository.NewBookRepository(db)
	service := services.NewBookService(repo)
	bookHandlers := bookshandler.NewBookHandlers(service)

	router.GET("/books", bookHandlers.ListBooks)
	router.POST("/books", bookHandlers.CreateBook)

	return router
}
