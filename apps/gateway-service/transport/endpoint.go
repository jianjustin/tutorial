package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/gateway-service/pb"
	"go.guide/gateway-service/proxying"
	"os"
)

func Endpoints() EndpointsSet {
	return EndpointsSet{
		AddEndpoint: makeAddEndpoint(),
	}
}

type EndpointsSet struct {
	pb.UnimplementedAddServiceServer
	AddEndpoint endpoint.Endpoint
}

func makeAddEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddRequest)
		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := proxying.ProxyingMiddleware(context.Background(), "/services/add", logger)
		addService := svc1(nil)
		_, v, err := addService.Add(ctx, req.A)
		if err != nil {
			return nil, err
		}
		return &pb.AddResponse{
			V: v,
		}, err
	}
}
