package endpoints_test

import (
	"context"
	"fmt"
	"testing"
)

func TestForHandlePriceChain(t *testing.T) {
}

type Goods struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PriceRes struct {
	Price int64 `json:"price"`
}

type HandlePriceFunc func(ctx context.Context, request *PriceRes) (response *PriceRes, err error)
type HandlePriceMiddleware func(HandlePriceFunc) HandlePriceFunc

func HandleAddPrice(ctx context.Context, request *PriceRes) (response *PriceRes, err error) {
	response.Price = request.Price + 1
	fmt.Printf("HandleAddPrice: %d\n", request.Price)
	defer fmt.Printf("HandleAddPrice res: %d\n", response.Price)
	return
}

func HandleSubPrice(ctx context.Context, request *PriceRes) (response *PriceRes, err error) {
	response.Price = request.Price - 2
	fmt.Printf("HandleSubPrice: %d\n", response.Price)
	defer fmt.Printf("HandleSubPrice res: %d\n", response.Price)
	return
}

func Chain(m ...HandlePriceMiddleware) HandlePriceMiddleware {
	return func(next HandlePriceFunc) HandlePriceFunc {
		for i := len(m) - 1; i >= 0; i-- { // reverse
			next = m[i](next)
		}
		return next
	}
}
