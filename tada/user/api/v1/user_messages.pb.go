// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/user/api/v1/user_messages.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid          []string `protobuf:"bytes,1,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	ExcludingTeamUuid *string  `protobuf:"bytes,2,opt,name=excluding_team_uuid,json=excludingTeamUuid,proto3,oneof" json:"excluding_team_uuid,omitempty"`
	Limit             *uint32  `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset            *uint64  `protobuf:"varint,4,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *GetRequest) GetExcludingTeamUuid() string {
	if x != nil && x.ExcludingTeamUuid != nil {
		return *x.ExcludingTeamUuid
	}
	return ""
}

func (x *GetRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetRequest) GetOffset() uint64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*Element `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetUserList() []*Element {
	if x != nil {
		return x.UserList
	}
	return nil
}

type GetByContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactUuid []string `protobuf:"bytes,1,rep,name=contact_uuid,json=contactUuid,proto3" json:"contact_uuid,omitempty"`
}

func (x *GetByContactRequest) Reset() {
	*x = GetByContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByContactRequest) ProtoMessage() {}

func (x *GetByContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByContactRequest.ProtoReflect.Descriptor instead.
func (*GetByContactRequest) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetByContactRequest) GetContactUuid() []string {
	if x != nil {
		return x.ContactUuid
	}
	return nil
}

type GetByContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid []string `protobuf:"bytes,1,rep,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetByContactResponse) Reset() {
	*x = GetByContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByContactResponse) ProtoMessage() {}

func (x *GetByContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByContactResponse.ProtoReflect.Descriptor instead.
func (*GetByContactResponse) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetByContactResponse) GetUserUuid() []string {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

type GetContactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string   `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TeamUuid []string `protobuf:"bytes,2,rep,name=team_uuid,json=teamUuid,proto3" json:"team_uuid,omitempty"`
	Limit    *uint32  `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset   *uint64  `protobuf:"varint,4,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetContactsRequest) Reset() {
	*x = GetContactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsRequest) ProtoMessage() {}

func (x *GetContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsRequest.ProtoReflect.Descriptor instead.
func (*GetContactsRequest) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetContactsRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetContactsRequest) GetTeamUuid() []string {
	if x != nil {
		return x.TeamUuid
	}
	return nil
}

func (x *GetContactsRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetContactsRequest) GetOffset() uint64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetContactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactUuid []string `protobuf:"bytes,1,rep,name=contact_uuid,json=contactUuid,proto3" json:"contact_uuid,omitempty"`
}

func (x *GetContactsResponse) Reset() {
	*x = GetContactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsResponse) ProtoMessage() {}

func (x *GetContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_user_api_v1_user_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsResponse.ProtoReflect.Descriptor instead.
func (*GetContactsResponse) Descriptor() ([]byte, []int) {
	return file_tada_user_api_v1_user_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetContactsResponse) GetContactUuid() []string {
	if x != nil {
		return x.ContactUuid
	}
	return nil
}

var File_tada_user_api_v1_user_messages_proto protoreflect.FileDescriptor

var file_tada_user_api_v1_user_messages_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x13, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x45, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x42, 0x74, 0x0a, 0x10, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61,
	0x64, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0xaa, 0x02, 0x10, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x55, 0x73,
	0x65, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_user_api_v1_user_messages_proto_rawDescOnce sync.Once
	file_tada_user_api_v1_user_messages_proto_rawDescData = file_tada_user_api_v1_user_messages_proto_rawDesc
)

func file_tada_user_api_v1_user_messages_proto_rawDescGZIP() []byte {
	file_tada_user_api_v1_user_messages_proto_rawDescOnce.Do(func() {
		file_tada_user_api_v1_user_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_user_api_v1_user_messages_proto_rawDescData)
	})
	return file_tada_user_api_v1_user_messages_proto_rawDescData
}

var file_tada_user_api_v1_user_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tada_user_api_v1_user_messages_proto_goTypes = []interface{}{
	(*GetRequest)(nil),           // 0: tada.user.api.v1.GetRequest
	(*GetResponse)(nil),          // 1: tada.user.api.v1.GetResponse
	(*GetByContactRequest)(nil),  // 2: tada.user.api.v1.GetByContactRequest
	(*GetByContactResponse)(nil), // 3: tada.user.api.v1.GetByContactResponse
	(*GetContactsRequest)(nil),   // 4: tada.user.api.v1.GetContactsRequest
	(*GetContactsResponse)(nil),  // 5: tada.user.api.v1.GetContactsResponse
	(*Element)(nil),              // 6: tada.user.api.v1.Element
}
var file_tada_user_api_v1_user_messages_proto_depIdxs = []int32{
	6, // 0: tada.user.api.v1.GetResponse.user_list:type_name -> tada.user.api.v1.Element
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tada_user_api_v1_user_messages_proto_init() }
func file_tada_user_api_v1_user_messages_proto_init() {
	if File_tada_user_api_v1_user_messages_proto != nil {
		return
	}
	file_tada_user_api_v1_user_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_user_api_v1_user_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_tada_user_api_v1_user_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_tada_user_api_v1_user_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByContactRequest); i {
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
		file_tada_user_api_v1_user_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByContactResponse); i {
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
		file_tada_user_api_v1_user_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactsRequest); i {
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
		file_tada_user_api_v1_user_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactsResponse); i {
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
	file_tada_user_api_v1_user_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_tada_user_api_v1_user_messages_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_user_api_v1_user_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_user_api_v1_user_messages_proto_goTypes,
		DependencyIndexes: file_tada_user_api_v1_user_messages_proto_depIdxs,
		MessageInfos:      file_tada_user_api_v1_user_messages_proto_msgTypes,
	}.Build()
	File_tada_user_api_v1_user_messages_proto = out.File
	file_tada_user_api_v1_user_messages_proto_rawDesc = nil
	file_tada_user_api_v1_user_messages_proto_goTypes = nil
	file_tada_user_api_v1_user_messages_proto_depIdxs = nil
}
