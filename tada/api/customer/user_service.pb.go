// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: tada/api/customer/user_service.proto

package customer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FullName     *string                `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Phone        *string                `protobuf:"bytes,3,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Email        *string                `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	LastActivity *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_activity,json=lastActivity,proto3,oneof" json:"last_activity,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserInfo) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UserInfo) GetLastActivity() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActivity
	}
	return nil
}

// UsersInfo users information
type UsersInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo []*UserInfo `protobuf:"bytes,1,rep,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *UsersInfo) Reset() {
	*x = UsersInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersInfo) ProtoMessage() {}

func (x *UsersInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersInfo.ProtoReflect.Descriptor instead.
func (*UsersInfo) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UsersInfo) GetUserInfo() []*UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// GetUsersInfoByUserUUIDArrayRequest request on get user information by array of UUID's users
type GetUsersInfoByUserUUIDArrayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid []string `protobuf:"bytes,1,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Limit    *int32   `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset   *int32   `protobuf:"varint,3,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetUsersInfoByUserUUIDArrayRequest) Reset() {
	*x = GetUsersInfoByUserUUIDArrayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersInfoByUserUUIDArrayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersInfoByUserUUIDArrayRequest) ProtoMessage() {}

func (x *GetUsersInfoByUserUUIDArrayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersInfoByUserUUIDArrayRequest.ProtoReflect.Descriptor instead.
func (*GetUsersInfoByUserUUIDArrayRequest) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersInfoByUserUUIDArrayRequest) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *GetUsersInfoByUserUUIDArrayRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetUsersInfoByUserUUIDArrayRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

// GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest request on get user information by array of UUID's users excluding team members in uuid team
type GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid []string `protobuf:"bytes,1,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TeamUuid string   `protobuf:"bytes,2,opt,name=team_uuid,json=teamUuid,proto3" json:"team_uuid,omitempty"`
	Limit    *int32   `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset   *int32   `protobuf:"varint,4,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) Reset() {
	*x = GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) ProtoMessage() {}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) GetTeamUuid() string {
	if x != nil {
		return x.TeamUuid
	}
	return ""
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

// GetUserUUIDByContactRequest request on get user UUIDs by array of UUID's contacts
type GetUserUUIDByContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactUuid []string `protobuf:"bytes,1,rep,name=contact_uuid,json=contactUuid,proto3" json:"contact_uuid,omitempty"`
	Limit       *int32   `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset      *int32   `protobuf:"varint,3,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetUserUUIDByContactRequest) Reset() {
	*x = GetUserUUIDByContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserUUIDByContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserUUIDByContactRequest) ProtoMessage() {}

func (x *GetUserUUIDByContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserUUIDByContactRequest.ProtoReflect.Descriptor instead.
func (*GetUserUUIDByContactRequest) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserUUIDByContactRequest) GetContactUuid() []string {
	if x != nil {
		return x.ContactUuid
	}
	return nil
}

func (x *GetUserUUIDByContactRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetUserUUIDByContactRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

// GetUserUUIDByContactResponse response on get user UUIDs by array of UUID's contacts
type GetUserUUIDByContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid []string `protobuf:"bytes,1,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetUserUUIDByContactResponse) Reset() {
	*x = GetUserUUIDByContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserUUIDByContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserUUIDByContactResponse) ProtoMessage() {}

func (x *GetUserUUIDByContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserUUIDByContactResponse.ProtoReflect.Descriptor instead.
func (*GetUserUUIDByContactResponse) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserUUIDByContactResponse) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

