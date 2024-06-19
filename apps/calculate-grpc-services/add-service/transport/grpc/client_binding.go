package transportgrpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"google.golang.org/grpc"
)

type AddClientBinding struct {
	E endpoint.Endpoint
}

func (s *AddClientBinding) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, &transport.AddRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*transport.AddResponse)
	return ctx, r.V, nil
}

func (s *AddClientBinding) AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, &transport.AddRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*transport.AddResponse)
	return ctx, r.V, nil
}

func NewAddClient(cc *grpc.ClientConn) service.AddService {
	return &AddClientBinding{
		E: grpctransport.NewClient(
			cc,
			"pb.AddService",
			"Add",
			_Encode_Add_Request,
			_Decode_Add_Response,
			&pb.AddResponse{},
		).Endpoint(),
	}
}

func NewAddAfterMulClient(cc *grpc.ClientConn) service.AddService {
	return &AddClientBinding{
		E: grpctransport.NewClient(
			cc,
			"pb.AddService",
			"AddAfterMul",
			_Encode_Add_Request,
			_Decode_Add_Response,
			&pb.AddResponse{},
		).Endpoint(),
	}
}
