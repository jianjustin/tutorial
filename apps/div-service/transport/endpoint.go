package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.guide/div-grpc-service/pb"
	"go.guide/div-grpc-service/service"
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
		_, v, err := svc.Div(ctx, req.A)
		return &pb.DivResponse{
			V: v,
		}, err
	}
}
