package telemetry

import (
	"context"
	"errors"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

// SetupOTelSDK sets up the OTel SDK and returns a shutdown function to call
// when the SDK is to be shutdown.
func SetupOTelSDK(ctx context.Context) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error
	var setupErr error

	shutdown := func(shutdownCtx context.Context) error {
		var shutdownErr error

		for _, fn := range shutdownFuncs {
			shutdownErr = errors.Join(shutdownErr, fn(shutdownCtx))
		}
		shutdownFuncs = nil

		return shutdownErr
	}

	handleErr := func(e error) {
		setupErr = errors.Join(e, shutdown(ctx))
	}

	propagator := newPropagator()
	otel.SetTextMapPropagator(propagator)

	tracerProvider, err := newTracerProvider(ctx)
	if err != nil {
		handleErr(err)
		return shutdown, setupErr
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider(ctx)
	if err != nil {
		handleErr(err)
		return shutdown, setupErr
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	loggerProvider, err := newLoggerProvider(ctx)
	if err != nil {
		handleErr(err)
		return shutdown, setupErr
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return shutdown, setupErr
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	otlpEndpoint, err := getOTLPEndpoint()
	if err != nil {
		return nil, err
	}

	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(otlpEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := newResource()
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
	)

	return tracerProvider, err
}

func newMeterProvider(ctx context.Context) (*sdkmetric.MeterProvider, error) {
	otlpEndpoint, err := getOTLPEndpoint()
	if err != nil {
		return nil, err
	}

	metricExporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(otlpEndpoint),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := newResource()
	if err != nil {
		return nil, err
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(5*time.Second),
			),
		),
	)

	return meterProvider, nil
}

func newLoggerProvider(ctx context.Context) (*sdklog.LoggerProvider, error) {
	otlpEndpoint, err := getOTLPEndpoint()
	if err != nil {
		return nil, err
	}

	logExporter, err := otlploggrpc.New(
		ctx,
		otlploggrpc.WithEndpoint(otlpEndpoint),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := newResource()
	if err != nil {
		return nil, err
	}

	processor := sdklog.NewBatchProcessor(logExporter)

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(processor),
	)

	return loggerProvider, nil
}

func getOTLPEndpoint() (string, error) {
	endpoint := os.Getenv("OTLP_ENDPOINT")
	if endpoint == "" {
		err := errors.New(
			"Value of environment variable OTLP_ENDPOINT is required",
		)
		return "", err
	}

	return endpoint, nil
}

func newResource() (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("website-backend"),
		),
	)
}
