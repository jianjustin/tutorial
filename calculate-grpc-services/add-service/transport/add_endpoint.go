package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
)

type addService struct{}

func (addService) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a + int64(3), nil
}

func (addService) AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a + int64(4), nil
}

func NewAddService() service.AddService {
	return addService{}
}

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
	str, _ := json.Marshal(resp)
	res := &pb2.AddResponse{}
	json.Unmarshal(str, res)
	return res, err
}

func (s *addGrpcServer) AddAfterMul(ctx context.Context, req *pb2.AddRequest) (*pb2.AddResponse, error) {
	resp, err := s.addAfterMul(ctx, req)
	str, _ := json.Marshal(resp)
	res := &pb2.AddResponse{}
	json.Unmarshal(str, res)
	return res, err
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
