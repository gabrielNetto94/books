package main

import (
	"books/internal/adapters/grpc/books/handler"
	pb "books/internal/adapters/grpc/books/proto"
	"books/internal/adapters/rest/routes"
	"books/internal/config/logger"
	bookrepository "books/internal/core/repositories/book"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"books/internal/core/services"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

func main() {

	db := db.ConnectDatabase()
	cache := cache.ConnectCache()

	repo := bookrepository.NewBookRepository(db, cache)
	service := services.NewBookService(repo)

	go initRestAPi(service)
	initGrpcServer(service)

}

func initRestAPi(service *services.BookService) {
	router := gin.Default()
	routes.InitRoutes(router, service)
	if err := router.Run(REST_API_PORT); err != nil {
		logger.Log.Fatal("Error running server: ", err.Error())
	}
	logger.Log.Info("Server running on port " + REST_API_PORT)
}

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
