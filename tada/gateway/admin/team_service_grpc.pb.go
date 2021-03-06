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

// TeamClient is the client API for Team service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamClient interface {
	AddTeamOnPersonalAccount(ctx context.Context, in *billing.AddTeamOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.AddTeamOnPersonalAccountResponse, error)
	DeleteTeamFromPersonalAccount(ctx context.Context, in *billing.DeleteTeamFromPersonalAccountRequest, opts ...grpc.CallOption) (*billing.DeleteTeamFromPersonalAccountResponse, error)
	GetTeamsOnPersonalAccount(ctx context.Context, in *billing.GetTeamsOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.GetTeamsOnPersonalAccountResponse, error)
	CheckTeamOnPersonalAccount(ctx context.Context, in *billing.CheckTeamOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.CheckTeamOnPersonalAccountResponse, error)
	GetTeamBitrate(ctx context.Context, in *billing.GetTeamBitrateRequest, opts ...grpc.CallOption) (*billing.GetTeamBitrateResponse, error)
}

type teamClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamClient(cc grpc.ClientConnInterface) TeamClient {
	return &teamClient{cc}
}

func (c *teamClient) AddTeamOnPersonalAccount(ctx context.Context, in *billing.AddTeamOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.AddTeamOnPersonalAccountResponse, error) {
	out := new(billing.AddTeamOnPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Team/AddTeamOnPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) DeleteTeamFromPersonalAccount(ctx context.Context, in *billing.DeleteTeamFromPersonalAccountRequest, opts ...grpc.CallOption) (*billing.DeleteTeamFromPersonalAccountResponse, error) {
	out := new(billing.DeleteTeamFromPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Team/DeleteTeamFromPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) GetTeamsOnPersonalAccount(ctx context.Context, in *billing.GetTeamsOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.GetTeamsOnPersonalAccountResponse, error) {
	out := new(billing.GetTeamsOnPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Team/GetTeamsOnPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) CheckTeamOnPersonalAccount(ctx context.Context, in *billing.CheckTeamOnPersonalAccountRequest, opts ...grpc.CallOption) (*billing.CheckTeamOnPersonalAccountResponse, error) {
	out := new(billing.CheckTeamOnPersonalAccountResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Team/CheckTeamOnPersonalAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) GetTeamBitrate(ctx context.Context, in *billing.GetTeamBitrateRequest, opts ...grpc.CallOption) (*billing.GetTeamBitrateResponse, error) {
	out := new(billing.GetTeamBitrateResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.Team/GetTeamBitrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServer is the server API for Team service.
// All implementations must embed UnimplementedTeamServer
// for forward compatibility
type TeamServer interface {
	AddTeamOnPersonalAccount(context.Context, *billing.AddTeamOnPersonalAccountRequest) (*billing.AddTeamOnPersonalAccountResponse, error)
	DeleteTeamFromPersonalAccount(context.Context, *billing.DeleteTeamFromPersonalAccountRequest) (*billing.DeleteTeamFromPersonalAccountResponse, error)
	GetTeamsOnPersonalAccount(context.Context, *billing.GetTeamsOnPersonalAccountRequest) (*billing.GetTeamsOnPersonalAccountResponse, error)
	CheckTeamOnPersonalAccount(context.Context, *billing.CheckTeamOnPersonalAccountRequest) (*billing.CheckTeamOnPersonalAccountResponse, error)
	GetTeamBitrate(context.Context, *billing.GetTeamBitrateRequest) (*billing.GetTeamBitrateResponse, error)
	mustEmbedUnimplementedTeamServer()
}

// UnimplementedTeamServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServer struct {
}

func (UnimplementedTeamServer) AddTeamOnPersonalAccount(context.Context, *billing.AddTeamOnPersonalAccountRequest) (*billing.AddTeamOnPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeamOnPersonalAccount not implemented")
}
func (UnimplementedTeamServer) DeleteTeamFromPersonalAccount(context.Context, *billing.DeleteTeamFromPersonalAccountRequest) (*billing.DeleteTeamFromPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeamFromPersonalAccount not implemented")
}
func (UnimplementedTeamServer) GetTeamsOnPersonalAccount(context.Context, *billing.GetTeamsOnPersonalAccountRequest) (*billing.GetTeamsOnPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsOnPersonalAccount not implemented")
}
func (UnimplementedTeamServer) CheckTeamOnPersonalAccount(context.Context, *billing.CheckTeamOnPersonalAccountRequest) (*billing.CheckTeamOnPersonalAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTeamOnPersonalAccount not implemented")
}
func (UnimplementedTeamServer) GetTeamBitrate(context.Context, *billing.GetTeamBitrateRequest) (*billing.GetTeamBitrateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamBitrate not implemented")
}
func (UnimplementedTeamServer) mustEmbedUnimplementedTeamServer() {}

// UnsafeTeamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServer will
// result in compilation errors.
type UnsafeTeamServer interface {
	mustEmbedUnimplementedTeamServer()
}

func RegisterTeamServer(s grpc.ServiceRegistrar, srv TeamServer) {
	s.RegisterService(&Team_ServiceDesc, srv)
}

func _Team_AddTeamOnPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.AddTeamOnPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).AddTeamOnPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Team/AddTeamOnPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).AddTeamOnPersonalAccount(ctx, req.(*billing.AddTeamOnPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_DeleteTeamFromPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.DeleteTeamFromPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).DeleteTeamFromPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Team/DeleteTeamFromPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).DeleteTeamFromPersonalAccount(ctx, req.(*billing.DeleteTeamFromPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_GetTeamsOnPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.GetTeamsOnPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).GetTeamsOnPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Team/GetTeamsOnPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).GetTeamsOnPersonalAccount(ctx, req.(*billing.GetTeamsOnPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_CheckTeamOnPersonalAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.CheckTeamOnPersonalAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).CheckTeamOnPersonalAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Team/CheckTeamOnPersonalAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).CheckTeamOnPersonalAccount(ctx, req.(*billing.CheckTeamOnPersonalAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_GetTeamBitrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.GetTeamBitrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).GetTeamBitrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.Team/GetTeamBitrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).GetTeamBitrate(ctx, req.(*billing.GetTeamBitrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Team_ServiceDesc is the grpc.ServiceDesc for Team service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Team_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.gateway.admin.Team",
	HandlerType: (*TeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTeamOnPersonalAccount",
			Handler:    _Team_AddTeamOnPersonalAccount_Handler,
		},
		{
			MethodName: "DeleteTeamFromPersonalAccount",
			Handler:    _Team_DeleteTeamFromPersonalAccount_Handler,
		},
		{
			MethodName: "GetTeamsOnPersonalAccount",
			Handler:    _Team_GetTeamsOnPersonalAccount_Handler,
		},
		{
			MethodName: "CheckTeamOnPersonalAccount",
			Handler:    _Team_CheckTeamOnPersonalAccount_Handler,
		},
		{
			MethodName: "GetTeamBitrate",
			Handler:    _Team_GetTeamBitrate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/gateway/admin/team_service.proto",
}
