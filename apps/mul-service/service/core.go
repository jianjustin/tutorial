package service

import "context"

type CoreServiceMiddleware func(CoreService) CoreService

type CoreService interface {
	Random(ctx context.Context, a int64) (context.Context, int64, error)
}
