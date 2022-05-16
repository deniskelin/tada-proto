// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: tada/gateway/admin/meeting/calendar_service.proto

package meetingGatewayPb

import (
	context "context"
	meeting "github.com/deniskelin/tada-proto/tada/api/meeting"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeetingGatewayClient is the client API for MeetingGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingGatewayClient interface {
	GetMeetings(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error)
	GetMeetingsDates(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsDatesResponse, error)
	GetMeetingsCounts(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsCountsResponse, error)
	GetMeetingById(ctx context.Context, in *wrapperspb.UInt64Value, opts ...grpc.CallOption) (*meeting.Meeting, error)
	GetMeetingByChatUUID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*meeting.Meeting, error)
	CreateMeeting(ctx context.Context, in *meeting.CreateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error)
	UpdateMeeting(ctx context.Context, in *meeting.UpdateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error)
	DeleteMeeting(ctx context.Context, in *meeting.DeleteMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	GetMeetingMembers(ctx context.Context, in *meeting.GetMeetingMembersRequest, opts ...grpc.CallOption) (*meeting.GetMeetingMembersResponse, error)
	AddMemberInMeeting(ctx context.Context, in *meeting.AddMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error)
	UpdateMemberInMeeting(ctx context.Context, in *meeting.UpdateMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error)
	DeleteMemberFromMeeting(ctx context.Context, in *meeting.DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	GenerateSells(ctx context.Context, in *meeting.GenerateSellsRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
}

type meetingGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingGatewayClient(cc grpc.ClientConnInterface) MeetingGatewayClient {
	return &meetingGatewayClient{cc}
}

func (c *meetingGatewayClient) GetMeetings(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error) {
	out := new(meeting.GetMeetingsResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingsDates(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsDatesResponse, error) {
	out := new(meeting.GetMeetingsDatesResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingsCounts(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsCountsResponse, error) {
	out := new(meeting.GetMeetingsCountsResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingById(ctx context.Context, in *wrapperspb.UInt64Value, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingByChatUUID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingByChatUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) CreateMeeting(ctx context.Context, in *meeting.CreateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/CreateMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) UpdateMeeting(ctx context.Context, in *meeting.UpdateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/UpdateMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) DeleteMeeting(ctx context.Context, in *meeting.DeleteMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/DeleteMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingMembers(ctx context.Context, in *meeting.GetMeetingMembersRequest, opts ...grpc.CallOption) (*meeting.GetMeetingMembersResponse, error) {
	out := new(meeting.GetMeetingMembersResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) AddMemberInMeeting(ctx context.Context, in *meeting.AddMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error) {
	out := new(meeting.Member)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/AddMemberInMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) UpdateMemberInMeeting(ctx context.Context, in *meeting.UpdateMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error) {
	out := new(meeting.Member)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/UpdateMemberInMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) DeleteMemberFromMeeting(ctx context.Context, in *meeting.DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/DeleteMemberFromMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GenerateSells(ctx context.Context, in *meeting.GenerateSellsRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GenerateSells", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingGatewayServer is the server API for MeetingGateway service.
// All implementations must embed UnimplementedMeetingGatewayServer
// for forward compatibility
type MeetingGatewayServer interface {
	GetMeetings(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsResponse, error)
	GetMeetingsDates(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsDatesResponse, error)
	GetMeetingsCounts(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsCountsResponse, error)
	GetMeetingById(context.Context, *wrapperspb.UInt64Value) (*meeting.Meeting, error)
	GetMeetingByChatUUID(context.Context, *wrapperspb.StringValue) (*meeting.Meeting, error)
	CreateMeeting(context.Context, *meeting.CreateMeetingRequest) (*meeting.Meeting, error)
	UpdateMeeting(context.Context, *meeting.UpdateMeetingRequest) (*meeting.Meeting, error)
	DeleteMeeting(context.Context, *meeting.DeleteMeetingRequest) (*meeting.SuccessResponse, error)
	GetMeetingMembers(context.Context, *meeting.GetMeetingMembersRequest) (*meeting.GetMeetingMembersResponse, error)
	AddMemberInMeeting(context.Context, *meeting.AddMemberInMeetingRequest) (*meeting.Member, error)
	UpdateMemberInMeeting(context.Context, *meeting.UpdateMemberInMeetingRequest) (*meeting.Member, error)
	DeleteMemberFromMeeting(context.Context, *meeting.DeleteMemberFromMeetingRequest) (*meeting.SuccessResponse, error)
	GenerateSells(context.Context, *meeting.GenerateSellsRequest) (*meeting.SuccessResponse, error)
	mustEmbedUnimplementedMeetingGatewayServer()
}

// UnimplementedMeetingGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingGatewayServer struct {
}

func (UnimplementedMeetingGatewayServer) GetMeetings(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetings not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsDates(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsDates not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsCounts(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsCounts not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingById(context.Context, *wrapperspb.UInt64Value) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingById not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingByChatUUID(context.Context, *wrapperspb.StringValue) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingByChatUUID not implemented")
}
func (UnimplementedMeetingGatewayServer) CreateMeeting(context.Context, *meeting.CreateMeetingRequest) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) UpdateMeeting(context.Context, *meeting.UpdateMeetingRequest) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) DeleteMeeting(context.Context, *meeting.DeleteMeetingRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingMembers(context.Context, *meeting.GetMeetingMembersRequest) (*meeting.GetMeetingMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingMembers not implemented")
}
func (UnimplementedMeetingGatewayServer) AddMemberInMeeting(context.Context, *meeting.AddMemberInMeetingRequest) (*meeting.Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberInMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) UpdateMemberInMeeting(context.Context, *meeting.UpdateMemberInMeetingRequest) (*meeting.Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemberInMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) DeleteMemberFromMeeting(context.Context, *meeting.DeleteMemberFromMeetingRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemberFromMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) GenerateSells(context.Context, *meeting.GenerateSellsRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSells not implemented")
}
func (UnimplementedMeetingGatewayServer) mustEmbedUnimplementedMeetingGatewayServer() {}

// UnsafeMeetingGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingGatewayServer will
// result in compilation errors.
type UnsafeMeetingGatewayServer interface {
	mustEmbedUnimplementedMeetingGatewayServer()
}

func RegisterMeetingGatewayServer(s grpc.ServiceRegistrar, srv MeetingGatewayServer) {
	s.RegisterService(&MeetingGateway_ServiceDesc, srv)
}

func _MeetingGateway_GetMeetings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetings(ctx, req.(*meeting.GetMeetingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingsDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingsDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingsDates(ctx, req.(*meeting.GetMeetingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingsCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingsCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingsCounts(ctx, req.(*meeting.GetMeetingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.UInt64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingById(ctx, req.(*wrapperspb.UInt64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingByChatUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingByChatUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingByChatUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingByChatUUID(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_CreateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.CreateMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).CreateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/CreateMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).CreateMeeting(ctx, req.(*meeting.CreateMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_UpdateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.UpdateMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).UpdateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/UpdateMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).UpdateMeeting(ctx, req.(*meeting.UpdateMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_DeleteMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.DeleteMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).DeleteMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/DeleteMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).DeleteMeeting(ctx, req.(*meeting.DeleteMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingMembers(ctx, req.(*meeting.GetMeetingMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_AddMemberInMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.AddMemberInMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).AddMemberInMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/AddMemberInMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).AddMemberInMeeting(ctx, req.(*meeting.AddMemberInMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_UpdateMemberInMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.UpdateMemberInMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).UpdateMemberInMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/UpdateMemberInMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).UpdateMemberInMeeting(ctx, req.(*meeting.UpdateMemberInMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_DeleteMemberFromMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.DeleteMemberFromMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).DeleteMemberFromMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/DeleteMemberFromMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).DeleteMemberFromMeeting(ctx, req.(*meeting.DeleteMemberFromMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GenerateSells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GenerateSellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GenerateSells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GenerateSells",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GenerateSells(ctx, req.(*meeting.GenerateSellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeetingGateway_ServiceDesc is the grpc.ServiceDesc for MeetingGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeetingGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.gateway.admin.meeting.MeetingGateway",
	HandlerType: (*MeetingGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeetings",
			Handler:    _MeetingGateway_GetMeetings_Handler,
		},
		{
			MethodName: "GetMeetingsDates",
			Handler:    _MeetingGateway_GetMeetingsDates_Handler,
		},
		{
			MethodName: "GetMeetingsCounts",
			Handler:    _MeetingGateway_GetMeetingsCounts_Handler,
		},
		{
			MethodName: "GetMeetingById",
			Handler:    _MeetingGateway_GetMeetingById_Handler,
		},
		{
			MethodName: "GetMeetingByChatUUID",
			Handler:    _MeetingGateway_GetMeetingByChatUUID_Handler,
		},
		{
			MethodName: "CreateMeeting",
			Handler:    _MeetingGateway_CreateMeeting_Handler,
		},
		{
			MethodName: "UpdateMeeting",
			Handler:    _MeetingGateway_UpdateMeeting_Handler,
		},
		{
			MethodName: "DeleteMeeting",
			Handler:    _MeetingGateway_DeleteMeeting_Handler,
		},
		{
			MethodName: "GetMeetingMembers",
			Handler:    _MeetingGateway_GetMeetingMembers_Handler,
		},
		{
			MethodName: "AddMemberInMeeting",
			Handler:    _MeetingGateway_AddMemberInMeeting_Handler,
		},
		{
			MethodName: "UpdateMemberInMeeting",
			Handler:    _MeetingGateway_UpdateMemberInMeeting_Handler,
		},
		{
			MethodName: "DeleteMemberFromMeeting",
			Handler:    _MeetingGateway_DeleteMemberFromMeeting_Handler,
		},
		{
			MethodName: "GenerateSells",
			Handler:    _MeetingGateway_GenerateSells_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/gateway/admin/meeting/calendar_service.proto",
}
