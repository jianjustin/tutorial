package handler

import (
	"context"
	"github.com/jianjustin/sub/proto/mul"
	"go-micro.dev/v4/logger"

	pb "github.com/jianjustin/sub/proto/sub"
)

type Sub struct{}

func (e *Sub) Sub(ctx context.Context, req *pb.SubRequest, rsp *pb.SubResponse) error {
	logger.Infof("Received Sub.Call request: %v", req)
	rsp.Result = req.A - req.B
	// call mul service
	r, err := MulClient.Mul(ctx, &mul.MulRequest{A: req.A, B: req.B})
	if err != nil {
		return err
	}
	rsp.Result = rsp.Result + r.Result
	return nil
}
