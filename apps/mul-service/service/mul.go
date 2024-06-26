package service

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/mul-grpc-service/pb"
	"time"
)

type MulServiceMiddleware func(MulService) MulService

type MulService interface {
	Mul(ctx context.Context, a int64) (context.Context, int64, error)
}

type mulService struct{}

func (mulService) Mul(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a * int64(3), nil
}

func NewMulService() MulService {
	return mulService{}
}

func LoggingMulServiceMiddleware(logger log.Logger) MulServiceMiddleware {
	return func(next MulService) MulService {
		return &loggingMulServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMulServiceMiddleware struct {
	logger log.Logger
	next   MulService
}

func (M loggingMulServiceMiddleware) Mul(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Mul",
			"request", pb.MulRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Mul(ctx, a)
}
