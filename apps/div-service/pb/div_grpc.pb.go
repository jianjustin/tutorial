// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc1
// source: div.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DivService_Div_FullMethodName = "/pb.DivService/Div"
)

// DivServiceClient is the client API for DivService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DivServiceClient interface {
	Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error)
}

type divServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDivServiceClient(cc grpc.ClientConnInterface) DivServiceClient {
	return &divServiceClient{cc}
}

func (c *divServiceClient) Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error) {
	out := new(DivResponse)
	err := c.cc.Invoke(ctx, DivService_Div_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivServiceServer is the server API for DivService service.
// All implementations should embed UnimplementedDivServiceServer
// for forward compatibility
type DivServiceServer interface {
	Div(context.Context, *DivRequest) (*DivResponse, error)
}

// UnimplementedDivServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDivServiceServer struct {
}

func (UnimplementedDivServiceServer) Div(context.Context, *DivRequest) (*DivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}

// UnsafeDivServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DivServiceServer will
// result in compilation errors.
type UnsafeDivServiceServer interface {
	mustEmbedUnimplementedDivServiceServer()
}

func RegisterDivServiceServer(s grpc.ServiceRegistrar, srv DivServiceServer) {
	s.RegisterService(&DivService_ServiceDesc, srv)
}

func _DivService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DivService_Div_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivServiceServer).Div(ctx, req.(*DivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DivService_ServiceDesc is the grpc.ServiceDesc for DivService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DivService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DivService",
	HandlerType: (*DivServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Div",
			Handler:    _DivService_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "div.proto",
}
