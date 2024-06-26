package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/proxying"
	"go.guide/add-grpc-service/service"
	"os"
)

func Endpoints(svc service.AddService) EndpointsSet {
	return EndpointsSet{
		AddEndpoint: makeAddEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedAddServiceServer
	AddEndpoint endpoint.Endpoint
}

func makeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddRequest)

		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/core", logger)
		coreService := svc1(nil)
		ctx, a, err := coreService.Random(ctx, req.A)
		if err != nil {
			return nil, err
		}
		req.A = a
		_, v, err := svc.Add(ctx, req.A)
		return &pb.AddResponse{
			V: v,
		}, err
	}
}
