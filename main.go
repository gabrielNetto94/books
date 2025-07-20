package main

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"
	"books/internal/adapters/http/routes"
	bookrepository "books/internal/core/repositories/book"
	cache "books/internal/core/repositories/cache/redis"
	"books/internal/core/repositories/db/gorm"
	userrepository "books/internal/core/repositories/user"
	"books/internal/core/services"

	"books/internal/infra/log"
	"books/internal/infra/log/logrus"
	"books/pkg/env"
	"books/pkg/observability/opentelemetry"
	"context"
	"os"
	"os/signal"

	// 1. Import your new metrics library

	// Assuming you use Gin
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const REST_API_PORT = ":3005"
const GRPC_SERVER_PORT = ":3001"

var serviceName = semconv.ServiceNameKey.String("asdf-test")

func main() {
	log := setupLogger()

	db := gorm.NewDatabaseRepository("postgres://postgres:password@db:5432")
	cacheRepo := cache.NewCacheInstance("redis://cache:6379")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	conn, err := opentelemetry.InitConn()
	if err != nil {
		log.Fatal("Error creating gRPC connection: ", err.Error())
	}

	res, err := resource.New(ctx, resource.WithAttributes(serviceName))
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := opentelemetry.InitTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal("Error initializing tracer provider: ", err.Error())
	}
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatal("Error shutting down tracer provider: ", err.Error())
		}
	}()

	opentelemetry := opentelemetry.NewObservability(serviceName.Value.AsString())

	bookRepo := bookrepository.NewBookRepository(db, cacheRepo)
	service := services.NewBookService(bookRepo, log, opentelemetry)
	bookHandler := bookhandler.NewBookHandlers(service, log)

	userRepo := userrepository.NewUserRepository(db, cacheRepo)
	userService := services.NewUserService(userRepo, log)
	userHandler := userhandler.NewUserHandlers(userService, log)

	router := routes.InitRouter(bookHandler, userHandler)

	if err := router.Run(REST_API_PORT); err != nil {
		log.Fatal("Error running server: ", err.Error())
	}
}

// metricsMiddleware creates a gin.HandlerFunc to collect standard HTTP metrics.

func setupLogger() log.Logger {
	logrusLog := logrus.NewLogrusAdapter()

	envName := os.Getenv("RAILWAY_ENVIRONMENT_NAME")
	logrusLog.Info("RAILWAY_ENVIRONMENT_NAME:", envName)

	if envName == "production" {
		logrusLog.SetLevel(log.ErrorLevel)
		return logrusLog
	}

	envLoader := env.NewLoader()
	if err := envLoader.Load(); err != nil {
		logrusLog.Fatal("Error loading env: ", err.Error())
	}

	if envLoader.IsProduction() {
		logrusLog.SetLevel(log.ErrorLevel)
	} else {
		logrusLog.SetLevel(log.InfoLevel)
	}

	return logrusLog
}
