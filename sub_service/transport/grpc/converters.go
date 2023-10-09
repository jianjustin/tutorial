package transportgrpc

import (
	"context"
	"jianjustin/sub-grpc-service/pb"
	"jianjustin/sub-grpc-service/transport"
)

func _Encode_Sub_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*transport.SubRequest)
	return &pb.SubRequest{A: r.A}, nil
}

func _Decode_Sub_Request(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.SubRequest)
	return &transport.SubRequest{A: r.A}, nil
}

func _Encode_Sub_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*transport.SubResponse)
	return &pb.SubResponse{V: r.V}, nil
}

func _Decode_Sub_Response(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.SubResponse)
	//add Ctx response
	return &transport.SubResponse{V: r.V}, nil
}
