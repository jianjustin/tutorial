package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "github.com/jianjustin/mul/proto"
)

type Mul struct{}

func (e *Mul) Mul(ctx context.Context, req *pb.MulRequest, rsp *pb.MulResponse) error {
	logger.Infof("Received Mul.Call request: %v", req)
	rsp.Result = req.A * req.B
	return nil
}
