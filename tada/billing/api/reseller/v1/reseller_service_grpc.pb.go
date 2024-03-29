// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/billing/api/reseller/v1/reseller_service.proto

package reseller

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

// ResellerServiceClient is the client API for ResellerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResellerServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetCountsList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetCountsListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type resellerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResellerServiceClient(cc grpc.ClientConnInterface) ResellerServiceClient {
	return &resellerServiceClient{cc}
}

func (c *resellerServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resellerServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resellerServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resellerServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resellerServiceClient) GetCountsList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetCountsListResponse, error) {
	out := new(GetCountsListResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/GetCountsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resellerServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/tada.billing.api.reseller.v1.ResellerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResellerServiceServer is the server API for ResellerService service.
// All implementations must embed UnimplementedResellerServiceServer
// for forward compatibility
type ResellerServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Get(context.Context, *GetRequest) (*GetListResponse, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetCountsList(context.Context, *GetListRequest) (*GetCountsListResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedResellerServiceServer()
}

// UnimplementedResellerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResellerServiceServer struct {
}

func (UnimplementedResellerServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResellerServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResellerServiceServer) Get(context.Context, *GetRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResellerServiceServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedResellerServiceServer) GetCountsList(context.Context, *GetListRequest) (*GetCountsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountsList not implemented")
}
func (UnimplementedResellerServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResellerServiceServer) mustEmbedUnimplementedResellerServiceServer() {}

// UnsafeResellerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResellerServiceServer will
// result in compilation errors.
type UnsafeResellerServiceServer interface {
	mustEmbedUnimplementedResellerServiceServer()
}

func RegisterResellerServiceServer(s grpc.ServiceRegistrar, srv ResellerServiceServer) {
	s.RegisterService(&ResellerService_ServiceDesc, srv)
}

func _ResellerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResellerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResellerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResellerService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResellerService_GetCountsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).GetCountsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/GetCountsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).GetCountsList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResellerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResellerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.billing.api.reseller.v1.ResellerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResellerServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResellerService_ServiceDesc is the grpc.ServiceDesc for ResellerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResellerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.billing.api.reseller.v1.ResellerService",
	HandlerType: (*ResellerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResellerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResellerService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResellerService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ResellerService_GetList_Handler,
		},
		{
			MethodName: "GetCountsList",
			Handler:    _ResellerService_GetCountsList_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResellerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/billing/api/reseller/v1/reseller_service.proto",
}
