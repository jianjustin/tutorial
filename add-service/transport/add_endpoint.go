package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"jianjustin/add-grpc-service/middleware/otel"
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
		ctx, span := otel.Tracer.Start(ctx, "add")
		defer span.End()

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
