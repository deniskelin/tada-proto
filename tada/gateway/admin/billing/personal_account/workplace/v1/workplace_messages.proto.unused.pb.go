// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/gateway/admin/billing/personal_account/workplace/v1/workplace_messages.proto.unused

package workplace

import (
	v1 "github.com/deniskelin/tada-proto/tada/user/api/v1"
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

type AddUserInWorkplaceByJidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetByContactRequest *v1.GetByContactRequest `protobuf:"bytes,1,opt,name=get_by_contact_request,json=getByContactRequest,proto3" json:"get_by_contact_request,omitempty"`
	PersonalAccountId   int64                   `protobuf:"varint,2,opt,name=personal_account_id,json=personalAccountId,proto3" json:"personal_account_id,omitempty"`
}

func (x *AddUserInWorkplaceByJidRequest) Reset() {
	*x = AddUserInWorkplaceByJidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserInWorkplaceByJidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserInWorkplaceByJidRequest) ProtoMessage() {}

func (x *AddUserInWorkplaceByJidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserInWorkplaceByJidRequest.ProtoReflect.Descriptor instead.
func (*AddUserInWorkplaceByJidRequest) Descriptor() ([]byte, []int) {
	return file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescGZIP(), []int{0}
}

func (x *AddUserInWorkplaceByJidRequest) GetGetByContactRequest() *v1.GetByContactRequest {
	if x != nil {
		return x.GetByContactRequest
	}
	return nil
}

func (x *AddUserInWorkplaceByJidRequest) GetPersonalAccountId() int64 {
	if x != nil {
		return x.PersonalAccountId
	}
	return 0
}

var File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused protoreflect.FileDescriptor

var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDesc = []byte{
	0x0a, 0x58, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x12, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x74, 0x61, 0x64, 0x61,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x01, 0x0a, 0x1e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x4a, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x13, 0x67, 0x65, 0x74, 0x42,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42,
	0xf0, 0x01, 0x0a, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0xaa, 0x02, 0x2d, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x2d, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescOnce sync.Once
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescData = file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDesc
)

func file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescGZIP() []byte {
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescOnce.Do(func() {
		file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescData = protoimpl.X.CompressGZIP(file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescData)
	})
	return file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDescData
}

var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_goTypes = []interface{}{
	(*AddUserInWorkplaceByJidRequest)(nil), // 0: tada.billing.api.personal_account.workplace.v1.AddUserInWorkplaceByJidRequest
	(*v1.GetByContactRequest)(nil),         // 1: tada.user.api.v1.GetByContactRequest
}
var file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_depIdxs = []int32{
	1, // 0: tada.billing.api.personal_account.workplace.v1.AddUserInWorkplaceByJidRequest.get_by_contact_request:type_name -> tada.user.api.v1.GetByContactRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_init()
}
func file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_init() {
	if File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserInWorkplaceByJidRequest); i {
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
			RawDescriptor: file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_goTypes,
		DependencyIndexes: file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_depIdxs,
		MessageInfos:      file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_msgTypes,
	}.Build()
	File_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused = out.File
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_rawDesc = nil
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_goTypes = nil
	file_tada_gateway_admin_billing_personal_account_workplace_v1_workplace_messages_proto_unused_depIdxs = nil
}
