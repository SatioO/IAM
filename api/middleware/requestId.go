package middleware

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if span := trace.SpanFromContext(r.Context()); span != nil {
			fmt.Println(span.SpanContext().TraceID())
		}

		next.ServeHTTP(w, r)
	})
}
