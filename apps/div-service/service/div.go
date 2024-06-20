package service

import "context"

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
