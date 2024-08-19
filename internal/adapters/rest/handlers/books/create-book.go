package bookshandler

import (
	httpreponse "books/internal/adapters/rest/http-reponse"
	"books/internal/core/domain"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/db"
	"books/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	db := db.ConnectDatabase()

	repo := bookrepository.NewBookRepository(db)
	bookService := services.NewBookService(repo)

	var book domain.Book
	ctx.BindJSON(&book)
	err := bookService.CreateBook(book)
	if err != nil {
		httpreponse.InternalServerError(ctx.Writer, httpreponse.InternalError{Message: "Error on create book", Error: err})
		return
	}

	ctx.Status(http.StatusNoContent)
}
