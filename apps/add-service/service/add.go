package service

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/add-grpc-service/pb"
	"time"
)

type AddServiceMiddleware func(AddService) AddService

type AddService interface {
	Add(ctx context.Context, a int64) (context.Context, int64, error)
}

type addService struct{}

func (addService) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a + int64(3), nil
}

func NewAddService() AddService {
	return addService{}
}

func LoggingAddServiceMiddleware(logger log.Logger) AddServiceMiddleware {
	return func(next AddService) AddService {
		return &loggingAddServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingAddServiceMiddleware struct {
	logger log.Logger
	next   AddService
}

func (M loggingAddServiceMiddleware) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Add",
			"request", pb.AddRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Add(ctx, a)
}
