// GENERATED BY go:generate. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: authwise/integration/v1alpha1/grpc.proto

package v1alpha1

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
	AuthwiseIntegrationService_ResolveProvider_FullMethodName  = "/authwise.integration.v1alpha1.AuthwiseIntegrationService/ResolveProvider"
	AuthwiseIntegrationService_Login_FullMethodName            = "/authwise.integration.v1alpha1.AuthwiseIntegrationService/Login"
	AuthwiseIntegrationService_UserPostProcess_FullMethodName  = "/authwise.integration.v1alpha1.AuthwiseIntegrationService/UserPostProcess"
	AuthwiseIntegrationService_GrantPostProcess_FullMethodName = "/authwise.integration.v1alpha1.AuthwiseIntegrationService/GrantPostProcess"
)

// AuthwiseIntegrationServiceClient is the client API for AuthwiseIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthwiseIntegrationServiceClient interface {
	ResolveProvider(ctx context.Context, in *ResolveProviderRequest, opts ...grpc.CallOption) (*ResolveProviderResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UserPostProcess(ctx context.Context, in *UserPostProcessRequest, opts ...grpc.CallOption) (*UserPostProcessResponse, error)
	GrantPostProcess(ctx context.Context, in *GrantPostProcessRequest, opts ...grpc.CallOption) (*GrantPostProcessResponse, error)
}

type authwiseIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthwiseIntegrationServiceClient(cc grpc.ClientConnInterface) AuthwiseIntegrationServiceClient {
	return &authwiseIntegrationServiceClient{cc}
}

func (c *authwiseIntegrationServiceClient) ResolveProvider(ctx context.Context, in *ResolveProviderRequest, opts ...grpc.CallOption) (*ResolveProviderResponse, error) {
	out := new(ResolveProviderResponse)
	err := c.cc.Invoke(ctx, AuthwiseIntegrationService_ResolveProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authwiseIntegrationServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthwiseIntegrationService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authwiseIntegrationServiceClient) UserPostProcess(ctx context.Context, in *UserPostProcessRequest, opts ...grpc.CallOption) (*UserPostProcessResponse, error) {
	out := new(UserPostProcessResponse)
	err := c.cc.Invoke(ctx, AuthwiseIntegrationService_UserPostProcess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authwiseIntegrationServiceClient) GrantPostProcess(ctx context.Context, in *GrantPostProcessRequest, opts ...grpc.CallOption) (*GrantPostProcessResponse, error) {
	out := new(GrantPostProcessResponse)
	err := c.cc.Invoke(ctx, AuthwiseIntegrationService_GrantPostProcess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthwiseIntegrationServiceServer is the server API for AuthwiseIntegrationService service.
// All implementations should embed UnimplementedAuthwiseIntegrationServiceServer
// for forward compatibility
type AuthwiseIntegrationServiceServer interface {
	ResolveProvider(context.Context, *ResolveProviderRequest) (*ResolveProviderResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UserPostProcess(context.Context, *UserPostProcessRequest) (*UserPostProcessResponse, error)
	GrantPostProcess(context.Context, *GrantPostProcessRequest) (*GrantPostProcessResponse, error)
}

// UnimplementedAuthwiseIntegrationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthwiseIntegrationServiceServer struct {
}

func (UnimplementedAuthwiseIntegrationServiceServer) ResolveProvider(context.Context, *ResolveProviderRequest) (*ResolveProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveProvider not implemented")
}
func (UnimplementedAuthwiseIntegrationServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthwiseIntegrationServiceServer) UserPostProcess(context.Context, *UserPostProcessRequest) (*UserPostProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPostProcess not implemented")
}
func (UnimplementedAuthwiseIntegrationServiceServer) GrantPostProcess(context.Context, *GrantPostProcessRequest) (*GrantPostProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantPostProcess not implemented")
}

// UnsafeAuthwiseIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthwiseIntegrationServiceServer will
// result in compilation errors.
type UnsafeAuthwiseIntegrationServiceServer interface {
	mustEmbedUnimplementedAuthwiseIntegrationServiceServer()
}

func RegisterAuthwiseIntegrationServiceServer(s grpc.ServiceRegistrar, srv AuthwiseIntegrationServiceServer) {
	s.RegisterService(&AuthwiseIntegrationService_ServiceDesc, srv)
}

func _AuthwiseIntegrationService_ResolveProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthwiseIntegrationServiceServer).ResolveProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthwiseIntegrationService_ResolveProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthwiseIntegrationServiceServer).ResolveProvider(ctx, req.(*ResolveProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthwiseIntegrationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthwiseIntegrationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthwiseIntegrationService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthwiseIntegrationServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthwiseIntegrationService_UserPostProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPostProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthwiseIntegrationServiceServer).UserPostProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthwiseIntegrationService_UserPostProcess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthwiseIntegrationServiceServer).UserPostProcess(ctx, req.(*UserPostProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthwiseIntegrationService_GrantPostProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantPostProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthwiseIntegrationServiceServer).GrantPostProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthwiseIntegrationService_GrantPostProcess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthwiseIntegrationServiceServer).GrantPostProcess(ctx, req.(*GrantPostProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthwiseIntegrationService_ServiceDesc is the grpc.ServiceDesc for AuthwiseIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthwiseIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authwise.integration.v1alpha1.AuthwiseIntegrationService",
	HandlerType: (*AuthwiseIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveProvider",
			Handler:    _AuthwiseIntegrationService_ResolveProvider_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthwiseIntegrationService_Login_Handler,
		},
		{
			MethodName: "UserPostProcess",
			Handler:    _AuthwiseIntegrationService_UserPostProcess_Handler,
		},
		{
			MethodName: "GrantPostProcess",
			Handler:    _AuthwiseIntegrationService_GrantPostProcess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authwise/integration/v1alpha1/grpc.proto",
}
