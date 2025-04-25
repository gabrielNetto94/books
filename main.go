package main

import (
	"books/internal/adapters/grpc/books/handler"
	pb "books/internal/adapters/grpc/books/proto"
	bookhandler "books/internal/adapters/rest/handlers/books"
	"books/internal/adapters/rest/routes"
	"books/internal/config/logger"
	bookmock "books/internal/core/repositories/book-mock"
	"books/internal/core/services"
	"net"

	"google.golang.org/grpc"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

// repo := &InMemoryRepo{}
// 	service := services.NewBookService(repo)
// 	handler := rest.NewHandler(service)
// 	router := rest.InitRouter(handler)

func main() {

	// db := db.ConnectDatabase()
	// cache := cache.ConnectCache()

	// repo := bookrepository.NewBookRepository(db, cache)
	repo := bookmock.NewBookRepositoryMock()

	service := services.NewBookService(repo)
	bookHandler := bookhandler.NewBookHandlers(service)
	router := routes.InitRouter(bookHandler)

	if err := router.Run(REST_API_PORT); err != nil {
		logger.Log.Fatal("Error running server: ", err.Error())
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

// func initRestAPi(service *services.BookService) {
// 	router := gin.Default()
// 	routes.InitRoutes(router, service)
// 	if err := router.Run(REST_API_PORT); err != nil {
// 		logger.Log.Fatal("Error running server: ", err.Error())
// 	}
// 	logger.Log.Info("Server running on port " + REST_API_PORT)
// }

func initGrpcServer(service *services.BookService) {

	lis, err := net.Listen("tcp", GRPC_SERVER_PORT)
	if err != nil {
		logger.Log.Fatal("Error listen: " + err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterBookServiceServer(server, &handler.Server{
		Service: service,
	})

	if err := server.Serve(lis); err != nil {
		logger.Log.Fatal("Error serve: " + err.Error())
	}
	logger.Log.Info("Server running on port " + GRPC_SERVER_PORT)

}
