package model

import (
	"context"
	"go.guide/add-grpc-service/pb"
)

type AddRequest struct {
	A int64
}

type AddResponse struct {
	Ctx context.Context
	V   int64
}

func EncodeAddRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(AddRequest)
	return &pb.AddRequest{A: r.A}, nil
}

func DecodeAddRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.AddRequest)
	return AddRequest{A: r.A}, nil
}

func EncodeAddResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*AddResponse)
	return &pb.AddResponse{V: r.V}, nil
}

func DecodeAddResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.AddResponse)
	return &AddResponse{V: r.V, Ctx: ctx}, nil
}

func EmptyResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	return resp, nil
}
