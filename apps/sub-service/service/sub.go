package service

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/sub-grpc-service/pb"
	"time"
)

type SubServiceMiddleware func(SubService) SubService

type SubService interface {
	Sub(ctx context.Context, a int64) (context.Context, int64, error)
}

type subService struct{}

func (subService) Sub(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a - int64(3), nil
}

func NewSubService() SubService {
	return subService{}
}

func LoggingSubServiceMiddleware(logger log.Logger) SubServiceMiddleware {
	return func(next SubService) SubService {
		return &loggingSubServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingSubServiceMiddleware struct {
	logger log.Logger
	next   SubService
}

func (M loggingSubServiceMiddleware) Sub(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Sub",
			"request", pb.SubRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Sub(ctx, a)
}
