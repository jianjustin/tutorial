package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.guide/core-service/pb"
	"go.guide/core-service/service"
)

func Endpoints(svc service.CoreService) EndpointsSet {
	return EndpointsSet{
		RandomEndpoint: makeRandomEndpoint(svc),
	}
}

type EndpointsSet struct {
	pb.UnimplementedCoreServiceServer
	RandomEndpoint endpoint.Endpoint
}

func makeRandomEndpoint(svc service.CoreService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.RandomRequest)
		_, v, err := svc.Random(ctx, req.A)
		return &pb.RandomResponse{
			V: v,
		}, err
	}
}
