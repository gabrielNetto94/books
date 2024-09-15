package main

import (
	pb "books/internal/adapters/grpc/books/proto"
	"books/internal/adapters/rest/routes"
	"books/internal/config/logger"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	pb.BookServiceServer
}

func main() {

	db := db.ConnectDatabase()
	cache := cache.ConnectCache()

	initRestAPi(db, cache)
	//initGrpcServer()

}

func initRestAPi(db *gorm.DB, cache *cache.CacheRepository) {
	router := gin.Default()
	routes.InitRoutes(router, db, cache)
	if err := router.Run(":3000"); err != nil {
		logger.Log.Fatal("Error running server: ", err.Error())
	}
	logger.Log.Info("Server running on port 3000")
}

func initGrpcServer() {

	var addr string = ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterBookServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Erorr serve: %v", err)
	}
	logger.Log.Info("Server running on port 50051")

}
