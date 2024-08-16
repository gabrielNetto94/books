package handlers

import (
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
		ctx.JSON(500, err.Error())
	}

	ctx.JSON(200, books)
}
