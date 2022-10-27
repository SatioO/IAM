package log

import (
	"context"
	"fmt"
	"reflect"
	"unsafe"

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

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}

func (b Factory) For(ctx context.Context) Logger {
	if span := trace.SpanFromContext(ctx); span != nil {
		return b.Bg().With(
			zap.Field{
				Key:    "x-trace-id",
				Type:   zapcore.StringType,
				String: span.SpanContext().TraceID().String(),
			},
		).With(
			zap.Field{
				Key:    "x-span-id",
				Type:   zapcore.StringType,
				String: span.SpanContext().SpanID().String(),
			},
		)
	}
	return b.Bg()
}
