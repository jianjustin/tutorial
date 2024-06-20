package transport

import (
	"context"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	"go.guide/div-grpc-service/pb"
)

type divServiceServer struct {
	pb.UnimplementedDivServiceServer
	div         grpc.Handler
	divAfterAdd grpc.Handler
}

func (a divServiceServer) Div(ctx context.Context, request *pb.DivRequest) (*pb.DivResponse, error) {
	_, resp, err := a.div.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DivResponse), nil
}

func (a divServiceServer) DivAfterAdd(ctx context.Context, request *pb.DivRequest) (*pb.DivResponse, error) {
	_, resp, err := a.divAfterAdd.ServeGRPC(ctx, request)
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
		divAfterAdd: grpc.NewServer(
			endpoints.DivAfterAddEndpoint,
			_Decode_Grpc_Div_Request,
			_Encode_Grpc_Div_Response,
			append(opts, grpc.ServerBefore(
				opentracing.GRPCToContext(tracer, "div_after_add", logger)))...,
		),
	}
}
