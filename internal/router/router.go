package router

import (
	"bbone/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"net/http"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Tracer(HomeHandler))

	return r
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	sp := request.Context().Value(middleware.SpanKey)
	span, ok := sp.(jaeger.Span)
	if !ok {
		log.Error().Msg("can not conver jaeger span from request context")
	}
	tracer := opentracing.GlobalTracer()
	child := tracer.StartSpan("inner_span_"+request.RequestURI, opentracing.ChildOf(span.SpanContext()))
	child.SetTag("inner", "tag")

	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "Response!!!")
}
