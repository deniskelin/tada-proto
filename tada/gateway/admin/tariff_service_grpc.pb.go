// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	billing "github.com/deniskelin/tada-proto/tada/api/billing"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TariffClient is the client API for Tariff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TariffClient interface {
	GetTariffById(ctx context.Context, in *billing.GetTariffByIdRequest, opts ...grpc.CallOption) (*billing.GetTariffByIdResponse, error)
	GetTariffsList(ctx context.Context, in *billing.GetTariffsListRequest, opts ...grpc.CallOption) (*billing.GetTariffsListResponse, error)
	GetActiveTariffsList(ctx context.Context, in *billing.GetActiveTariffsListRequest, opts ...grpc.CallOption) (*billing.GetActiveTariffsListResponse, error)
	CreateTariff(ctx context.Context, in *billing.CreateTariffRequest, opts ...grpc.CallOption) (*billing.CreateTariffResponse, error)
	CloseTariff(ctx context.Context, in *billing.CloseTariffRequest, opts ...grpc.CallOption) (*billing.CloseTariffResponse, error)
	SetTariffAsDefault(ctx context.Context, in *billing.SetTariffAsDefaultRequest, opts ...grpc.CallOption) (*billing.SetTariffAsDefaultResponse, error)
	UpdateTariff(ctx context.Context, in *UpdateTariffRequest, opts ...grpc.CallOption) (*UpdateTariffResponse, error)
}

type tariffClient struct {
	cc grpc.ClientConnInterface
}

func NewTariffClient(cc grpc.ClientConnInterface) TariffClient {
	return &tariffClient{cc}
}

