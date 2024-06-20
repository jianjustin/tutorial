package service

import "context"

type MulService interface {
	Mul(ctx context.Context, a int64) (context.Context, int64, error)
	MulAfterAdd(ctx context.Context, a int64) (context.Context, int64, error)
}

type mulService struct{}

func (mulService) Mul(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a * int64(3), nil
}

func (mulService) MulAfterAdd(ctx context.Context, a int64) (context.Context, int64, error) {
	return nil, a * int64(4), nil
}

func NewMulService() MulService {
	return mulService{}
}
