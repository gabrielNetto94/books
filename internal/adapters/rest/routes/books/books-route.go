package booksroute

import (
	bookshandler "books/internal/adapters/rest/handlers/books"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/cache"
	"books/internal/core/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadBooksRoute(router *gin.RouterGroup, db *gorm.DB, cache *cache.CacheRepository) {

	repo := bookrepository.NewBookRepository(db, cache)
	service := services.NewBookService(repo)
	bookHandlers := bookshandler.NewBookHandlers(service)

	router.GET("/books", bookHandlers.ListBooks)
	router.GET("/books/:id", bookHandlers.GetBookById)
	router.POST("/books", bookHandlers.CreateBook)
	router.PUT("/books/:id", bookHandlers.UpdateBook)
}
