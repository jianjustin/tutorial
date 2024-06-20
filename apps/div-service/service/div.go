package service

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/div-grpc-service/pb"
	"time"
)

type DivServiceMiddleware func(DivService) DivService

type DivService interface {
	Div(ctx context.Context, a int64) (context.Context, int64, error)
}

type divService struct{}

func (divService) Div(ctx context.Context, a int64) (context.Context, int64, error) {
	return ctx, a / int64(2), nil
}

func NewDivService() DivService {
	return divService{}
}

type loggingDivServiceMiddleware struct {
	logger log.Logger
	next   DivService
}

func (M loggingDivServiceMiddleware) Div(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Div",
			"request", pb.DivRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	//time.Sleep(time.Duration(rand.Int()%10) * time.Second)
	return M.next.Div(ctx, a)
}

func LoggingDivServiceMiddleware(logger log.Logger) DivServiceMiddleware {
	return func(next DivService) DivService {
		return &loggingDivServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}
