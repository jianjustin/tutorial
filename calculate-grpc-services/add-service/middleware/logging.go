package middleware

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/add-grpc-service/model"
	"go.guide/add-grpc-service/service"
	"time"
)

func LoggingAddServiceMiddleware(logger log.Logger) service.AddServiceMiddleware {
	return func(next service.AddService) service.AddService {
		return &loggingAddServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingAddServiceMiddleware struct {
	logger log.Logger
	next   service.AddService
}

func (M loggingAddServiceMiddleware) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Add",
			"request", model.AddRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Add(ctx, a)
}

func (M loggingAddServiceMiddleware) AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AddAfterMul",
			"request", model.AddRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.AddAfterMul(ctx, a)
}
