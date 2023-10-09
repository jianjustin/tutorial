package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	pb "jianjustin/sub-grpc-service/pb"
	"jianjustin/sub-grpc-service/proxying"
	"jianjustin/sub-grpc-service/service"
	"os"
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
		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/add", logger)
		addService := svc1(nil)
		ctx, a, err := addService.Add(ctx, req.A)
		if err != nil {
			return nil, err
		}
		req.A = a
		_, v, err := svc.Sub(ctx, req.A)
		return &SubResponse{
			V: v,
		}, err
	}
}
