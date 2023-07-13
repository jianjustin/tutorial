package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "go.guide/sub-grpc-service/pb"
	"go.guide/sub-grpc-service/service"
)

func Endpoints(svc service.SubService) EndpointsSet {
	return EndpointsSet{
		SubEndpoint:         makeSubEndpoint(svc),
		SubAfterAddEndpoint: makeSubAfterAddEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedSubServiceServer
	SubEndpoint         endpoint.Endpoint
	SubAfterAddEndpoint endpoint.Endpoint
}

func (s *EndpointsSet) Sub(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	resp, err := s.SubEndpoint(ctx, req)
	return resp.(*pb.SubResponse), err
}

func (s *EndpointsSet) SubAfterAdd(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	resp, err := s.SubAfterAddEndpoint(ctx, req)
	return resp.(*pb.SubResponse), err
}

func makeSubEndpoint(svc service.SubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*SubRequest)
		_, v, err := svc.Sub(ctx, req.A)
		return &SubResponse{
			V: v,
		}, err
	}
}

func makeSubAfterAddEndpoint(svc service.SubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*SubRequest)
		_, v, err := svc.SubAfterAdd(ctx, req.A)
		return &SubResponse{
			V: v,
		}, err
	}
}
