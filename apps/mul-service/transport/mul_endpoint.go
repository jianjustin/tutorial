package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/mul-grpc-service/middleware"
	"go.guide/mul-grpc-service/pb"
	"go.guide/mul-grpc-service/service"
	"os"
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
	logger := log.NewLogfmtLogger(os.Stdout)
	svc1 := middleware.ProxyingMiddleware(context.Background(), "/services/add", logger)
	addService := svc1(nil)
	ctx, a, err := addService.Add(ctx, req.A)
	if err != nil {
		return nil, err
	}
	req.A = a

	resp, err := s.mul(ctx, req)
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
