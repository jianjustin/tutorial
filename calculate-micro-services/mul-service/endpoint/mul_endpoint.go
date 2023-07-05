package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"go.guide/mul-service/middleware"
	"go.guide/mul-service/model"
	"go.guide/mul-service/service"
	"os"
)

func MakeMulEndpoint(svc service.MulService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.MulRequest)
		v, err := svc.Mul(req.A)
		if err != nil {
			return model.MulResponse{V: v, Err: err.Error()}, nil
		}
		return model.MulResponse{V: v}, nil
	}
}

func MakeMulAfterAddEndpoint(svc service.MulService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.MulRequest)

		logger := log.NewLogfmtLogger(os.Stdout)
		svc1 := middleware.ProxyingMiddleware(context.Background(), "/services/add", logger)
		addService := svc1(nil)
		a, err := addService.Add(req.A)
		if err != nil {
			return model.MulResponse{V: 0, Err: err.Error()}, nil
		}

		v, err := svc.Mul(a)
		if err != nil {
			return model.MulResponse{V: 0, Err: err.Error()}, nil
		}
		return model.MulResponse{V: v}, nil
	}
}
