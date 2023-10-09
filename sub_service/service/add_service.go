package service

import "context"

type AddServiceMiddleware func(AddService) AddService

type AddService interface {
	Add(ctx context.Context, a int64) (context.Context, int64, error)
	AddAfterMul(ctx context.Context, a int64) (context.Context, int64, error)
}
