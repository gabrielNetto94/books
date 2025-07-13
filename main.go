package main

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"
	"books/internal/adapters/http/routes"
	loglevel "books/internal/infra/log"
	"books/internal/infra/log/logrus"
	"context"
	"os"
	"os/signal"

	bookrepository "books/internal/core/repositories/book"
	cache "books/internal/core/repositories/cache/redis"
	"books/internal/core/repositories/db/gorm"
	userrepository "books/internal/core/repositories/user"
	"books/internal/core/services"
	"books/pkg/env"
	"books/pkg/observability/opentelemetry"

	gormm "gorm.io/gorm"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

var serviceName = semconv.ServiceNameKey.String("asdf-test")

func main() {

	db := gorm.NewDatabaseRepository("postgres://postgres:password@db:5432")
	cache := cache.NewCacheInstance("redis://cache:6379")

	log := setupLogger()

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

	bookRepo := bookrepository.NewBookRepository(&gormm.DB{}, cache) //@todo fix dependency injection
	service := services.NewBookService(bookRepo, log, opentelemetry)
	bookHandler := bookhandler.NewBookHandlers(service, log)

	userRepo := userrepository.NewUserRepository(db, cache)
	userService := services.NewUserService(userRepo, log)
	userHandler := userhandler.NewUserHandlers(userService, log)

	router := routes.InitRouter(bookHandler, userHandler)

	// tracer := otel.Tracer(serviceName.Value.AsString())
	// _, iSpan := tracer.Start(ctx, fmt.Sprintf("Sample-%d", 123))
	// iSpan.End()

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
