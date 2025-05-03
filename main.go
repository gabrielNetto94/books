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

	log := logrus.NewLogrusAdapter()

	log.Info("RAILWAY_ENVIRONMENT_NAME:", os.Getenv("RAILWAY_ENVIRONMENT_NAME"))

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") == "production" {
		log.SetLevel(loglevel.ErrorLevel)
	} else {

		envv := env.NewLoader()
		if err := envv.Load(); err != nil {
			log.Fatal("Error loading env: ", err.Error())
		}

		if envv.IsProduction() {
			log.SetLevel(loglevel.ErrorLevel)
		} else {
			log.SetLevel(loglevel.InfoLevel)
		}
	}

	//db := db.ConnectDatabase("postgres://postgres:password@db:5432")
	//cache := cache.ConnectCache("redis://cache:6379")

	//repo := bookrepository.NewBookRepository(db, cache2)
	repo := bookmock.NewBookRepositoryMock(datafake.NewFaker())

	service := services.NewBookService(repo, log)
	bookHandler := bookhandler.NewBookHandlers(service, log)
	router := routes.InitRouter(bookHandler)

	if err := router.Run(REST_API_PORT); err != nil {
		log.Fatal("Error running server: ", err.Error())
	}

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	logger.Log.Info("Starting REST API...")
	// 	// initRestAPi(service)
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	logger.Log.Info("Starting gRPC server...")
	// 	initGrpcServer(service)
	// }()

	// wg.Wait()

}

// func initGrpcServer(service *services.BookService) {

// 	lis, err := net.Listen("tcp", GRPC_SERVER_PORT)
// 	if err != nil {
// 		logger.Log.Fatal("Error listen: " + err.Error())
// 	}

// 	server := grpc.NewServer()
// 	pb.RegisterBookServiceServer(server, &handler.Server{
// 		Service: service,
// 	})

// 	if err := server.Serve(lis); err != nil {
// 		logger.Log.Fatal("Error serve: " + err.Error())
// 	}
// 	logger.Log.Info("Server running on port " + GRPC_SERVER_PORT)

// }
