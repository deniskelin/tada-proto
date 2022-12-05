// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/api/customer/user_service.proto

package customer

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

// CustomerUserClient is the client API for CustomerUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerUserClient interface {
	GetUsersInfoByUserUUIDArray(ctx context.Context, in *GetUsersInfoByUserUUIDArrayRequest, opts ...grpc.CallOption) (*UsersInfo, error)
	GetUsersInfoByUserUUIDArrayExcludingTeamMembers(ctx context.Context, in *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest, opts ...grpc.CallOption) (*UsersInfo, error)
	GetUserUUIDByContact(ctx context.Context, in *GetUserUUIDByContactRequest, opts ...grpc.CallOption) (*GetUserUUIDByContactResponse, error)
	GetContactsByUserAndTeams(ctx context.Context, in *GetContactsByUserAndTeamsRequest, opts ...grpc.CallOption) (*GetContactsByUserAndTeamsResponse, error)
}

type customerUserClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserClient(cc grpc.ClientConnInterface) CustomerUserClient {
	return &customerUserClient{cc}
}

func (c *customerUserClient) GetUsersInfoByUserUUIDArray(ctx context.Context, in *GetUsersInfoByUserUUIDArrayRequest, opts ...grpc.CallOption) (*UsersInfo, error) {
	out := new(UsersInfo)
	err := c.cc.Invoke(ctx, "/tada.api.customer.CustomerUser/GetUsersInfoByUserUUIDArray", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserClient) GetUsersInfoByUserUUIDArrayExcludingTeamMembers(ctx context.Context, in *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest, opts ...grpc.CallOption) (*UsersInfo, error) {
	out := new(UsersInfo)
	err := c.cc.Invoke(ctx, "/tada.api.customer.CustomerUser/GetUsersInfoByUserUUIDArrayExcludingTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserClient) GetUserUUIDByContact(ctx context.Context, in *GetUserUUIDByContactRequest, opts ...grpc.CallOption) (*GetUserUUIDByContactResponse, error) {
	out := new(GetUserUUIDByContactResponse)
	err := c.cc.Invoke(ctx, "/tada.api.customer.CustomerUser/GetUserUUIDByContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserClient) GetContactsByUserAndTeams(ctx context.Context, in *GetContactsByUserAndTeamsRequest, opts ...grpc.CallOption) (*GetContactsByUserAndTeamsResponse, error) {
	out := new(GetContactsByUserAndTeamsResponse)
	err := c.cc.Invoke(ctx, "/tada.api.customer.CustomerUser/GetContactsByUserAndTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserServer is the server API for CustomerUser service.
// All implementations must embed UnimplementedCustomerUserServer
// for forward compatibility
type CustomerUserServer interface {
	GetUsersInfoByUserUUIDArray(context.Context, *GetUsersInfoByUserUUIDArrayRequest) (*UsersInfo, error)
	GetUsersInfoByUserUUIDArrayExcludingTeamMembers(context.Context, *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) (*UsersInfo, error)
	GetUserUUIDByContact(context.Context, *GetUserUUIDByContactRequest) (*GetUserUUIDByContactResponse, error)
	GetContactsByUserAndTeams(context.Context, *GetContactsByUserAndTeamsRequest) (*GetContactsByUserAndTeamsResponse, error)
	mustEmbedUnimplementedCustomerUserServer()
}

// UnimplementedCustomerUserServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerUserServer struct {
}

func (UnimplementedCustomerUserServer) GetUsersInfoByUserUUIDArray(context.Context, *GetUsersInfoByUserUUIDArrayRequest) (*UsersInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInfoByUserUUIDArray not implemented")
}
func (UnimplementedCustomerUserServer) GetUsersInfoByUserUUIDArrayExcludingTeamMembers(context.Context, *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) (*UsersInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInfoByUserUUIDArrayExcludingTeamMembers not implemented")
}
func (UnimplementedCustomerUserServer) GetUserUUIDByContact(context.Context, *GetUserUUIDByContactRequest) (*GetUserUUIDByContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserUUIDByContact not implemented")
}
func (UnimplementedCustomerUserServer) GetContactsByUserAndTeams(context.Context, *GetContactsByUserAndTeamsRequest) (*GetContactsByUserAndTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactsByUserAndTeams not implemented")
}
func (UnimplementedCustomerUserServer) mustEmbedUnimplementedCustomerUserServer() {}

// UnsafeCustomerUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerUserServer will
// result in compilation errors.
type UnsafeCustomerUserServer interface {
	mustEmbedUnimplementedCustomerUserServer()
}

func RegisterCustomerUserServer(s grpc.ServiceRegistrar, srv CustomerUserServer) {
	s.RegisterService(&CustomerUser_ServiceDesc, srv)
}

func _CustomerUser_GetUsersInfoByUserUUIDArray_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInfoByUserUUIDArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserServer).GetUsersInfoByUserUUIDArray(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.customer.CustomerUser/GetUsersInfoByUserUUIDArray",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserServer).GetUsersInfoByUserUUIDArray(ctx, req.(*GetUsersInfoByUserUUIDArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUser_GetUsersInfoByUserUUIDArrayExcludingTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserServer).GetUsersInfoByUserUUIDArrayExcludingTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.customer.CustomerUser/GetUsersInfoByUserUUIDArrayExcludingTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserServer).GetUsersInfoByUserUUIDArrayExcludingTeamMembers(ctx, req.(*GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUser_GetUserUUIDByContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserUUIDByContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserServer).GetUserUUIDByContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.customer.CustomerUser/GetUserUUIDByContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserServer).GetUserUUIDByContact(ctx, req.(*GetUserUUIDByContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUser_GetContactsByUserAndTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsByUserAndTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserServer).GetContactsByUserAndTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.customer.CustomerUser/GetContactsByUserAndTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserServer).GetContactsByUserAndTeams(ctx, req.(*GetContactsByUserAndTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerUser_ServiceDesc is the grpc.ServiceDesc for CustomerUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.customer.CustomerUser",
	HandlerType: (*CustomerUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsersInfoByUserUUIDArray",
			Handler:    _CustomerUser_GetUsersInfoByUserUUIDArray_Handler,
		},
		{
			MethodName: "GetUsersInfoByUserUUIDArrayExcludingTeamMembers",
			Handler:    _CustomerUser_GetUsersInfoByUserUUIDArrayExcludingTeamMembers_Handler,
		},
		{
			MethodName: "GetUserUUIDByContact",
			Handler:    _CustomerUser_GetUserUUIDByContact_Handler,
		},
		{
			MethodName: "GetContactsByUserAndTeams",
			Handler:    _CustomerUser_GetContactsByUserAndTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/customer/user_service.proto",
}
