package routes

import (
	httpresponse "books/internal/infra/web/http-response"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		httpresponse.HandleResponse(ctx.Writer, 200, map[string]string{"message": "hello world"})
	})

	return router
}
