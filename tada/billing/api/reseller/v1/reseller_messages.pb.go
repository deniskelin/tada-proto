// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/reseller/v1/reseller_messages.proto

package reseller

import (
	common "github.com/deniskelin/tada-proto/tada/common"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reseller *Element `protobuf:"bytes,1,opt,name=reseller,proto3" json:"reseller,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetReseller() *Element {
	if x != nil {
		return x.Reseller
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[2]
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
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  *uint32 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset *uint64 `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetListRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetListRequest) GetOffset() uint64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResellerList []*Element `protobuf:"bytes,1,rep,name=reseller_list,json=resellerList,proto3" json:"reseller_list,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetListResponse) GetResellerList() []*Element {
	if x != nil {
		return x.ResellerList
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Code *string `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reseller *Element `protobuf:"bytes,1,opt,name=reseller,proto3" json:"reseller,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetReseller() *Element {
	if x != nil {
		return x.Reseller
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResellerIds *common.ArrayId `protobuf:"bytes,1,opt,name=reseller_ids,json=resellerIds,proto3" json:"reseller_ids,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetResellerIds() *common.ArrayId {
	if x != nil {
		return x.ResellerIds
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP(), []int{8}
}

var File_tada_billing_api_reseller_v1_reseller_messages_proto protoreflect.FileDescriptor

var file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDesc = []byte{
	0x0a, 0x34, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x22, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x5d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x63, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x49, 0x64, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa9, 0x01, 0x0a, 0x1c, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0xaa, 0x02, 0x1c, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescOnce sync.Once
	file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescData = file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDesc
)

func file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescGZIP() []byte {
	file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescData)
	})
	return file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDescData
}

var file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tada_billing_api_reseller_v1_reseller_messages_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: tada.billing.api.reseller.v1.CreateRequest
	(*CreateResponse)(nil),  // 1: tada.billing.api.reseller.v1.CreateResponse
	(*GetRequest)(nil),      // 2: tada.billing.api.reseller.v1.GetRequest
	(*GetListRequest)(nil),  // 3: tada.billing.api.reseller.v1.GetListRequest
	(*GetListResponse)(nil), // 4: tada.billing.api.reseller.v1.GetListResponse
	(*UpdateRequest)(nil),   // 5: tada.billing.api.reseller.v1.UpdateRequest
	(*UpdateResponse)(nil),  // 6: tada.billing.api.reseller.v1.UpdateResponse
	(*DeleteRequest)(nil),   // 7: tada.billing.api.reseller.v1.DeleteRequest
	(*DeleteResponse)(nil),  // 8: tada.billing.api.reseller.v1.DeleteResponse
	(*Element)(nil),         // 9: tada.billing.api.reseller.v1.Element
	(*Identity)(nil),        // 10: tada.billing.api.reseller.v1.Identity
	(*common.ArrayId)(nil),  // 11: tada.common.ArrayId
}
var file_tada_billing_api_reseller_v1_reseller_messages_proto_depIdxs = []int32{
	9,  // 0: tada.billing.api.reseller.v1.CreateResponse.reseller:type_name -> tada.billing.api.reseller.v1.Element
	10, // 1: tada.billing.api.reseller.v1.GetRequest.identity:type_name -> tada.billing.api.reseller.v1.Identity
	9,  // 2: tada.billing.api.reseller.v1.GetListResponse.reseller_list:type_name -> tada.billing.api.reseller.v1.Element
	9,  // 3: tada.billing.api.reseller.v1.UpdateResponse.reseller:type_name -> tada.billing.api.reseller.v1.Element
	11, // 4: tada.billing.api.reseller.v1.DeleteRequest.reseller_ids:type_name -> tada.common.ArrayId
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_tada_billing_api_reseller_v1_reseller_messages_proto_init() }
func file_tada_billing_api_reseller_v1_reseller_messages_proto_init() {
	if File_tada_billing_api_reseller_v1_reseller_messages_proto != nil {
		return
	}
	file_tada_billing_api_reseller_v1_reseller_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
	file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_reseller_v1_reseller_messages_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_reseller_v1_reseller_messages_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_reseller_v1_reseller_messages_proto_msgTypes,
	}.Build()
	File_tada_billing_api_reseller_v1_reseller_messages_proto = out.File
	file_tada_billing_api_reseller_v1_reseller_messages_proto_rawDesc = nil
	file_tada_billing_api_reseller_v1_reseller_messages_proto_goTypes = nil
	file_tada_billing_api_reseller_v1_reseller_messages_proto_depIdxs = nil
}
