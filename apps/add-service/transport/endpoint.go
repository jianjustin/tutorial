package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
)

func Endpoints(svc service.AddService) EndpointsSet {
	return EndpointsSet{
		AddEndpoint:         makeAddEndpoint(svc),
		AddAfterMulEndpoint: makeAddAfterMulEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedAddServiceServer
	AddEndpoint         endpoint.Endpoint
	AddAfterMulEndpoint endpoint.Endpoint
}

func makeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddRequest)
		_, v, err := svc.Add(ctx, req.A)
		return &pb.AddResponse{
			V: v,
		}, err
	}
}

func makeAddAfterMulEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddRequest)
		_, v, err := svc.AddAfterMul(ctx, req.A)
		return &pb.AddResponse{
			V: v,
		}, err
	}
}
