package model

import (
	"context"
	"go.guide/mul-grpc-service/pb"
)

type MulRequest struct {
	A int64
}

type MulResponse struct {
	Ctx context.Context
	V   int64
}

func EncodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(MulRequest)
	return &pb.MulRequest{A: r.A}, nil
}

func DecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.MulRequest)
	return MulRequest{A: r.A}, nil
}

func EncodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*MulResponse)
	return &pb.MulResponse{V: r.V}, nil
}

func DecodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.MulResponse)
	return &MulResponse{V: r.V, Ctx: ctx}, nil
}
