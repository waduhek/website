package telemetry

import (
	"net/http"
	"time"

	"github.com/waduhek/website/internal/telemetry/internal"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/semconv/v1.37.0/httpconv"
)

type TelemetryCollector struct {
	activeRequestsMeter  httpconv.ServerActiveRequests
	requestDurationMeter httpconv.ServerRequestDuration
}

// NewTelemetryCollector creates a new instance of the telemetry collector.
func NewTelemetryCollector(
	meter metric.Meter,
) (*TelemetryCollector, error) {
	activeRequestsMeter, err := httpconv.NewServerActiveRequests(meter)
	if err != nil {
		return nil, err
	}

	requestDurationMeter, err := httpconv.NewServerRequestDuration(meter)
	if err != nil {
		return nil, err
	}

	// Go runtime metrics. These are all observable meters and so are not saved
	// to the the struct's fields.

	err = internal.RegisterMemoryUsedMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterMemoryLimitMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterMemoryAllocatedMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterMemoryAllocationsMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterGCGoalMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterGoRoutineCountMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterProcessorLimitMetric(meter)
	if err != nil {
		return nil, err
	}

	err = internal.RegisterGoGCConfigMetric(meter)
	if err != nil {
		return nil, err
	}

	return &TelemetryCollector{activeRequestsMeter, requestDurationMeter}, nil
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

		c.activeRequestsMeter.Add(ctx, 1, method, scheme, routeAttribute)

		start := time.Now()

		// Perform the request.
		next.ServeHTTP(w, r)

		requestDuration := time.Since(start).Seconds()

		c.activeRequestsMeter.Add(ctx, -1, method, scheme)
		c.requestDurationMeter.Record(
			ctx, requestDuration, method, scheme,
			routeAttribute,
		)
	})
}
