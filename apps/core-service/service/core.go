package service

import (
	"context"
	"github.com/go-kit/log"
	"go.guide/core-service/pb"
	"math/rand"
	"time"
)

type CoreServiceMiddleware func(CoreService) CoreService

type CoreService interface {
	Random(ctx context.Context, a int64) (context.Context, int64, error)
}

type coreService struct{}

func (coreService) Random(ctx context.Context, a int64) (context.Context, int64, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := rand.Int63n(11) // 生成0到10的随机整数
	return ctx, a + randomNum, nil
}

func NewCoreService() CoreService {
	return coreService{}
}

func LoggingCoreServiceMiddleware(logger log.Logger) CoreServiceMiddleware {
	return func(next CoreService) CoreService {
		return &loggingCoreServiceMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingCoreServiceMiddleware struct {
	logger log.Logger
	next   CoreService
}

func (M loggingCoreServiceMiddleware) Random(ctx context.Context, a int64) (context.Context, int64, error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Random",
			"request", pb.RandomRequest{A: a},
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Random(ctx, a)
}
