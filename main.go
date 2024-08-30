package main

import (
	"books/internal/adapters/rest/routes"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.ConnectDatabase()
	cache := cache.ConnectCache()

	router := gin.Default()

	routes.InitRoutes(router, db, cache)

	if err := router.Run(":3000"); err != nil {
		log.Fatal("error running server:", err)
	}
}
