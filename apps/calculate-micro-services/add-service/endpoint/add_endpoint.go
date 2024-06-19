package endpoint

import (
	"context"
	"encoding/json"
	"go.guide/add-service/model"
	"go.guide/add-service/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeAddEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddRequest)
		v, err := svc.Add(req.A)
		if err != nil {
			return model.AddResponse{V: v, Err: err.Error()}, nil
		}
		return model.AddResponse{V: v}, nil
	}
}

func MakeAddAfterMulEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AddRequest)
		v, err := svc.Add(req.A)
		if err != nil {
			return model.AddResponse{V: v, Err: err.Error()}, nil
		}
		return model.AddResponse{V: v}, nil
	}
}

func DecodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
