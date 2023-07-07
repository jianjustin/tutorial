package service

import "context"

type AddServiceMiddleware func(AddService) AddService

type AddService interface {
	Add(ctx context.Context, a int64) (context.Context, int64, error)
	AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error)
}

type addService struct{}

func (addService) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a + int64(3), nil
}

func (addService) AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a + int64(4), nil
}

func NewAddService() AddService {
	return addService{}
}
