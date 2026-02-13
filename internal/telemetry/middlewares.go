package telemetry

import (
	"net/http"
	"time"

	"github.com/waduhek/website/internal/telemetry/internal"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"go.opentelemetry.io/otel/semconv/v1.39.0/httpconv"
)

type TelemetryCollector struct {
	requestDurationMeter httpconv.ServerRequestDuration
}

// NewTelemetryCollector creates a new instance of the telemetry collector.
func NewTelemetryCollector(
	meter metric.Meter,
) (*TelemetryCollector, error) {
	requestDurationMeter, err := httpconv.NewServerRequestDuration(
		meter,
		metric.WithExplicitBucketBoundaries(
			0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1, 2.5, 5,
			7.5, 10,
		),
	)
	if err != nil {
		return nil, err
	}

	// Go runtime metrics. These are all observable meters and so are not saved
	// to the the struct's fields.

	err = internal.RegisterMemoryUsedMetric(meter)
	if err != nil {
		return nil, err
	}

	return &TelemetryCollector{requestDurationMeter}, nil
}

// CollectDefaultMetricsMiddleware is a middleware function that collects all
// the default metrics for the HTTP server.
func (c *TelemetryCollector) CollectDefaultMetricsMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		method := httpconv.RequestMethodAttr(r.Method)
		scheme := r.URL.Scheme
		route := r.URL.Path
		routeAttribute := semconv.HTTPRoute(route)

		start := time.Now()

		// Perform the request.
		next.ServeHTTP(w, r)

		requestDuration := time.Since(start).Seconds()

		c.requestDurationMeter.Record(
			ctx, requestDuration, method, scheme,
			routeAttribute,
		)
	})
}
