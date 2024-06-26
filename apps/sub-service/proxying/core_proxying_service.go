package proxying

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"go.guide/sub-grpc-service/pb"
)

type CodeServiceProxy struct {
	ctx context.Context
	E   endpoint.Endpoint
}

func (proxy CodeServiceProxy) Random(ctx context.Context, a int64) (context.Context, int64, error) {
	response, err := proxy.E(proxy.ctx, pb.RandomRequest{A: a})
	if err != nil {
		return ctx, 0, err
	}

	str, _ := json.Marshal(response)

	resp := &pb.RandomResponse{}
	err = json.Unmarshal(str, resp)
	return ctx, resp.V, nil
}
