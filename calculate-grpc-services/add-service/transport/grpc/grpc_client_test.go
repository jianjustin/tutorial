package transportgrpc_test

import (
	"context"
	transportgrpc "go.guide/add-grpc-service/transport/grpc"
	"google.golang.org/grpc"
	"testing"
)

func TestForGrpcService(t *testing.T) {
	cc, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	client := transportgrpc.NewAddClient(cc)
	_, _, err = client.Add(context.Background(), int64(42))
	if err != nil {
		t.Fatalf("unable to Test: %+v", err)
	}

	client1 := transportgrpc.NewAddAfterMulClient(cc)
	_, _, err = client1.AddAfterMul(context.Background(), int64(42))
	if err != nil {
		t.Fatalf("unable to Test: %+v", err)
	}
}
