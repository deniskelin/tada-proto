// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/gateway/admin/meeting/calendar_service.proto

package meetingGatewayPb

import (
	context "context"
	meeting "gitlab.tada.team/tdapis/go-genproto/tada/api/meeting"
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
	GetAnotherMeetings(ctx context.Context, in *meeting.GetAnotherMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error)
	GetMeetingsInstances(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error)
	GetMeetingsDates(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsDatesResponse, error)
	GetMeetingsCounts(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsCountsResponse, error)
	GetMeetingsGroupsByInterval(ctx context.Context, in *meeting.GetMeetingsGroupsByIntervalRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsGroupsByIntervalResponse, error)
	GetMeetingById(ctx context.Context, in *wrapperspb.UInt64Value, opts ...grpc.CallOption) (*meeting.Meeting, error)
	GetMeetingByGroupUUID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*meeting.Meeting, error)
	CreateMeeting(ctx context.Context, in *meeting.CreateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error)
	UpdateMeeting(ctx context.Context, in *meeting.UpdateMeetingRequest, opts ...grpc.CallOption) (*meeting.Meeting, error)
	DeleteMeeting(ctx context.Context, in *meeting.DeleteMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	GetMeetingMembers(ctx context.Context, in *meeting.GetMeetingMembersRequest, opts ...grpc.CallOption) (*meeting.GetMeetingMembersResponse, error)
	AddMemberInMeeting(ctx context.Context, in *meeting.AddMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error)
	UpdateMemberInMeeting(ctx context.Context, in *meeting.UpdateMemberInMeetingRequest, opts ...grpc.CallOption) (*meeting.Member, error)
	DeleteMemberFromMeeting(ctx context.Context, in *DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	UpdateMeetingCell(ctx context.Context, in *meeting.UpdateMeetingCellRequest, opts ...grpc.CallOption) (*meeting.Meeting, error)
	DeleteMeetingCell(ctx context.Context, in *meeting.DeleteMeetingCellRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	GenerateMeetingCells(ctx context.Context, in *meeting.GenerateMeetingCellsRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error)
	ConvertFrequencyToString(ctx context.Context, in *meeting.FrequencyToStringRequest, opts ...grpc.CallOption) (*meeting.FrequencyToStringResponse, error)
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

func (c *meetingGatewayClient) GetAnotherMeetings(ctx context.Context, in *meeting.GetAnotherMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error) {
	out := new(meeting.GetMeetingsResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetAnotherMeetings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GetMeetingsInstances(ctx context.Context, in *meeting.GetMeetingsRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsResponse, error) {
	out := new(meeting.GetMeetingsResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsInstances", in, out, opts...)
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

func (c *meetingGatewayClient) GetMeetingsGroupsByInterval(ctx context.Context, in *meeting.GetMeetingsGroupsByIntervalRequest, opts ...grpc.CallOption) (*meeting.GetMeetingsGroupsByIntervalResponse, error) {
	out := new(meeting.GetMeetingsGroupsByIntervalResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsGroupsByInterval", in, out, opts...)
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

func (c *meetingGatewayClient) GetMeetingByGroupUUID(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingByGroupUUID", in, out, opts...)
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

func (c *meetingGatewayClient) DeleteMemberFromMeeting(ctx context.Context, in *DeleteMemberFromMeetingRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/DeleteMemberFromMeeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) UpdateMeetingCell(ctx context.Context, in *meeting.UpdateMeetingCellRequest, opts ...grpc.CallOption) (*meeting.Meeting, error) {
	out := new(meeting.Meeting)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/UpdateMeetingCell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) DeleteMeetingCell(ctx context.Context, in *meeting.DeleteMeetingCellRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/DeleteMeetingCell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) GenerateMeetingCells(ctx context.Context, in *meeting.GenerateMeetingCellsRequest, opts ...grpc.CallOption) (*meeting.SuccessResponse, error) {
	out := new(meeting.SuccessResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/GenerateMeetingCells", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingGatewayClient) ConvertFrequencyToString(ctx context.Context, in *meeting.FrequencyToStringRequest, opts ...grpc.CallOption) (*meeting.FrequencyToStringResponse, error) {
	out := new(meeting.FrequencyToStringResponse)
	err := c.cc.Invoke(ctx, "/tada.gateway.admin.meeting.MeetingGateway/ConvertFrequencyToString", in, out, opts...)
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
	GetAnotherMeetings(context.Context, *meeting.GetAnotherMeetingsRequest) (*meeting.GetMeetingsResponse, error)
	GetMeetingsInstances(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsResponse, error)
	GetMeetingsDates(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsDatesResponse, error)
	GetMeetingsCounts(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsCountsResponse, error)
	GetMeetingsGroupsByInterval(context.Context, *meeting.GetMeetingsGroupsByIntervalRequest) (*meeting.GetMeetingsGroupsByIntervalResponse, error)
	GetMeetingById(context.Context, *wrapperspb.UInt64Value) (*meeting.Meeting, error)
	GetMeetingByGroupUUID(context.Context, *wrapperspb.StringValue) (*meeting.Meeting, error)
	CreateMeeting(context.Context, *meeting.CreateMeetingRequest) (*meeting.Meeting, error)
	UpdateMeeting(context.Context, *meeting.UpdateMeetingRequest) (*meeting.Meeting, error)
	DeleteMeeting(context.Context, *meeting.DeleteMeetingRequest) (*meeting.SuccessResponse, error)
	GetMeetingMembers(context.Context, *meeting.GetMeetingMembersRequest) (*meeting.GetMeetingMembersResponse, error)
	AddMemberInMeeting(context.Context, *meeting.AddMemberInMeetingRequest) (*meeting.Member, error)
	UpdateMemberInMeeting(context.Context, *meeting.UpdateMemberInMeetingRequest) (*meeting.Member, error)
	DeleteMemberFromMeeting(context.Context, *DeleteMemberFromMeetingRequest) (*meeting.SuccessResponse, error)
	UpdateMeetingCell(context.Context, *meeting.UpdateMeetingCellRequest) (*meeting.Meeting, error)
	DeleteMeetingCell(context.Context, *meeting.DeleteMeetingCellRequest) (*meeting.SuccessResponse, error)
	GenerateMeetingCells(context.Context, *meeting.GenerateMeetingCellsRequest) (*meeting.SuccessResponse, error)
	ConvertFrequencyToString(context.Context, *meeting.FrequencyToStringRequest) (*meeting.FrequencyToStringResponse, error)
	mustEmbedUnimplementedMeetingGatewayServer()
}

// UnimplementedMeetingGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingGatewayServer struct {
}

func (UnimplementedMeetingGatewayServer) GetMeetings(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetings not implemented")
}
func (UnimplementedMeetingGatewayServer) GetAnotherMeetings(context.Context, *meeting.GetAnotherMeetingsRequest) (*meeting.GetMeetingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnotherMeetings not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsInstances(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsInstances not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsDates(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsDates not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsCounts(context.Context, *meeting.GetMeetingsRequest) (*meeting.GetMeetingsCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsCounts not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingsGroupsByInterval(context.Context, *meeting.GetMeetingsGroupsByIntervalRequest) (*meeting.GetMeetingsGroupsByIntervalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingsGroupsByInterval not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingById(context.Context, *wrapperspb.UInt64Value) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingById not implemented")
}
func (UnimplementedMeetingGatewayServer) GetMeetingByGroupUUID(context.Context, *wrapperspb.StringValue) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeetingByGroupUUID not implemented")
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
func (UnimplementedMeetingGatewayServer) DeleteMemberFromMeeting(context.Context, *DeleteMemberFromMeetingRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemberFromMeeting not implemented")
}
func (UnimplementedMeetingGatewayServer) UpdateMeetingCell(context.Context, *meeting.UpdateMeetingCellRequest) (*meeting.Meeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeetingCell not implemented")
}
func (UnimplementedMeetingGatewayServer) DeleteMeetingCell(context.Context, *meeting.DeleteMeetingCellRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeetingCell not implemented")
}
func (UnimplementedMeetingGatewayServer) GenerateMeetingCells(context.Context, *meeting.GenerateMeetingCellsRequest) (*meeting.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMeetingCells not implemented")
}
func (UnimplementedMeetingGatewayServer) ConvertFrequencyToString(context.Context, *meeting.FrequencyToStringRequest) (*meeting.FrequencyToStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertFrequencyToString not implemented")
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

func _MeetingGateway_GetAnotherMeetings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetAnotherMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetAnotherMeetings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetAnotherMeetings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetAnotherMeetings(ctx, req.(*meeting.GetAnotherMeetingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GetMeetingsInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingsInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingsInstances(ctx, req.(*meeting.GetMeetingsRequest))
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

func _MeetingGateway_GetMeetingsGroupsByInterval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GetMeetingsGroupsByIntervalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingsGroupsByInterval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingsGroupsByInterval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingsGroupsByInterval(ctx, req.(*meeting.GetMeetingsGroupsByIntervalRequest))
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

func _MeetingGateway_GetMeetingByGroupUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GetMeetingByGroupUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GetMeetingByGroupUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GetMeetingByGroupUUID(ctx, req.(*wrapperspb.StringValue))
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
	in := new(DeleteMemberFromMeetingRequest)
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
		return srv.(MeetingGatewayServer).DeleteMemberFromMeeting(ctx, req.(*DeleteMemberFromMeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_UpdateMeetingCell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.UpdateMeetingCellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).UpdateMeetingCell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/UpdateMeetingCell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).UpdateMeetingCell(ctx, req.(*meeting.UpdateMeetingCellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_DeleteMeetingCell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.DeleteMeetingCellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).DeleteMeetingCell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/DeleteMeetingCell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).DeleteMeetingCell(ctx, req.(*meeting.DeleteMeetingCellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_GenerateMeetingCells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.GenerateMeetingCellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).GenerateMeetingCells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/GenerateMeetingCells",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).GenerateMeetingCells(ctx, req.(*meeting.GenerateMeetingCellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingGateway_ConvertFrequencyToString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(meeting.FrequencyToStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingGatewayServer).ConvertFrequencyToString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.gateway.admin.meeting.MeetingGateway/ConvertFrequencyToString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingGatewayServer).ConvertFrequencyToString(ctx, req.(*meeting.FrequencyToStringRequest))
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
			MethodName: "GetAnotherMeetings",
			Handler:    _MeetingGateway_GetAnotherMeetings_Handler,
		},
		{
			MethodName: "GetMeetingsInstances",
			Handler:    _MeetingGateway_GetMeetingsInstances_Handler,
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
			MethodName: "GetMeetingsGroupsByInterval",
			Handler:    _MeetingGateway_GetMeetingsGroupsByInterval_Handler,
		},
		{
			MethodName: "GetMeetingById",
			Handler:    _MeetingGateway_GetMeetingById_Handler,
		},
		{
			MethodName: "GetMeetingByGroupUUID",
			Handler:    _MeetingGateway_GetMeetingByGroupUUID_Handler,
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
			MethodName: "UpdateMeetingCell",
			Handler:    _MeetingGateway_UpdateMeetingCell_Handler,
		},
		{
			MethodName: "DeleteMeetingCell",
			Handler:    _MeetingGateway_DeleteMeetingCell_Handler,
		},
		{
			MethodName: "GenerateMeetingCells",
			Handler:    _MeetingGateway_GenerateMeetingCells_Handler,
		},
		{
			MethodName: "ConvertFrequencyToString",
			Handler:    _MeetingGateway_ConvertFrequencyToString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/gateway/admin/meeting/calendar_service.proto",
}
