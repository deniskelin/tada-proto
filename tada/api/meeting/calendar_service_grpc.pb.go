// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: tada/api/meeting/calendar_service.proto

package meetingApiPb

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

// MeetingApiClient is the client API for MeetingApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingApiClient interface {
	GetMeetings(ctx context.Context, in *GetMeetingsRequest, opts ...grpc.CallOption) (*GetMeetingsResponse, error)
	GetMeetingById(ctx context.Context, in *GetMeetingByIdRequest, opts ...grpc.CallOption) (*Meeting, error)
	GetMeetingByChatUUID(ctx context.Context, in *GetMeetingByChatUUIDRequest, opts ...grpc.CallOption) (*Meeting, error)
	CreateMeeting(ctx context.Context, in *CreateMeetingRequest, opts ...grpc.CallOption) (*Meeting, error)
	UpdateMeeting(ctx context.Context, in *UpdateMeetingRequest, opts ...grpc.CallOption) (*Meeting, error)
	DeleteMeeting(ctx context.Context, in *DeleteMeetingRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	GetMeetingMembers(ctx context.Context, in *GetMeetingMembersRequest, opts ...grpc.CallOption) (*GetMeetingMembersResponse, error)
	AddMemberInMeeting(ctx context.Context, in *AddMemberInMeetingRequest, opts ...grpc.CallOption) (*Member, error)
	UpdateMemberInMeeting(ctx context.Context, in *UpdateMemberInMeetingRequest, opts ...grpc.CallOption) (*Member, error)
	DeleteMemberFromMeeting(ctx context.Context, in *DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type meetingApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingApiClient(cc grpc.ClientConnInterface) MeetingApiClient {
	return &meetingApiClient{cc}
}

func (c *meetingApiClient) GetMeetings(ctx context.Context, in *GetMeetingsRequest, opts ...grpc.CallOption) (*GetMeetingsResponse, error) {
	out := new(GetMeetingsResponse)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/GetMeetings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) GetMeetingById(ctx context.Context, in *GetMeetingByIdRequest, opts ...grpc.CallOption) (*Meeting, error) {
	out := new(Meeting)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/GetMeetingById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) GetMeetingByChatUUID(ctx context.Context, in *GetMeetingByChatUUIDRequest, opts ...grpc.CallOption) (*Meeting, error) {
	out := new(Meeting)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/GetMeetingByChatUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) CreateMeeting(ctx context.Context, in *CreateMeetingRequest, opts ...grpc.CallOption) (*Meeting, error) {
	out := new(Meeting)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/CreateMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) UpdateMeeting(ctx context.Context, in *UpdateMeetingRequest, opts ...grpc.CallOption) (*Meeting, error) {
	out := new(Meeting)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/UpdateMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) DeleteMeeting(ctx context.Context, in *DeleteMeetingRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/DeleteMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) GetMeetingMembers(ctx context.Context, in *GetMeetingMembersRequest, opts ...grpc.CallOption) (*GetMeetingMembersResponse, error) {
	out := new(GetMeetingMembersResponse)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/GetMeetingMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) AddMemberInMeeting(ctx context.Context, in *AddMemberInMeetingRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/AddMemberInMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) UpdateMemberInMeeting(ctx context.Context, in *UpdateMemberInMeetingRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/UpdateMemberInMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingApiClient) DeleteMemberFromMeeting(ctx context.Context, in *DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.api.meeting.MeetingApi/DeleteMemberFromMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingApiServer is the server API for MeetingApi service.
// All implementations must embed UnimplementedMeetingApiServer
// for forward compatibility
type MeetingApiServer interface {
	GetMeetings(context.Context, *GetMeetingsRequest) (*GetMeetingsResponse, error)
	GetMeetingById(context.Context, *GetMeetingByIdRequest) (*Meeting, error)
	GetMeetingByChatUUID(context.Context, *GetMeetingByChatUUIDRequest) (*Meeting, error)
	CreateMeeting(context.Context, *CreateMeetingRequest) (*Meeting, error)
	UpdateMeeting(context.Context, *UpdateMeetingRequest) (*Meeting, error)
	DeleteMeeting(context.Context, *DeleteMeetingRequest) (*SuccessResponse, error)
	GetMeetingMembers(context.Context, *GetMeetingMembersRequest) (*GetMeetingMembersResponse, error)
	AddMemberInMeeting(context.Context, *AddMemberInMeetingRequest) (*Member, error)
	UpdateMemberInMeeting(context.Context, *UpdateMemberInMeetingRequest) (*Member, error)
	DeleteMemberFromMeeting(context.Context, *DeleteMemberFromMeetingRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedMeetingApiServer()
}

// UnimplementedMeetingApiServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingApiServer struct {
}

func (UnimplementedMeetingApiServer) GetMeetings(context.Context, *GetMeetingsRequest) (*GetMeetingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetings not implemented")
}
func (UnimplementedMeetingApiServer) GetMeetingById(context.Context, *GetMeetingByIdRequest) (*Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingById not implemented")
}
func (UnimplementedMeetingApiServer) GetMeetingByChatUUID(context.Context, *GetMeetingByChatUUIDRequest) (*Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingByChatUUID not implemented")
}
func (UnimplementedMeetingApiServer) CreateMeeting(context.Context, *CreateMeetingRequest) (*Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeeting not implemented")
}
func (UnimplementedMeetingApiServer) UpdateMeeting(context.Context, *UpdateMeetingRequest) (*Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeeting not implemented")
}
func (UnimplementedMeetingApiServer) DeleteMeeting(context.Context, *DeleteMeetingRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeeting not implemented")
}
func (UnimplementedMeetingApiServer) GetMeetingMembers(context.Context, *GetMeetingMembersRequest) (*GetMeetingMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingMembers not implemented")
}
func (UnimplementedMeetingApiServer) AddMemberInMeeting(context.Context, *AddMemberInMeetingRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberInMeeting not implemented")
}
func (UnimplementedMeetingApiServer) UpdateMemberInMeeting(context.Context, *UpdateMemberInMeetingRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemberInMeeting not implemented")
}
func (UnimplementedMeetingApiServer) DeleteMemberFromMeeting(context.Context, *DeleteMemberFromMeetingRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemberFromMeeting not implemented")
}
func (UnimplementedMeetingApiServer) mustEmbedUnimplementedMeetingApiServer() {}

// UnsafeMeetingApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingApiServer will
// result in compilation errors.
type UnsafeMeetingApiServer interface {
	mustEmbedUnimplementedMeetingApiServer()
}

func RegisterMeetingApiServer(s grpc.ServiceRegistrar, srv MeetingApiServer) {
	s.RegisterService(&MeetingApi_ServiceDesc, srv)
}

func _MeetingApi_GetMeetings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).GetMeetings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/GetMeetings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).GetMeetings(ctx, req.(*GetMeetingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_GetMeetingById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).GetMeetingById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/GetMeetingById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).GetMeetingById(ctx, req.(*GetMeetingByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_GetMeetingByChatUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingByChatUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).GetMeetingByChatUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/GetMeetingByChatUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).GetMeetingByChatUUID(ctx, req.(*GetMeetingByChatUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_CreateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).CreateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/CreateMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).CreateMeeting(ctx, req.(*CreateMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_UpdateMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).UpdateMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/UpdateMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).UpdateMeeting(ctx, req.(*UpdateMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_DeleteMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).DeleteMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/DeleteMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).DeleteMeeting(ctx, req.(*DeleteMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_GetMeetingMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeetingMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).GetMeetingMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/GetMeetingMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).GetMeetingMembers(ctx, req.(*GetMeetingMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_AddMemberInMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberInMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).AddMemberInMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/AddMemberInMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).AddMemberInMeeting(ctx, req.(*AddMemberInMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_UpdateMemberInMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemberInMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).UpdateMemberInMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/UpdateMemberInMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).UpdateMemberInMeeting(ctx, req.(*UpdateMemberInMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingApi_DeleteMemberFromMeeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberFromMeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingApiServer).DeleteMemberFromMeeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.meeting.MeetingApi/DeleteMemberFromMeeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingApiServer).DeleteMemberFromMeeting(ctx, req.(*DeleteMemberFromMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeetingApi_ServiceDesc is the grpc.ServiceDesc for MeetingApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeetingApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.meeting.MeetingApi",
	HandlerType: (*MeetingApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeetings",
			Handler:    _MeetingApi_GetMeetings_Handler,
		},
		{
			MethodName: "GetMeetingById",
			Handler:    _MeetingApi_GetMeetingById_Handler,
		},
		{
			MethodName: "GetMeetingByChatUUID",
			Handler:    _MeetingApi_GetMeetingByChatUUID_Handler,
		},
		{
			MethodName: "CreateMeeting",
			Handler:    _MeetingApi_CreateMeeting_Handler,
		},
		{
			MethodName: "UpdateMeeting",
			Handler:    _MeetingApi_UpdateMeeting_Handler,
		},
		{
			MethodName: "DeleteMeeting",
			Handler:    _MeetingApi_DeleteMeeting_Handler,
		},
		{
			MethodName: "GetMeetingMembers",
			Handler:    _MeetingApi_GetMeetingMembers_Handler,
		},
		{
			MethodName: "AddMemberInMeeting",
			Handler:    _MeetingApi_AddMemberInMeeting_Handler,
		},
		{
			MethodName: "UpdateMemberInMeeting",
			Handler:    _MeetingApi_UpdateMemberInMeeting_Handler,
		},
		{
			MethodName: "DeleteMemberFromMeeting",
			Handler:    _MeetingApi_DeleteMemberFromMeeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/meeting/calendar_service.proto",
}
