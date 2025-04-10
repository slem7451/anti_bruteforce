// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: AppService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Auth_Auth_FullMethodName                = "/auth.Auth/Auth"
	Auth_Reset_FullMethodName               = "/auth.Auth/Reset"
	Auth_AddToBlacklist_FullMethodName      = "/auth.Auth/AddToBlacklist"
	Auth_AddToWhitelist_FullMethodName      = "/auth.Auth/AddToWhitelist"
	Auth_DeleteFromBlacklist_FullMethodName = "/auth.Auth/DeleteFromBlacklist"
	Auth_DeleteFromWhitelist_FullMethodName = "/auth.Auth/DeleteFromWhitelist"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Auth(ctx context.Context, in *Credits, opts ...grpc.CallOption) (*Response, error)
	Reset(ctx context.Context, in *Credits, opts ...grpc.CallOption) (*Response, error)
	AddToBlacklist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error)
	AddToWhitelist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error)
	DeleteFromBlacklist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error)
	DeleteFromWhitelist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Auth(ctx context.Context, in *Credits, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Reset(ctx context.Context, in *Credits, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_Reset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddToBlacklist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_AddToBlacklist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddToWhitelist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_AddToWhitelist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteFromBlacklist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_DeleteFromBlacklist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteFromWhitelist(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Auth_DeleteFromWhitelist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	Auth(context.Context, *Credits) (*Response, error)
	Reset(context.Context, *Credits) (*Response, error)
	AddToBlacklist(context.Context, *Subnet) (*Response, error)
	AddToWhitelist(context.Context, *Subnet) (*Response, error)
	DeleteFromBlacklist(context.Context, *Subnet) (*Response, error)
	DeleteFromWhitelist(context.Context, *Subnet) (*Response, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) Auth(context.Context, *Credits) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedAuthServer) Reset(context.Context, *Credits) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedAuthServer) AddToBlacklist(context.Context, *Subnet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBlacklist not implemented")
}
func (UnimplementedAuthServer) AddToWhitelist(context.Context, *Subnet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhitelist not implemented")
}
func (UnimplementedAuthServer) DeleteFromBlacklist(context.Context, *Subnet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFromBlacklist not implemented")
}
func (UnimplementedAuthServer) DeleteFromWhitelist(context.Context, *Subnet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFromWhitelist not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credits)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Auth(ctx, req.(*Credits))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credits)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Reset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Reset(ctx, req.(*Credits))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddToBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddToBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddToBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddToBlacklist(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddToWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddToWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddToWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddToWhitelist(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteFromBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteFromBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteFromBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteFromBlacklist(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteFromWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteFromWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteFromWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteFromWhitelist(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Auth_Auth_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Auth_Reset_Handler,
		},
		{
			MethodName: "AddToBlacklist",
			Handler:    _Auth_AddToBlacklist_Handler,
		},
		{
			MethodName: "AddToWhitelist",
			Handler:    _Auth_AddToWhitelist_Handler,
		},
		{
			MethodName: "DeleteFromBlacklist",
			Handler:    _Auth_DeleteFromBlacklist_Handler,
		},
		{
			MethodName: "DeleteFromWhitelist",
			Handler:    _Auth_DeleteFromWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AppService.proto",
}
