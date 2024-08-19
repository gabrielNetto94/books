package bookshandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/db"
	"books/internal/core/services"

	"github.com/gin-gonic/gin"
)

func ListBooks(ctx *gin.Context) {
	db := db.ConnectDatabase()

	repo := bookrepository.NewBookRepository(db)
	bookService := services.NewBookService(repo)

	books, err := bookService.ListAll()
	if err != nil {
		httpreponse.InternalServerError(ctx.Writer, httpreponse.InternalError{Message: "Error on create book", Error: err})
		return
	}

	ctx.JSON(200, books)
}
