package telemetry

import (
	"context"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"go.opentelemetry.io/otel/trace"
)

const PackageName string = "github.com/waduhek/website"

// ExtractContext takes an incoming request's context and the headers and
// extracts trace context into a new context object.
func ExtractContext(
	reqCtx context.Context,
	header http.Header,
) context.Context {
	propagator := otel.GetTextMapPropagator()

	return propagator.Extract(
		reqCtx,
		propagation.HeaderCarrier(header),
	)
}

// InjectContext takes a trace context and injects it into the provided header
// object.
func InjectContext(ctx context.Context, header http.Header) {
	propagator := otel.GetTextMapPropagator()

	carrier := make(propagation.HeaderCarrier)
	propagator.Inject(ctx, &carrier)

	for key, values := range carrier {
		for _, value := range values {
			header.Add(key, value)
		}
	}
}

// GetTracer gets the tracer for the application.
func GetTracer() trace.Tracer {
	return otel.GetTracerProvider().Tracer(PackageName)
}

// NewSpan creates a new span from the provided context. Returns a new context
// for the span along with the span object.
func NewSpan(
	ctx context.Context,
	spanName string,
) (context.Context, trace.Span) {
	hostName := os.Getenv("HOSTNAME")
	tracer := GetTracer()

	return tracer.Start(
		ctx,
		spanName,
		trace.WithAttributes(semconv.ContainerID(hostName)),
	)
}
