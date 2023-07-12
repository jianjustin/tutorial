package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/add-grpc-service/model"
	"go.guide/add-grpc-service/pb"
)

type addServiceServer struct {
	pb.UnimplementedAddServiceServer
	add         grpc.Handler
	addAfterMul grpc.Handler
}

func (a addServiceServer) Add(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	_, resp, err := a.add.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddResponse), nil
}

func (a addServiceServer) AddAfterMul(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {
	_, resp, err := a.add.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddResponse), nil
}

func NewGRPCServer(endpoints *EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...grpc.ServerOption) pb.AddServiceServer {
	return &addServiceServer{
		add: grpc.NewServer(
			endpoints.add,
			model.DecodeRequest,
			model.EncodeResponse,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "Concat", logger)))...,
		),
		addAfterMul: grpc.NewServer(
			endpoints.addAfterMul,
			model.DecodeRequest,
			model.EncodeResponse,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "Sum", logger)))...,
		),
	}
}
