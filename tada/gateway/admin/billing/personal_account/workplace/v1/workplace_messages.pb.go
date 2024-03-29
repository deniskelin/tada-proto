// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/gateway/admin/billing/personal_account/workplace/v1/workplace_messages.proto

package workplace

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

type AddUserByContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonalAccountId uint64 `protobuf:"varint,1,opt,name=personal_account_id,json=personalAccountId,proto3" json:"personal_account_id,omitempty"`
	ContactUuid       string `protobuf:"bytes,2,opt,name=contact_uuid,json=contactUuid,proto3" json:"contact_uuid,omitempty"`
}

func (x *AddUserByContactRequest) Reset() {
	*x = AddUserByContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserByContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserByContactRequest) ProtoMessage() {}

func (x *AddUserByContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserByContactRequest.ProtoReflect.Descriptor instead.
func (*AddUserByContactRequest) Descriptor() ([]byte, []int) {
	return file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserByContactRequest) GetPersonalAccountId() uint64 {
	if x != nil {
		return x.PersonalAccountId
	}
	return 0
}

func (x *AddUserByContactRequest) GetContactUuid() string {
	if x != nil {
		return x.ContactUuid
	}
	return ""
}

type AddUserByContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddUserByContactResponse) Reset() {
	*x = AddUserByContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserByContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserByContactResponse) ProtoMessage() {}

func (x *AddUserByContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserByContactResponse.ProtoReflect.Descriptor instead.
func (*AddUserByContactResponse) Descriptor() ([]byte, []int) {
	return file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescGZIP(), []int{1}
}

var File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto protoreflect.FileDescriptor

var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDesc = []byte{
	0x0a, 0x51, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x38, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x6c, 0x0a,
	0x17, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x98, 0x02, 0x0a, 0x38, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0xaa, 0x02,
	0x37, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x37, 0x54, 0x61, 0x64, 0x61, 0x5c,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescOnce sync.Once
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescData = file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDesc
)

func file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescGZIP() []byte {
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescOnce.Do(func() {
		file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescData)
	})
	return file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDescData
}

var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_goTypes = []interface{}{
	(*AddUserByContactRequest)(nil),  // 0: tada.gateway.admin.billing.personal_account.workplace.v1.AddUserByContactRequest
	(*AddUserByContactResponse)(nil), // 1: tada.gateway.admin.billing.personal_account.workplace.v1.AddUserByContactResponse
}
var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_init()
}
func file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_init() {
	if File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserByContactRequest); i {
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
		file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserByContactResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_goTypes,
		DependencyIndexes: file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_depIdxs,
		MessageInfos:      file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_msgTypes,
	}.Build()
	File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto = out.File
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_rawDesc = nil
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_goTypes = nil
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_depIdxs = nil
}
