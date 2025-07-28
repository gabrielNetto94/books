# API Books

A simple book management API built with Go, fully instrumented for observability.

## 🏛️ Architecture

This project is structured following the principles of **Hexagonal Architecture** (also known as Ports and Adapters). This approach isolates the core business logic from external concerns like databases, web frameworks, and other services.

-   **Core (`internal/core`)**: Contains the application's domain logic, services (use cases), and repository interfaces (ports).
-   **Adapters (`internal/adapters`)**: Implements the entry points to the core logic, such as the Gin HTTP handlers.
-   **Infrastructure (`internal/infra`)**: Provides concrete implementations for the interfaces defined in the core, like the GORM database repository and the Logrus logger.

## 🛠️ Technologies Used

-   **Framework**: [Gin](https://gin-gonic.com/) for the HTTP server.
-   **Database**: [PostgreSQL](https://www.postgresql.org/) with [GORM](https://gorm.io/) as the ORM.
-   **Logging**: [Logrus](https://github.com/sirupsen/logrus).
-   **Observability Stack**:
    -   **Tracing**: [OpenTelemetry](https://opentelemetry.io/) sending data to [Jaeger](https://www.jaegertracing.io/).
    -   **Metrics**: [Prometheus](https://prometheus.io/) for collecting and storing metrics.
    -   **Visualization**: [Grafana](https://grafana.com/) for dashboards.

## 🚀 Getting Started

### Prerequisites

-   [Docker](https://docs.docker.com/get-docker/)
-   [Docker Compose](https://docs.docker.com/compose/install/)

### Running the Project

To start the application and all its dependencies, run the following command from the root of the project:

```sh
docker compose up
```

-   **Application** will be available at `http://localhost:3005`.
-   **Healthcheck** will be available at `http://localhost:3005/ping`.
-   **Jaeger** will be available at `http://localhost:16686`.
-   **Prometheus** will be available at `http://localhost:9090`.
-   **Grafana** will be available at `http://localhost:3000`. The dashboard is automatically provisioned.

## 📋 To Do
- [ ] Fix the update book feature
- [ ] Fix the list all books feature
- [ ] Send metrics to OpenTelemetry instead of Prometheus
- [ ] Add Swagger/OpenAPI
- [ ] Finish gRPC implementation (currently unavailable)
- [ ] Remove DB and cache from the core and move them to adapters to follow Hexagonal Architecture
- [ ] Add rate limit
