package proxying

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"go.guide/div-grpc-service/pb"
)

type AddServiceProxy struct {
	ctx context.Context
	E   endpoint.Endpoint
}

func (proxy AddServiceProxy) Add(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := proxy.E(proxy.ctx, pb.AddRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}

	str, _ := json.Marshal(response)

	resp := &pb.AddResponse{}
	err = json.Unmarshal(str, resp)
	return ctx, resp.V, nil
}
