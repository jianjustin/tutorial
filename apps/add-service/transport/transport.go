package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/add-grpc-service/pb"
	http1 "net/http"
)

type addServiceServer struct {
	pb.UnimplementedAddServiceServer
	add grpc.Handler
}

func (a addServiceServer) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	_, resp, err := a.add.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.AddServiceServer {
	return &addServiceServer{
		add: grpc.NewServer(
			endpoints.AddEndpoint,
			_Decode_Grpc_Add_Request,
			_Encode_Grpc_Add_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "add", logger)))...,
		),
	}
}

func NewHTTPHandler(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/add").Handler(
		http.NewServer(
			endpoints.AddEndpoint,
			_Decode_Http_Add_Request,
			_Encode_Http_Add_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "add", logger)))...))
	return mux
}
