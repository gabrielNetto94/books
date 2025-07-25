package routes

import (
	bookhandler "books/internal/adapters/http/handlers/books"
	userhandler "books/internal/adapters/http/handlers/user"
	"books/pkg/observability/metrics"
	"books/pkg/observability/metrics/prometheus"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// InitRouter initializes the Gin router with all routes.
func InitRouter(bookHandler bookhandler.BookHTTPHandler, userHandler userhandler.UserHTTPHandler) *gin.Engine {

	r := gin.Default()
	r.Use(otelgin.Middleware("asdf-test"))

	collector := prometheus.NewPrometheusCollector()
	appMetrics, err := metrics.NewAppMetrics(collector)
	if err != nil {
		log.Fatal("Failed to create app metrics: ", err.Error())
	}
	r.Use(prometheus.MetricsMiddleware(appMetrics))

	r.GET("/metrics", gin.WrapH(collector.Handler()))
	// Health check
	r.GET("/ping", healthCheckHandler)

	// Book routes grouped by version
	v1 := r.Group("/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("", gin.WrapF(bookHandler.ListBooks))
			books.GET("/:id", gin.WrapF(bookHandler.GetBookById))
			books.POST("", gin.WrapF(bookHandler.CreateBook))
			books.PUT("/:id", gin.WrapF(bookHandler.UpdateBook))
		}
		users := v1.Group("/users")
		{
			users.POST("", gin.WrapF(userHandler.CreateUser))
		}
	}

	return r
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
