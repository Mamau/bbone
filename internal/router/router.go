package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	return r
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("do tracer!!!")
	defer span.Finish()

	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "Response!!!")
}