package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/div-grpc-service/pb"
	http1 "net/http"
)

type divServiceServer struct {
	pb.UnimplementedDivServiceServer
	div grpc.Handler
}

func (a divServiceServer) Div(ctx context.Context, request *pb.DivRequest) (*pb.DivResponse, error) {
	_, resp, err := a.div.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DivResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.DivServiceServer {
	return &divServiceServer{
		div: grpc.NewServer(
			endpoints.DivEndpoint,
			_Decode_Grpc_Div_Request,
			_Encode_Grpc_Div_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "div", logger)))...,
		),
	}
}

func NewHTTPHandler(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/div").Handler(
		http.NewServer(
			endpoints.DivEndpoint,
			_Decode_Http_Div_Request,
			_Encode_Http_Div_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "div", logger)))...))
	return mux
}
