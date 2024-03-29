// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/gateway/admin/billing/personal_account/responsible_person/v1/responsible_person_service.proto

package responsible_person

import (
	context "context"
	v1 "github.com/deniskelin/tada-proto/tada/billing/api/personal_account/responsible_person/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResponsiblePersonServiceClient is the client API for ResponsiblePersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResponsiblePersonServiceClient interface {
	Create(ctx context.Context, in *v1.CreateRequest, opts ...grpc.CallOption) (*v1.CreateResponse, error)
	Update(ctx context.Context, in *v1.UpdateRequest, opts ...grpc.CallOption) (*v1.UpdateResponse, error)
	Get(ctx context.Context, in *v1.GetRequest, opts ...grpc.CallOption) (*v1.GetListResponse, error)
	GetList(ctx context.Context, in *v1.GetListRequest, opts ...grpc.CallOption) (*v1.GetListResponse, error)
	GetCountsList(ctx context.Context, in *v1.GetListRequest, opts ...grpc.CallOption) (*v1.GetCountsListResponse, error)
	Delete(ctx context.Context, in *v1.DeleteRequest, opts ...grpc.CallOption) (*v1.DeleteResponse, error)
}

type responsiblePersonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResponsiblePersonServiceClient(cc grpc.ClientConnInterface) ResponsiblePersonServiceClient {
	return &responsiblePersonServiceClient{cc}
}

func (c *responsiblePersonServiceClient) Create(ctx context.Context, in *v1.CreateRequest, opts ...grpc.CallOption) (*v1.CreateResponse, error) {
	out := new(v1.CreateResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responsiblePersonServiceClient) Update(ctx context.Context, in *v1.UpdateRequest, opts ...grpc.CallOption) (*v1.UpdateResponse, error) {
	out := new(v1.UpdateResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responsiblePersonServiceClient) Get(ctx context.Context, in *v1.GetRequest, opts ...grpc.CallOption) (*v1.GetListResponse, error) {
	out := new(v1.GetListResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responsiblePersonServiceClient) GetList(ctx context.Context, in *v1.GetListRequest, opts ...grpc.CallOption) (*v1.GetListResponse, error) {
	out := new(v1.GetListResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responsiblePersonServiceClient) GetCountsList(ctx context.Context, in *v1.GetListRequest, opts ...grpc.CallOption) (*v1.GetCountsListResponse, error) {
	out := new(v1.GetCountsListResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/GetCountsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *responsiblePersonServiceClient) Delete(ctx context.Context, in *v1.DeleteRequest, opts ...grpc.CallOption) (*v1.DeleteResponse, error) {
	out := new(v1.DeleteResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResponsiblePersonServiceServer is the server API for ResponsiblePersonService service.
// All implementations must embed UnimplementedResponsiblePersonServiceServer
// for forward compatibility
type ResponsiblePersonServiceServer interface {
	Create(context.Context, *v1.CreateRequest) (*v1.CreateResponse, error)
	Update(context.Context, *v1.UpdateRequest) (*v1.UpdateResponse, error)
	Get(context.Context, *v1.GetRequest) (*v1.GetListResponse, error)
	GetList(context.Context, *v1.GetListRequest) (*v1.GetListResponse, error)
	GetCountsList(context.Context, *v1.GetListRequest) (*v1.GetCountsListResponse, error)
	Delete(context.Context, *v1.DeleteRequest) (*v1.DeleteResponse, error)
	mustEmbedUnimplementedResponsiblePersonServiceServer()
}

// UnimplementedResponsiblePersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResponsiblePersonServiceServer struct {
}

func (UnimplementedResponsiblePersonServiceServer) Create(context.Context, *v1.CreateRequest) (*v1.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) Update(context.Context, *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) Get(context.Context, *v1.GetRequest) (*v1.GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) GetList(context.Context, *v1.GetListRequest) (*v1.GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) GetCountsList(context.Context, *v1.GetListRequest) (*v1.GetCountsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountsList not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) Delete(context.Context, *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResponsiblePersonServiceServer) mustEmbedUnimplementedResponsiblePersonServiceServer() {
}

// UnsafeResponsiblePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResponsiblePersonServiceServer will
// result in compilation errors.
type UnsafeResponsiblePersonServiceServer interface {
	mustEmbedUnimplementedResponsiblePersonServiceServer()
}

func RegisterResponsiblePersonServiceServer(s grpc.ServiceRegistrar, srv ResponsiblePersonServiceServer) {
	s.RegisterService(&ResponsiblePersonService_ServiceDesc, srv)
}

func _ResponsiblePersonService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).Create(ctx, req.(*v1.CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResponsiblePersonService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).Update(ctx, req.(*v1.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResponsiblePersonService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).Get(ctx, req.(*v1.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResponsiblePersonService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).GetList(ctx, req.(*v1.GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResponsiblePersonService_GetCountsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).GetCountsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/GetCountsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).GetCountsList(ctx, req.(*v1.GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResponsiblePersonService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponsiblePersonServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponsiblePersonServiceServer).Delete(ctx, req.(*v1.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResponsiblePersonService_ServiceDesc is the grpc.ServiceDesc for ResponsiblePersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResponsiblePersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.gateway.admin.billing.personal_account.responsible_person.v1.ResponsiblePersonService",
	HandlerType: (*ResponsiblePersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResponsiblePersonService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResponsiblePersonService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResponsiblePersonService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ResponsiblePersonService_GetList_Handler,
		},
		{
			MethodName: "GetCountsList",
			Handler:    _ResponsiblePersonService_GetCountsList_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResponsiblePersonService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/gateway/admin/billing/personal_account/responsible_person/v1/responsible_person_service.proto",
}
