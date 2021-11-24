// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Get(ctx context.Context, in *Health, opts ...grpc.CallOption) (*Health, error)
	GetLogs(ctx context.Context, in *Logs, opts ...grpc.CallOption) (*Logs, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Get(ctx context.Context, in *Health, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/v1.health/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) GetLogs(ctx context.Context, in *Logs, opts ...grpc.CallOption) (*Logs, error) {
	out := new(Logs)
	err := c.cc.Invoke(ctx, "/v1.health/GetLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Get(context.Context, *Health) (*Health, error)
	GetLogs(context.Context, *Logs) (*Logs, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Get(context.Context, *Health) (*Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHealthServer) GetLogs(context.Context, *Logs) (*Logs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Health)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.health/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Get(ctx, req.(*Health))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Logs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.health/GetLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).GetLogs(ctx, req.(*Logs))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Health_Get_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _Health_GetLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}