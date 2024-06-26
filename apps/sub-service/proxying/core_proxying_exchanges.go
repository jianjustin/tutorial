package proxying

import (
	"context"
	"go.guide/sub-grpc-service/pb"
)

func _Encode_Grpc_Random_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(pb.RandomRequest)
	return &pb.RandomRequest{A: r.A}, nil
}

func _Decode_Grpc_Random_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.RandomResponse)
	return &pb.RandomResponse{V: r.V}, nil
}
