package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/mul-grpc-service/pb"
	"go.guide/mul-grpc-service/proxying"
	"go.guide/mul-grpc-service/service"
	"os"
)

func Endpoints(svc service.MulService) EndpointsSet {
	return EndpointsSet{
		MulEndpoint: makeMulEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedMulServiceServer
	MulEndpoint endpoint.Endpoint
}

func makeMulEndpoint(svc service.MulService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.MulRequest)

		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/core", logger)
		coreService := svc1(nil)
		ctx, a, err := coreService.Random(ctx, req.A)
		if err != nil {
			return nil, err
		}
		req.A = a
		_, v, err := svc.Mul(ctx, req.A)
		return &pb.MulResponse{
			V: v,
		}, err
	}
}
