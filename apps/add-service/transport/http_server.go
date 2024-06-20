package transport

import (
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/add").Handler(
		http.NewServer(
			endpoints.AddEndpoint,
			_Decode_Add_Request,
			_Encode_Add_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "add", logger)))...))
	mux.Methods("POST").Path("/addAfterMul").Handler(
		http.NewServer(
			endpoints.AddAfterMulEndpoint,
			_Decode_Add_Request,
			_Encode_Add_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "addAfterMul", logger)))...))
	return mux
}
