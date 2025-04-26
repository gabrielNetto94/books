package main

import (
	bookhandler "books/internal/adapters/rest/handlers/books"
	"books/internal/adapters/rest/routes"
	"books/internal/infra/logger"

	bookmock "books/internal/core/repositories/book-mock"
	"books/internal/core/services"
	datafake "books/pkg/data-fake"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

func main() {

	// db := db.ConnectDatabase()
	// cache := cache.ConnectCache()
	// repo := bookrepository.NewBookRepository(db, cache)
	repo := bookmock.NewBookRepositoryMock(datafake.NewFaker())
	log := logger.NewLogrusAdapter()

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
