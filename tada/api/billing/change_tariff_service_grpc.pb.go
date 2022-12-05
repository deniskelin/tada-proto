// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/api/billing/change_tariff_service.proto

package billing

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

// ChangeTariffClient is the client API for ChangeTariff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChangeTariffClient interface {
	GetChangesTariffsByPersonalAccount(ctx context.Context, in *GetChangesTariffsByPersonalAccountRequest, opts ...grpc.CallOption) (*GetChangesTariffsByPersonalAccountResponse, error)
	CreateChangeTariffOnPersonalAccount(ctx context.Context, in *CreateChangeTariffOnPersonalAccountRequest, opts ...grpc.CallOption) (*CreateChangeTariffOnPersonalAccountResponse, error)
	DeleteChangeTariffOnPersonalAccount(ctx context.Context, in *DeleteChangeTariffOnPersonalAccountRequest, opts ...grpc.CallOption) (*DeleteChangeTariffOnPersonalAccountResponse, error)
}

type changeTariffClient struct {
	cc grpc.ClientConnInterface
}

func NewChangeTariffClient(cc grpc.ClientConnInterface) ChangeTariffClient {
	return &changeTariffClient{cc}
}

func (c *changeTariffClient) GetChangesTariffsByPersonalAccount(ctx context.Context, in *GetChangesTariffsByPersonalAccountRequest, opts ...grpc.CallOption) (*GetChangesTariffsByPersonalAccountResponse, error) {
	out := new(GetChangesTariffsByPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.api.billing.ChangeTariff/GetChangesTariffsByPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *changeTariffClient) CreateChangeTariffOnPersonalAccount(ctx context.Context, in *CreateChangeTariffOnPersonalAccountRequest, opts ...grpc.CallOption) (*CreateChangeTariffOnPersonalAccountResponse, error) {
	out := new(CreateChangeTariffOnPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.api.billing.ChangeTariff/CreateChangeTariffOnPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *changeTariffClient) DeleteChangeTariffOnPersonalAccount(ctx context.Context, in *DeleteChangeTariffOnPersonalAccountRequest, opts ...grpc.CallOption) (*DeleteChangeTariffOnPersonalAccountResponse, error) {
	out := new(DeleteChangeTariffOnPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.api.billing.ChangeTariff/DeleteChangeTariffOnPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeTariffServer is the server API for ChangeTariff service.
// All implementations must embed UnimplementedChangeTariffServer
// for forward compatibility
type ChangeTariffServer interface {
	GetChangesTariffsByPersonalAccount(context.Context, *GetChangesTariffsByPersonalAccountRequest) (*GetChangesTariffsByPersonalAccountResponse, error)
	CreateChangeTariffOnPersonalAccount(context.Context, *CreateChangeTariffOnPersonalAccountRequest) (*CreateChangeTariffOnPersonalAccountResponse, error)
	DeleteChangeTariffOnPersonalAccount(context.Context, *DeleteChangeTariffOnPersonalAccountRequest) (*DeleteChangeTariffOnPersonalAccountResponse, error)
	mustEmbedUnimplementedChangeTariffServer()
}

// UnimplementedChangeTariffServer must be embedded to have forward compatible implementations.
type UnimplementedChangeTariffServer struct {
}

func (UnimplementedChangeTariffServer) GetChangesTariffsByPersonalAccount(context.Context, *GetChangesTariffsByPersonalAccountRequest) (*GetChangesTariffsByPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangesTariffsByPersonalAccount not implemented")
}
func (UnimplementedChangeTariffServer) CreateChangeTariffOnPersonalAccount(context.Context, *CreateChangeTariffOnPersonalAccountRequest) (*CreateChangeTariffOnPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChangeTariffOnPersonalAccount not implemented")
}
func (UnimplementedChangeTariffServer) DeleteChangeTariffOnPersonalAccount(context.Context, *DeleteChangeTariffOnPersonalAccountRequest) (*DeleteChangeTariffOnPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChangeTariffOnPersonalAccount not implemented")
}
func (UnimplementedChangeTariffServer) mustEmbedUnimplementedChangeTariffServer() {}

// UnsafeChangeTariffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChangeTariffServer will
// result in compilation errors.
type UnsafeChangeTariffServer interface {
	mustEmbedUnimplementedChangeTariffServer()
}

func RegisterChangeTariffServer(s grpc.ServiceRegistrar, srv ChangeTariffServer) {
	s.RegisterService(&ChangeTariff_ServiceDesc, srv)
}

func _ChangeTariff_GetChangesTariffsByPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangesTariffsByPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeTariffServer).GetChangesTariffsByPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.billing.ChangeTariff/GetChangesTariffsByPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeTariffServer).GetChangesTariffsByPersonalAccount(ctx, req.(*GetChangesTariffsByPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChangeTariff_CreateChangeTariffOnPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChangeTariffOnPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeTariffServer).CreateChangeTariffOnPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.billing.ChangeTariff/CreateChangeTariffOnPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeTariffServer).CreateChangeTariffOnPersonalAccount(ctx, req.(*CreateChangeTariffOnPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChangeTariff_DeleteChangeTariffOnPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChangeTariffOnPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeTariffServer).DeleteChangeTariffOnPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.billing.ChangeTariff/DeleteChangeTariffOnPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeTariffServer).DeleteChangeTariffOnPersonalAccount(ctx, req.(*DeleteChangeTariffOnPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChangeTariff_ServiceDesc is the grpc.ServiceDesc for ChangeTariff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChangeTariff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.billing.ChangeTariff",
	HandlerType: (*ChangeTariffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChangesTariffsByPersonalAccount",
			Handler:    _ChangeTariff_GetChangesTariffsByPersonalAccount_Handler,
		},
		{
			MethodName: "CreateChangeTariffOnPersonalAccount",
			Handler:    _ChangeTariff_CreateChangeTariffOnPersonalAccount_Handler,
		},
		{
			MethodName: "DeleteChangeTariffOnPersonalAccount",
			Handler:    _ChangeTariff_DeleteChangeTariffOnPersonalAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/billing/change_tariff_service.proto",
}
