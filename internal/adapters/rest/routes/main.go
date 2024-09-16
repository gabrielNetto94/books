package routes

import (
	booksroute "books/internal/adapters/rest/routes/books"
	"books/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, service *services.BookService) {

	router.GET("/ping", Pong)

	v1Group := router.Group("/v1")
	booksroute.LoadBooksRoute(v1Group, service)
}

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
