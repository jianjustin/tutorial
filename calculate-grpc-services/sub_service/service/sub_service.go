package service

import "context"

type SubServiceMiddleware func(SubService) SubService

type SubService interface {
	Sub(ctx context.Context, a int64) (context.Context, int64, error)
	SubAfterAdd(ctx context.Context, a int64) (context.Context, int64, error)
}

type subService struct{}

func (subService) Sub(ctx context.Context, a int64) (context.Context, int64, error) {
	return ctx, a - int64(1), nil
}

func (subService) SubAfterAdd(ctx context.Context, a int64) (context.Context, int64, error) {
	return ctx, a + int64(100) - int64(1), nil
}

func NewSubService() SubService {
	return subService{}
}
