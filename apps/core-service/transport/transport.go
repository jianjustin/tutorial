package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/core-service/pb"
)

type coreServiceServer struct {
	pb.UnimplementedCoreServiceServer
	random grpc.Handler
}

func (a coreServiceServer) Random(ctx context.Context, request *pb.RandomRequest) (*pb.RandomResponse, error) {
	_, resp, err := a.random.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RandomResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.CoreServiceServer {
	return &coreServiceServer{
		random: grpc.NewServer(
			endpoints.RandomEndpoint,
			_Decode_Grpc_Random_Request,
			_Encode_Grpc_Random_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "random", logger)))...,
		),
	}
}
