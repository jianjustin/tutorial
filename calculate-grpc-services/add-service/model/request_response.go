package model

import (
	"context"
	pb2 "go.guide/add-grpc-service/pb"
)

type AddRequest struct {
	A int64
}

type AddResponse struct {
	Ctx context.Context
	V   int64
}

func EncodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(AddRequest)
	return &pb2.AddRequest{A: r.A}, nil
}

func DecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb2.AddRequest)
	return AddRequest{A: r.A}, nil
}

func EncodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*AddResponse)
	return &pb2.AddResponse{V: r.V}, nil
}

func DecodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb2.AddResponse)
	return &AddResponse{V: r.V, Ctx: ctx}, nil
}
