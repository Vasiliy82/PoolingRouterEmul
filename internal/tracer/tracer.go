package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

const serviceName = "delivery-api"

var global trace.Tracer

func init() {
	SetTracer(New())
}

func New() trace.Tracer {
	return otel.Tracer(serviceName)
}

func SetTracer(tr trace.Tracer) {
	global = tr
}

func Tracer() trace.Tracer {
	return global
}

func InitHTTPProvider(url string, appName string, appID int64) (func(context.Context) error, error) {
	ctx := context.Background()
	client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint(url), otlptracehttp.WithInsecure())

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(appName),
			attribute.Int64("ID", appID),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider.Shutdown, nil
}
