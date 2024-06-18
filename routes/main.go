package routes

import (
	httpresponse "books/internal/http-response"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		httpresponse.HandleResponse(ctx.Writer, 200, map[string]string{"message": "hello world"})
	})

	router.GET("/error", func(ctx *gin.Context) {
		httpresponse.HandleError(ctx.Writer, 500, "Error message", nil)
	})

	router.GET("/no-content", func(ctx *gin.Context) {
		httpresponse.HandleResponseNoContent(ctx.Writer)
	})

	return router
}
