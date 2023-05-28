package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"go.guide/tutorial/go-kit/stringsvc/model"
	"go.guide/tutorial/go-kit/stringsvc/service"
	"net/http"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return model.UppercaseResponse{v, err.Error()}, nil
		}
		return model.UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CountRequest)
		v := svc.Count(req.S)
		return model.CountResponse{v}, nil
	}
}

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
