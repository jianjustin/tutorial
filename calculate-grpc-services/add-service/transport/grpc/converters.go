package transportgrpc

import (
	"context"
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/transport"
)

func _Encode_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*transport.AddRequest)
	return &pb2.AddRequest{A: r.A}, nil
}

func _Decode_Add_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb2.AddRequest)
	return &transport.AddRequest{A: r.A}, nil
}

func _Encode_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*transport.AddResponse)
	return &pb2.AddResponse{V: r.V}, nil
}

func _Decode_Add_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb2.AddResponse)
	//add Ctx response
	return &transport.AddResponse{V: r.V}, nil
}
