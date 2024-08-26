package booksroute

import (
	bookshandler "books/internal/adapters/rest/handlers/books"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"books/internal/core/services"

	"github.com/gin-gonic/gin"
)

func LoadBooksRoute(router *gin.RouterGroup) *gin.RouterGroup {

	db := db.ConnectDatabase()
	cache := cache.ConnectCache()

	repo := bookrepository.NewBookRepository(db, cache)
	service := services.NewBookService(repo)
	bookHandlers := bookshandler.NewBookHandlers(service)

	router.GET("/books", bookHandlers.ListBooks)
	router.GET("/books/:id", bookHandlers.GetBookById)
	router.POST("/books", bookHandlers.CreateBook)
	router.PUT("/books/:id", bookHandlers.UpdateBook)

	return router
}
