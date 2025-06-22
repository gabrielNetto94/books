package main

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"
	"books/internal/adapters/http/routes"
	loglevel "books/internal/infra/log"
	"books/internal/infra/log/logrus"
	"os"

	bookrepository "books/internal/core/repositories/book"
	cache "books/internal/core/repositories/cache/redis"
	"books/internal/core/repositories/db"
	userrepository "books/internal/core/repositories/user"
	"books/internal/core/services"
	"books/pkg/env"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

func main() {

	db := db.ConnectDatabase("postgres://postgres:password@db:5432")
	cache := cache.ConnectCache("redis://cache:6379")

	log := setupLogger()

	bookRepo := bookrepository.NewBookRepository(db, cache)
	service := services.NewBookService(bookRepo, log)
	bookHandler := bookhandler.NewBookHandlers(service, log)

	userRepo := userrepository.NewUserRepository(db, cache)
	userService := services.NewUserService(userRepo, log)
	userHandler := userhandler.NewUserHandlers(userService, log)

	router := routes.InitRouter(bookHandler, userHandler)

	if err := router.Run(REST_API_PORT); err != nil {
		log.Fatal("Error running server: ", err.Error())
	}

}

func setupLogger() loglevel.Logger {
	log := logrus.NewLogrusAdapter()

	envName := os.Getenv("RAILWAY_ENVIRONMENT_NAME")
	log.Info("RAILWAY_ENVIRONMENT_NAME:", envName)

	if envName == "production" {
		log.SetLevel(loglevel.ErrorLevel)
		return log
	}

	envLoader := env.NewLoader()
	if err := envLoader.Load(); err != nil {
		log.Fatal("Error loading env: ", err.Error())
	}

	if envLoader.IsProduction() {
		log.SetLevel(loglevel.ErrorLevel)
	} else {
		log.SetLevel(loglevel.InfoLevel)
	}

	return log
}
