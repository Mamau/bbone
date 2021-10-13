package middleware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type Span string

const SpanKey Span = "span"

// Tracer позволяет делать трассировку запросов
func Tracer(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		url := request.RequestURI
		tracer := opentracing.GlobalTracer()
		span := tracer.StartSpan("url_" + url)
		defer span.Finish()

		ctx := context.WithValue(request.Context(), SpanKey, &span)
		request = request.WithContext(ctx)

		h(writer, request)
	}
}
