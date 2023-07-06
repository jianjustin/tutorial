package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.guide/mul-grpc-service/pb"
	"go.guide/mul-grpc-service/service"
)

func MakeMulGRPCServer(svc service.MulService) pb.MulServiceServer {
	return &mulGrpcServer{
		mul:         makeMulEndpoint(svc),
		mulAfterAdd: makeMulAfterAddEndpoint(svc),
	}
}

type mulGrpcServer struct {
	pb.UnimplementedMulServiceServer
	mul         endpoint.Endpoint
	mulAfterAdd endpoint.Endpoint
}

func (s *mulGrpcServer) Mul(ctx context.Context, req *pb.MulRequest) (*pb.MulResponse, error) {
	resp, err := s.mul(ctx, req)
	return resp.(*pb.MulResponse), err
}

func (s *mulGrpcServer) MulAfterAdd(ctx context.Context, req *pb.MulRequest) (*pb.MulResponse, error) {
	resp, err := s.mulAfterAdd(ctx, req)
	return resp.(*pb.MulResponse), err
}

func makeMulEndpoint(svc service.MulService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.MulRequest)
		_, v, err := svc.Mul(ctx, req.A)
		return &pb.MulResponse{
			V: v,
			//Ctx: newCtx,
		}, err
	}
}

func makeMulAfterAddEndpoint(svc service.MulService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.MulRequest)
		_, v, err := svc.MulAfterAdd(ctx, req.A)
		return &pb.MulResponse{
			V: v,
			//Ctx: newCtx,
		}, err
	}
}
