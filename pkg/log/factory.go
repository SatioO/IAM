package log

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Factory struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// Bg creates a context-unaware logger.
func (b Factory) Bg() Logger {
	return logger(b)
}

func (b Factory) For(ctx context.Context) Logger {
	if span := trace.SpanFromContext(ctx); span != nil {
		return b.Bg().With(
			zap.Field{
				Key:    "trace-id",
				Type:   zapcore.StringType,
				String: span.SpanContext().TraceID().String(),
			},
		).With(
			zap.Field{
				Key:    "span-id",
				Type:   zapcore.StringType,
				String: span.SpanContext().SpanID().String(),
			},
		)
	}
	return b.Bg()
}
