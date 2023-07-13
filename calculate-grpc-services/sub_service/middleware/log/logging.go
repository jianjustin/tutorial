package log

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/sub-grpc-service/service"
	"go.guide/sub-grpc-service/transport"
	"time"
)

func LoggingSubServiceMiddleware(logger log.Logger) service.SubServiceMiddleware {
	return func(next service.SubService) service.SubService {
		return &loggingSubServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingSubServiceMiddleware struct {
	logger log.Logger
	next   service.SubService
}

func (M loggingSubServiceMiddleware) Sub(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Sub",
			"request", transport.SubRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Sub(ctx, a)
}

func (M loggingSubServiceMiddleware) SubAfterAdd(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "SubAfterAdd",
			"request", transport.SubRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.SubAfterAdd(ctx, a)
}
