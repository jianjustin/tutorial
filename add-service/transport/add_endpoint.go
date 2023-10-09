package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb2 "jianjustin/add-grpc-service/pb"
	"jianjustin/add-grpc-service/service"
)

func Endpoints(svc service.AddService) EndpointsSet {
	return EndpointsSet{
		AddEndpoint:         makeAddEndpoint(svc),
		AddAfterMulEndpoint: makeAddAfterMulEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb2.UnimplementedAddServiceServer
	AddEndpoint         endpoint.Endpoint
	AddAfterMulEndpoint endpoint.Endpoint
}

func makeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*AddRequest)
		_, v, err := svc.Add(ctx, req.A)
		return &AddResponse{
			V: v,
		}, err
	}
}

func makeAddAfterMulEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*AddRequest)
		_, v, err := svc.AddAfterMul(ctx, req.A)
		return &AddResponse{
			V: v,
		}, err
	}
}
