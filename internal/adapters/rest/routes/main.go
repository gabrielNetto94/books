package routes

import (
	"books/internal/adapters/rest/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	routes := gin.Default()

	v1Group := routes.Group("v1")
	v1Group.GET("/books", handlers.ListBooks)
	return routes
}
