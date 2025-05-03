package main

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	"books/internal/adapters/http/routes"
	loglevel "books/internal/infra/log"
	"books/internal/infra/log/logrus"
	"os"

	bookmock "books/internal/core/repositories/book-mock"
	"books/internal/core/services"
	datafake "books/pkg/data-fake"
	"books/pkg/env"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

func main() {

	//db := db.ConnectDatabase("postgres://postgres:password@db:5432")
	//cache := cache.ConnectCache("redis://cache:6379")

	//repo := bookrepository.NewBookRepository(db, cache2)
	log := setupLogger()

	repo := bookmock.NewBookRepositoryMock(datafake.NewFaker())

	service := services.NewBookService(repo, log)
	bookHandler := bookhandler.NewBookHandlers(service, log)
	router := routes.InitRouter(bookHandler)

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
