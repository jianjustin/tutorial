package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "github.com/jianjustin/add/proto"
)

type Add struct{}

func (e *Add) Add(ctx context.Context, req *pb.AddRequest, rsp *pb.AddResponse) error {
	logger.Infof("Received Add.Add request: %v", req)
	rsp.Result = req.A + req.B
	return nil
}