func (c *tariffClient) GetTariffById(ctx context.Context, in *billing.GetTariffByIdRequest, opts ...grpc.CallOption) (*billing.GetTariffByIdResponse, error) {
	out := new(billing.GetTariffByIdResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/GetTariffById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) GetTariffsList(ctx context.Context, in *billing.GetTariffsListRequest, opts ...grpc.CallOption) (*billing.GetTariffsListResponse, error) {
	out := new(billing.GetTariffsListResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/GetTariffsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) GetActiveTariffsList(ctx context.Context, in *billing.GetActiveTariffsListRequest, opts ...grpc.CallOption) (*billing.GetActiveTariffsListResponse, error) {
	out := new(billing.GetActiveTariffsListResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/GetActiveTariffsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) CreateTariff(ctx context.Context, in *billing.CreateTariffRequest, opts ...grpc.CallOption) (*billing.CreateTariffResponse, error) {
	out := new(billing.CreateTariffResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/CreateTariff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) CloseTariff(ctx context.Context, in *billing.CloseTariffRequest, opts ...grpc.CallOption) (*billing.CloseTariffResponse, error) {
	out := new(billing.CloseTariffResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/CloseTariff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) SetTariffAsDefault(ctx context.Context, in *billing.SetTariffAsDefaultRequest, opts ...grpc.CallOption) (*billing.SetTariffAsDefaultResponse, error) {
	out := new(billing.SetTariffAsDefaultResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/SetTariffAsDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffClient) UpdateTariff(ctx context.Context, in *UpdateTariffRequest, opts ...grpc.CallOption) (*UpdateTariffResponse, error) {
	out := new(UpdateTariffResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Tariff/UpdateTariff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TariffServer is the server API for Tariff service.
// All implementations must embed UnimplementedTariffServer
// for forward compatibility
type TariffServer interface {
	GetTariffById(context.Context, *billing.GetTariffByIdRequest) (*billing.GetTariffByIdResponse, error)
	GetTariffsList(context.Context, *billing.GetTariffsListRequest) (*billing.GetTariffsListResponse, error)
	GetActiveTariffsList(context.Context, *billing.GetActiveTariffsListRequest) (*billing.GetActiveTariffsListResponse, error)
	CreateTariff(context.Context, *billing.CreateTariffRequest) (*billing.CreateTariffResponse, error)
	CloseTariff(context.Context, *billing.CloseTariffRequest) (*billing.CloseTariffResponse, error)
	SetTariffAsDefault(context.Context, *billing.SetTariffAsDefaultRequest) (*billing.SetTariffAsDefaultResponse, error)
	UpdateTariff(context.Context, *UpdateTariffRequest) (*UpdateTariffResponse, error)
	mustEmbedUnimplementedTariffServer()
}

// UnimplementedTariffServer must be embedded to have forward compatible implementations.
type UnimplementedTariffServer struct {
}

func (UnimplementedTariffServer) GetTariffById(context.Context, *billing.GetTariffByIdRequest) (*billing.GetTariffByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTariffById not implemented")
}
func (UnimplementedTariffServer) GetTariffsList(context.Context, *billing.GetTariffsListRequest) (*billing.GetTariffsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTariffsList not implemented")
}
func (UnimplementedTariffServer) GetActiveTariffsList(context.Context, *billing.GetActiveTariffsListRequest) (*billing.GetActiveTariffsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveTariffsList not implemented")
}
func (UnimplementedTariffServer) CreateTariff(context.Context, *billing.CreateTariffRequest) (*billing.CreateTariffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTariff not implemented")
}
func (UnimplementedTariffServer) CloseTariff(context.Context, *billing.CloseTariffRequest) (*billing.CloseTariffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTariff not implemented")
}
func (UnimplementedTariffServer) SetTariffAsDefault(context.Context, *billing.SetTariffAsDefaultRequest) (*billing.SetTariffAsDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTariffAsDefault not implemented")
}
func (UnimplementedTariffServer) UpdateTariff(context.Context, *UpdateTariffRequest) (*UpdateTariffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTariff not implemented")
}
func (UnimplementedTariffServer) mustEmbedUnimplementedTariffServer() {}

// UnsafeTariffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TariffServer will
// result in compilation errors.
type UnsafeTariffServer interface {
	mustEmbedUnimplementedTariffServer()
}

func RegisterTariffServer(s grpc.ServiceRegistrar, srv TariffServer) {
	s.RegisterService(&Tariff_ServiceDesc, srv)
}

func _Tariff_GetTariffById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.GetTariffByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).GetTariffById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/GetTariffById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).GetTariffById(ctx, req.(*billing.GetTariffByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_GetTariffsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.GetTariffsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).GetTariffsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/GetTariffsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).GetTariffsList(ctx, req.(*billing.GetTariffsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_GetActiveTariffsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.GetActiveTariffsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).GetActiveTariffsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/GetActiveTariffsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).GetActiveTariffsList(ctx, req.(*billing.GetActiveTariffsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_CreateTariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.CreateTariffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).CreateTariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/CreateTariff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).CreateTariff(ctx, req.(*billing.CreateTariffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_CloseTariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.CloseTariffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).CloseTariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/CloseTariff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).CloseTariff(ctx, req.(*billing.CloseTariffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_SetTariffAsDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.SetTariffAsDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).SetTariffAsDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/SetTariffAsDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).SetTariffAsDefault(ctx, req.(*billing.SetTariffAsDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tariff_UpdateTariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTariffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServer).UpdateTariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Tariff/UpdateTariff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServer).UpdateTariff(ctx, req.(*UpdateTariffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tariff_ServiceDesc is the grpc.ServiceDesc for Tariff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tariff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.gateway.admin.Tariff",
	HandlerType: (*TariffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTariffById",
			Handler:    _Tariff_GetTariffById_Handler,
		},
		{
			MethodName: "GetTariffsList",
			Handler:    _Tariff_GetTariffsList_Handler,
		},
		{
			MethodName: "GetActiveTariffsList",
			Handler:    _Tariff_GetActiveTariffsList_Handler,
		},
		{
			MethodName: "CreateTariff",
			Handler:    _Tariff_CreateTariff_Handler,
		},
		{
			MethodName: "CloseTariff",
			Handler:    _Tariff_CloseTariff_Handler,
		},
		{
			MethodName: "SetTariffAsDefault",
			Handler:    _Tariff_SetTariffAsDefault_Handler,
		},
		{
			MethodName: "UpdateTariff",
			Handler:    _Tariff_UpdateTariff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/gateway/admin/tariff_service.proto",
}
