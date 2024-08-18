package routes

import (
	booksroute "books/internal/adapters/rest/routes/books"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	routes := gin.Default()

	v1Group := routes.Group("v1")
	booksroute.LoadBooksRoute(v1Group)
	return routes
}
