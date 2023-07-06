package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
)

func MakeAddGRPCServer(svc service.AddService) pb2.AddServiceServer {
	return &addGrpcServer{
		add:         makeAddEndpoint(svc),
		addAfterMul: makeAddAfterMulEndpoint(svc),
	}
}

type addGrpcServer struct {
	pb2.UnimplementedAddServiceServer
	add         endpoint.Endpoint
	addAfterMul endpoint.Endpoint
}

func (s *addGrpcServer) Add(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.add(ctx, req)
	return resp.(*pb2.AddResponse), err
}

func (s *addGrpcServer) AddAfterMul(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.addAfterMul(ctx, req)
	return resp.(*pb2.AddResponse), err
}

func makeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb2.AddRequest)
		_, v, err := svc.Add(ctx, req.A)
		return &pb2.AddResponse{
			V: v,
			//Ctx: newCtx,
		}, err
	}
}

func makeAddAfterMulEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb2.AddRequest)
		_, v, err := svc.AddAfterMul(ctx, req.A)
		return &pb2.AddResponse{
			V: v,
			//Ctx: newCtx,
		}, err
	}
}
