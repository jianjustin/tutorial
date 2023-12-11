package handler

import (
	"context"
	"go-micro.dev/v4/logger"

	pb "github.com/jianjustin/sub/proto"
)

type Sub struct{}

func (e *Sub) Sub(ctx context.Context, req *pb.SubRequest, rsp *pb.SubResponse) error {
	logger.Infof("Received Sub.Call request: %v", req)
	rsp.Result = req.A - req.B
	return nil
}
