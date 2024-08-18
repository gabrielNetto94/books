package booksroute

import (
	bookshandler "books/internal/adapters/rest/handlers/books"

	"github.com/gin-gonic/gin"
)

func LoadBooksRoute(router *gin.RouterGroup) *gin.RouterGroup {

	router.GET("/books", bookshandler.ListBooks)
	return router
}
