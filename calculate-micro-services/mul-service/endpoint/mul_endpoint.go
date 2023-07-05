package endpoint

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"go.guide/mul-service/model"
	"go.guide/mul-service/service"
	"net/http"
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
		v, err := svc.Mul(req.A)
		if err != nil {
			return model.MulResponse{V: v, Err: err.Error()}, nil
		}
		return model.MulResponse{V: v}, nil
	}
}

func DecodeMulRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.MulRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
