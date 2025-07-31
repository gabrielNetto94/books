package observability

import (
	"context"
)

type Tracer interface {
	Span(ctx context.Context, name string) (context.Context, Span)
}

type Span interface {
	End()
	AddEvent(name string)
	SetAttribute(key string, value any)
	RecordError(err error)
}
