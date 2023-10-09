package proxying

import (
	"context"
	"jianjustin/add-grpc-service/pb"
)

type AddRequest struct {
	A int64
}

type AddResponse struct {
	V int64
}

func _Encode_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(AddRequest)
	return &pb.AddRequest{A: r.A}, nil
}

func _Decode_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.AddRequest)
	return AddRequest{A: r.A}, nil
}

func _Encode_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*AddResponse)
	return &pb.AddResponse{V: r.V}, nil
}

func _Decode_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.AddResponse)
	return &AddResponse{V: r.V}, nil
}
