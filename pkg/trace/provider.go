package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type ProviderConfig struct {
	JaegerEndpoint string
	ServiceName    string
	ServiceVersion string
	Environment    string
	Disabled       bool
}

type Provider struct {
	provider trace.TracerProvider
}

func NewProvider(config ProviderConfig) (Provider, error) {
	if config.Disabled {
		return Provider{provider: trace.NewNoopTracerProvider()}, nil
	}

	exporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerEndpoint)),
	)

	if err != nil {
		return Provider{}, err
	}

	traceProvider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSpanProcessor(tracesdk.NewBatchSpanProcessor(exporter)),
		tracesdk.WithResource(sdkresource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		)),
	)

	otel.SetTracerProvider(traceProvider)

	// Register the W3C trace context and baggage propagators so data is propagated across services/processes
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return Provider{provider: traceProvider}, nil
}

// Close shuts down the tracer provider only if it was not "no operations"
// version.
func (p Provider) Close(ctx context.Context) error {
	if prv, ok := p.provider.(*tracesdk.TracerProvider); ok {
		return prv.Shutdown(ctx)
	}

	return nil
}
