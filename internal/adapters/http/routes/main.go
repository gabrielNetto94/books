package routes

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router with all routes.
func InitRouter(bookHandler bookhandler.BookHTTPHandler, userHandler userhandler.UserHTTPHandler) *gin.Engine {

	r := gin.Default()
	// Health check
	r.GET("/ping", healthCheckHandler)

	// Book routes grouped by version
	v1 := r.Group("/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("", gin.WrapF(bookHandler.ListBooks))
			books.GET("/:id", gin.WrapF(bookHandler.GetBookById))
			books.POST("", gin.WrapF(bookHandler.CreateBook))
			books.PUT("/:id", gin.WrapF(bookHandler.UpdateBook))
		}
		users := v1.Group("/users")
		{
			users.POST("", gin.WrapF(userHandler.CreateUser))
		}
	}

	return r
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
