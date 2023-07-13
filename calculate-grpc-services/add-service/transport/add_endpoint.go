package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
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

func (s *EndpointsSet) Add(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.AddEndpoint(ctx, req)
	return resp.(*pb2.AddResponse), err
}

func (s *EndpointsSet) AddAfterMul(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.AddAfterMulEndpoint(ctx, req)
	return resp.(*pb2.AddResponse), err
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
