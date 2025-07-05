package observability

import "context"

type Observability interface {
	// Tracing
	StartSpan(ctx context.Context, name string) (context.Context, Span)

	// // Logging
	// Info(ctx context.Context, msg string, fields map[string]any)
	// Error(ctx context.Context, msg string, fields map[string]any)
	// Debug(ctx context.Context, msg string, fields map[string]any)
}

// Represents a simplified span abstraction
type Span interface {
	End()
	// SetAttribute(key string, value any)
	AddEvent(name string)
}
