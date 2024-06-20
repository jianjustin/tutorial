package transport

import (
	"context"
	pb2 "go.guide/add-grpc-service/pb"
)

func _Encode_Grpc_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*AddRequest)
	return &pb2.AddRequest{A: r.A}, nil
}

func _Decode_Grpc_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb2.AddRequest)
	return &AddRequest{A: r.A}, nil
}

func _Encode_Grpc_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*AddResponse)
	return &pb2.AddResponse{V: r.V}, nil
}

func _Decode_Grpc_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb2.AddResponse)
	//add Ctx response
	return &AddResponse{V: r.V}, nil
}
