package routes

import (
	booksroute "books/internal/adapters/rest/routes/books"
	"books/internal/core/repositories/cache"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB, cache *cache.CacheRepository) {

	router.GET("/ping", Pong)

	v1Group := router.Group("/v1")
	booksroute.LoadBooksRoute(v1Group, db, cache)
}

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
