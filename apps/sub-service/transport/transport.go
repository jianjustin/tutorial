package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/sub-grpc-service/pb"
	http1 "net/http"
)

type subServiceServer struct {
	pb.UnimplementedSubServiceServer
	sub grpc.Handler
}

func (a subServiceServer) Add(ctx context.Context, request *pb.SubRequest) (*pb.SubResponse, error) {
	_, resp, err := a.sub.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SubResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.SubServiceServer {
	return &subServiceServer{
		sub: grpc.NewServer(
			endpoints.SubEndpoint,
			_Decode_Grpc_Sub_Request,
			_Encode_Grpc_Sub_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "sub", logger)))...,
		),
	}
}

func NewHTTPHandler(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/sub").Handler(
		http.NewServer(
			endpoints.SubEndpoint,
			_Decode_Http_Sub_Request,
			_Encode_Http_Sub_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "sub", logger)))...))
	return mux
}
