package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.guide/add-grpc-service/model"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
)

func Endpoints(svc service.AddService) EndpointsSet {
	return EndpointsSet{
		add:         makeAddEndpoint(svc),
		addAfterMul: makeAddAfterMulEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb2.UnimplementedAddServiceServer
	add         endpoint.Endpoint
	addAfterMul endpoint.Endpoint
}

func (s *EndpointsSet) Add(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.add(ctx, req)
	return resp.(*pb2.AddResponse), err
}

func (s *EndpointsSet) AddAfterMul(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.addAfterMul(ctx, req)
	return resp.(*pb2.AddResponse), err
}

func makeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddRequest)
		_, v, err := svc.Add(ctx, req.A)
		return model.AddResponse{
			V:   v,
			Ctx: ctx,
		}, err
	}
}

func makeAddAfterMulEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddRequest)
		_, v, err := svc.AddAfterMul(ctx, req.A)
		return model.AddResponse{
			V:   v,
			Ctx: ctx,
		}, err
	}
}
