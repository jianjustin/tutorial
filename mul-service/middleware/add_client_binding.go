package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"jianjustin/add-grpc-service/pb"
	"jianjustin/add-grpc-service/service"
	"jianjustin/mul-grpc-service/model"
)

type AddClientBinding struct {
	E endpoint.Endpoint
}

func (s *AddClientBinding) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, model.AddRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*model.AddResponse)
	return r.Ctx, r.V, nil
}

func (s *AddClientBinding) AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, model.AddRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*model.AddResponse)
	return r.Ctx, r.V, nil
}

func NewAddClient(cc *grpc.ClientConn) service.AddService {
	return &AddClientBinding{
		E: grpctransport.NewClient(
			cc,
			"pb.AddService",
			"Add",
			model.EncodeRequest,
			model.DecodeResponse,
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
			model.EncodeRequest,
			model.DecodeResponse,
			&pb.AddResponse{},
		).Endpoint(),
	}
}
