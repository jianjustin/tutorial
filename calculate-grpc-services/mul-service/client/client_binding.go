package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.guide/mul-grpc-service/model"
	"go.guide/mul-grpc-service/pb"
	"go.guide/mul-grpc-service/service"
	"google.golang.org/grpc"
)

type MulClientBinding struct {
	E endpoint.Endpoint
}

func (s *MulClientBinding) Mul(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, model.MulRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*model.MulResponse)
	return r.Ctx, r.V, nil
}

func (s *MulClientBinding) MulAfterAdd(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := s.E(ctx, model.MulRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}
	r := response.(*model.MulResponse)
	return r.Ctx, r.V, nil
}

func NewMulClient(cc *grpc.ClientConn) service.MulService {
	return &MulClientBinding{
		E: grpctransport.NewClient(
			cc,
			"pb.MulService",
			"Mul",
			model.EncodeRequest,
			model.DecodeResponse,
			&pb.MulResponse{},
		).Endpoint(),
	}
}

func NewMulAfterAddClient(cc *grpc.ClientConn) service.MulService {
	return &MulClientBinding{
		E: grpctransport.NewClient(
			cc,
			"pb.MulService",
			"MulAfterAdd",
			model.EncodeRequest,
			model.DecodeResponse,
			&pb.MulResponse{},
		).Endpoint(),
	}
}
