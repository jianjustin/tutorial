package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/mul-grpc-service/pb"
	http1 "net/http"
)

type mulServiceServer struct {
	pb.UnimplementedMulServiceServer
	mul grpc.Handler
}

func (a mulServiceServer) Mul(ctx context.Context, request *pb.MulRequest) (*pb.MulResponse, error) {
	_, resp, err := a.mul.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MulResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.MulServiceServer {
	return &mulServiceServer{
		mul: grpc.NewServer(
			endpoints.MulEndpoint,
			_Decode_Grpc_Mul_Request,
			_Encode_Grpc_Mul_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "mul", logger)))...,
		),
	}
}

func NewHTTPHandler(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/mul").Handler(
		http.NewServer(
			endpoints.MulEndpoint,
			_Decode_Http_Mul_Request,
			_Encode_Http_Mul_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "mul", logger)))...))
	return mux
}
