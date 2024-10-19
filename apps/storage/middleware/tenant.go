package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func GetTenantId(ctx context.Context) (tenantId string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata found in context")
	}
	tenantIds := md.Get("tenant_id")
	if len(tenantIds) == 0 {
		return "", errors.New("no tenant_id found in context")
	}
	if tenantIds[0] == "" {
		return "", errors.New("no tenant_id found in context")
	}
	return tenantIds[0], nil
}
