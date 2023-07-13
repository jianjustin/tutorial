package transporthttp

import (
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/sub-grpc-service/transport"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/sub").Handler(
		http.NewServer(
			endpoints.SubEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "sub", logger)))...))
	mux.Methods("POST").Path("/subAfterAdd").Handler(
		http.NewServer(
			endpoints.SubAfterAddEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "subAfterAdd", logger)))...))
	return mux
}
