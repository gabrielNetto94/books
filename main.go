package main

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"
	"books/internal/adapters/http/routes"
	loglevel "books/internal/infra/log"
	"books/internal/infra/log/logrus"
	"context"
	"fmt"
	"os"
	"os/signal"

	bookrepository "books/internal/core/repositories/book"
	cache "books/internal/core/repositories/cache/redis"
	"books/internal/core/repositories/db"
	userrepository "books/internal/core/repositories/user"
	"books/internal/core/services"
	"books/pkg/env"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const REST_API_PORT = ":3000"
const GRPC_SERVER_PORT = ":3001"

// Initialize a gRPC connection to be used by both the tracer and meter
// providers.
func initConn() (*grpc.ClientConn, error) {
	// It connects the OpenTelemetry Collector through local gRPC connection.
	// You may replace `localhost:4317` with your endpoint.
	conn, err := grpc.NewClient("localhost:4317",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

// Initializes an OTLP exporter, and configures the corresponding trace provider.
func initTracerProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

func main() {

	db := db.ConnectDatabase("postgres://postgres:password@db:5432")
	cache := cache.ConnectCache("redis://cache:6379")

	log := setupLogger()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	conn, err := initConn()
	if err != nil {
		log.Fatal("Error creating gRPC connection: ", err.Error())
	}

	var serviceName = semconv.ServiceNameKey.String("test-service")
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := initTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatal("Error shutting down tracer provider: ", err.Error())
		}
	}()

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
