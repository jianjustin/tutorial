package transportgrpc

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/sub-grpc-service/pb"
	"go.guide/sub-grpc-service/transport"
)

type subServiceServer struct {
	pb.UnimplementedSubServiceServer
	sub         grpc.Handler
	subAfterAdd grpc.Handler
}

func (a subServiceServer) Sub(ctx context.Context, request *pb.SubRequest) (*pb.SubResponse, error) {
	_, resp, err := a.sub.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SubResponse), nil
}

func (a subServiceServer) SubAfterAdd(ctx context.Context, request *pb.SubRequest) (*pb.SubResponse, error) {
	_, resp, err := a.subAfterAdd.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SubResponse), nil
}

func NewGRPCServer(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.SubServiceServer {
	return &subServiceServer{
		sub: grpc.NewServer(
			endpoints.SubEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "sub", logger)))...,
		),
		subAfterAdd: grpc.NewServer(
			endpoints.SubAfterAddEndpoint,
			_Decode_Sub_Request,
			_Encode_Sub_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "subAfterMul", logger)))...,
		),
	}
}
