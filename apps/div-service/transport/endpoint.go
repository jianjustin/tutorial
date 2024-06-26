package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/div-grpc-service/pb"
	"go.guide/div-grpc-service/proxying"
	"go.guide/div-grpc-service/service"
	"os"
)

func Endpoints(svc service.DivService) EndpointsSet {
	return EndpointsSet{
		DivEndpoint: makeDivEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedDivServiceServer
	DivEndpoint endpoint.Endpoint
}

func makeDivEndpoint(svc service.DivService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.DivRequest)
		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/core", logger)
		addService := svc1(nil)
		ctx, a, err := addService.Random(ctx, req.A)
		if err != nil {
			return nil, err
		}
		req.A = a
		_, v, err := svc.Div(ctx, req.A)
		return &pb.DivResponse{
			V: v,
		}, err
	}
}