// GetContactsByUserAndTeamsRequest request on get contacts uuid (jid) by user UUID and teams
type GetContactsByUserAndTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string   `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TeamUuid []string `protobuf:"bytes,2,rep,name=team_uuid,json=teamUuid,proto3" json:"team_uuid,omitempty"`
	Limit    *int32   `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset   *int32   `protobuf:"varint,4,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetContactsByUserAndTeamsRequest) Reset() {
	*x = GetContactsByUserAndTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactsByUserAndTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsByUserAndTeamsRequest) ProtoMessage() {}

func (x *GetContactsByUserAndTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsByUserAndTeamsRequest.ProtoReflect.Descriptor instead.
func (*GetContactsByUserAndTeamsRequest) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetContactsByUserAndTeamsRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetContactsByUserAndTeamsRequest) GetTeamUuid() []string {
	if x != nil {
		return x.TeamUuid
	}
	return nil
}

func (x *GetContactsByUserAndTeamsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetContactsByUserAndTeamsRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

// GetContactsByUserAndTeamsResponse request on get contacts uuid (jid) by user UUID and teams
type GetContactsByUserAndTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactUuid []string `protobuf:"bytes,1,rep,name=contact_uuid,json=contactUuid,proto3" json:"contact_uuid,omitempty"`
}

func (x *GetContactsByUserAndTeamsResponse) Reset() {
	*x = GetContactsByUserAndTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_api_customer_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactsByUserAndTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsByUserAndTeamsResponse) ProtoMessage() {}

func (x *GetContactsByUserAndTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_api_customer_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsByUserAndTeamsResponse.ProtoReflect.Descriptor instead.
func (*GetContactsByUserAndTeamsResponse) Descriptor() ([]byte, []int) {
	return file_tada_api_customer_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetContactsByUserAndTeamsResponse) GetContactUuid() []string {
	if x != nil {
		return x.ContactUuid
	}
	return nil
}

var File_tada_api_customer_user_service_proto protoreflect.FileDescriptor

var file_tada_api_customer_user_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x44, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x36, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x61, 0x6d,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x55, 0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75,
	0x69, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x46,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x55, 0x75, 0x69, 0x64, 0x32, 0xc6, 0x06, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0xb8, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x35, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x44, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3e, 0x22, 0x39, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41, 0x72, 0x72, 0x61, 0x79, 0x3a,
	0x01, 0x2a, 0x12, 0xf4, 0x01, 0x0a, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x61, 0x6d, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x49, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49,
	0x44, 0x41, 0x72, 0x72, 0x61, 0x79, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x52, 0x22, 0x4d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x61, 0x6d, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xb6, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x2e, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x22, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0xca, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x33, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3c, 0x22, 0x37, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x42,
	0x93, 0x01, 0x0a, 0x11, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x14, 0x54, 0x61, 0x64, 0x61, 0x41, 0x50, 0x49, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f,
	0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0xaa, 0x02, 0x11,
	0x54, 0x61, 0x64, 0x61, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0xca, 0x02, 0x11, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x41, 0x50, 0x49, 0x5c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_api_customer_user_service_proto_rawDescOnce sync.Once
	file_tada_api_customer_user_service_proto_rawDescData = file_tada_api_customer_user_service_proto_rawDesc
)

func file_tada_api_customer_user_service_proto_rawDescGZIP() []byte {
	file_tada_api_customer_user_service_proto_rawDescOnce.Do(func() {
		file_tada_api_customer_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_api_customer_user_service_proto_rawDescData)
	})
	return file_tada_api_customer_user_service_proto_rawDescData
}

var file_tada_api_customer_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tada_api_customer_user_service_proto_goTypes = []interface{}{
	(*UserInfo)(nil),  // 0: tada.api.customer.UserInfo
	(*UsersInfo)(nil), // 1: tada.api.customer.UsersInfo
	(*GetUsersInfoByUserUUIDArrayRequest)(nil),                     // 2: tada.api.customer.GetUsersInfoByUserUUIDArrayRequest
	(*GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest)(nil), // 3: tada.api.customer.GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest
	(*GetUserUUIDByContactRequest)(nil),                            // 4: tada.api.customer.GetUserUUIDByContactRequest
	(*GetUserUUIDByContactResponse)(nil),                           // 5: tada.api.customer.GetUserUUIDByContactResponse
	(*GetContactsByUserAndTeamsRequest)(nil),                       // 6: tada.api.customer.GetContactsByUserAndTeamsRequest
	(*GetContactsByUserAndTeamsResponse)(nil),                      // 7: tada.api.customer.GetContactsByUserAndTeamsResponse
	(*timestamppb.Timestamp)(nil),                                  // 8: google.protobuf.Timestamp
}
var file_tada_api_customer_user_service_proto_depIdxs = []int32{
	8, // 0: tada.api.customer.UserInfo.last_activity:type_name -> google.protobuf.Timestamp
	0, // 1: tada.api.customer.UsersInfo.user_info:type_name -> tada.api.customer.UserInfo
	2, // 2: tada.api.customer.CustomerUser.GetUsersInfoByUserUUIDArray:input_type -> tada.api.customer.GetUsersInfoByUserUUIDArrayRequest
	3, // 3: tada.api.customer.CustomerUser.GetUsersInfoByUserUUIDArrayExcludingTeamMembers:input_type -> tada.api.customer.GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest
	4, // 4: tada.api.customer.CustomerUser.GetUserUUIDByContact:input_type -> tada.api.customer.GetUserUUIDByContactRequest
	6, // 5: tada.api.customer.CustomerUser.GetContactsByUserAndTeams:input_type -> tada.api.customer.GetContactsByUserAndTeamsRequest
	1, // 6: tada.api.customer.CustomerUser.GetUsersInfoByUserUUIDArray:output_type -> tada.api.customer.UsersInfo
	1, // 7: tada.api.customer.CustomerUser.GetUsersInfoByUserUUIDArrayExcludingTeamMembers:output_type -> tada.api.customer.UsersInfo
	5, // 8: tada.api.customer.CustomerUser.GetUserUUIDByContact:output_type -> tada.api.customer.GetUserUUIDByContactResponse
	7, // 9: tada.api.customer.CustomerUser.GetContactsByUserAndTeams:output_type -> tada.api.customer.GetContactsByUserAndTeamsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tada_api_customer_user_service_proto_init() }
func file_tada_api_customer_user_service_proto_init() {
	if File_tada_api_customer_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tada_api_customer_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersInfoByUserUUIDArrayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserUUIDByContactRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserUUIDByContactResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactsByUserAndTeamsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tada_api_customer_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactsByUserAndTeamsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tada_api_customer_user_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_tada_api_customer_user_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_tada_api_customer_user_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_tada_api_customer_user_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_tada_api_customer_user_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_api_customer_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tada_api_customer_user_service_proto_goTypes,
		DependencyIndexes: file_tada_api_customer_user_service_proto_depIdxs,
		MessageInfos:      file_tada_api_customer_user_service_proto_msgTypes,
	}.Build()
	File_tada_api_customer_user_service_proto = out.File
	file_tada_api_customer_user_service_proto_rawDesc = nil
	file_tada_api_customer_user_service_proto_goTypes = nil
	file_tada_api_customer_user_service_proto_depIdxs = nil
}
