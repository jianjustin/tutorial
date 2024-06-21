package proxying

import (
	"context"
	"go.guide/div-grpc-service/pb"
)

func _Encode_Grpc_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(pb.AddRequest)
	return &pb.AddRequest{A: r.A}, nil
}

func _Decode_Grpc_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.AddRequest)
	return pb.AddRequest{A: r.A}, nil
}

func _Encode_Grpc_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.AddResponse)
	return &pb.AddResponse{V: r.V}, nil
}

func _Decode_Grpc_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.AddResponse)
	return &pb.AddResponse{V: r.V}, nil
}
