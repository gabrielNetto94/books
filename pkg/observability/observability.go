package observability

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Span(ctx context.Context, name string) (context.Context, ISpan)
}

type tracer struct {
	tracer trace.Tracer
}

func NewTracer(service string) Tracer {
	return &tracer{
		tracer: otel.Tracer(service),
	}
}

func (t *tracer) Span(ctx context.Context, name string) (context.Context, ISpan) {
	ctx, s := t.tracer.Start(ctx, name)
	return ctx, &otelSpan{span: s}
}

type ISpan interface {
	End()
	AddEvent(name string)
	SetAttribute(key string, value any)
	RecordError(err error)
}

type otelSpan struct {
	span trace.Span
}

func (s *otelSpan) End() {
	s.span.End()
}

func (s *otelSpan) AddEvent(name string) {
	s.span.AddEvent(name)
}

func (s *otelSpan) SetAttribute(key string, value any) {
	switch v := value.(type) {
	case string:
		s.span.SetAttributes(attribute.String(key, v))
	case int:
		s.span.SetAttributes(attribute.Int(key, v))
	case int64:
		s.span.SetAttributes(attribute.Int64(key, v))
	case float64:
		s.span.SetAttributes(attribute.Float64(key, v))
	case bool:
		s.span.SetAttributes(attribute.Bool(key, v))
	case []string:
		s.span.SetAttributes(attribute.StringSlice(key, v))
	case []int:
		s.span.SetAttributes(attribute.IntSlice(key, v))
	case []int64:
		s.span.SetAttributes(attribute.Int64Slice(key, v))
	case []float64:
		s.span.SetAttributes(attribute.Float64Slice(key, v))
	case []bool:
		s.span.SetAttributes(attribute.BoolSlice(key, v))
	default:
		if jsonData, err := json.Marshal(value); err == nil {
			s.span.SetAttributes(attribute.String(key, string(jsonData)))
		}
	}
}

func (s *otelSpan) RecordError(err error) {
	s.span.RecordError(err)
}
