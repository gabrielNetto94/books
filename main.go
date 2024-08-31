package main

import (
	"books/internal/adapters/rest/routes"
	"books/internal/config/logger"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.ConnectDatabase()
	cache := cache.ConnectCache()

	router := gin.Default()

	routes.InitRoutes(router, db, cache)

	if err := router.Run(":3000"); err != nil {
		logger.Log.Fatal("Error running server: ", err.Error())
	}
	logger.Log.Info("Server running on port 3000")
}
