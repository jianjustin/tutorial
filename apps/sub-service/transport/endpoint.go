package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/sub-grpc-service/pb"
	"go.guide/sub-grpc-service/proxying"
	"go.guide/sub-grpc-service/service"
	"os"
)

func Endpoints(svc service.SubService) EndpointsSet {
	return EndpointsSet{
		SubEndpoint: makeSubEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedSubServiceServer
	SubEndpoint endpoint.Endpoint
}

func makeSubEndpoint(svc service.SubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SubRequest)

		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/core", logger)
		coreService := svc1(nil)
		ctx, a, err := coreService.Random(ctx, req.A)
		if err != nil {
			return nil, err
		}
		req.A = a
		_, v, err := svc.Sub(ctx, req.A)
		return &pb.SubResponse{
			V: v,
		}, err
	}
}
