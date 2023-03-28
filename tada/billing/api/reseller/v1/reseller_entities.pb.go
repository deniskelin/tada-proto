// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/reseller/v1/reseller_entities.proto

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

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identity:
	//	*Identity_ResellerIds
	//	*Identity_Code
	Identity isIdentity_Identity `protobuf_oneof:"identity"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescGZIP(), []int{0}
}

func (m *Identity) GetIdentity() isIdentity_Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (x *Identity) GetResellerIds() *common.ArrayId {
	if x, ok := x.GetIdentity().(*Identity_ResellerIds); ok {
		return x.ResellerIds
	}
	return nil
}

func (x *Identity) GetCode() string {
	if x, ok := x.GetIdentity().(*Identity_Code); ok {
		return x.Code
	}
	return ""
}

type isIdentity_Identity interface {
	isIdentity_Identity()
}

type Identity_ResellerIds struct {
	ResellerIds *common.ArrayId `protobuf:"bytes,1,opt,name=reseller_ids,json=resellerIds,proto3,oneof"`
}

type Identity_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*Identity_ResellerIds) isIdentity_Identity() {}

func (*Identity_Code) isIdentity_Identity() {}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Element) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Element) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Element) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_tada_billing_api_reseller_v1_reseller_entities_proto protoreflect.FileDescriptor

var file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDesc = []byte{
	0x0a, 0x34, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67,
	0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x49, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x06,
	0x42, 0xa9, 0x01, 0x0a, 0x1c, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0xaa, 0x02,
	0x1c, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c,
	0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescOnce sync.Once
	file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescData = file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDesc
)

func file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescGZIP() []byte {
	file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescData)
	})
	return file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDescData
}

var file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tada_billing_api_reseller_v1_reseller_entities_proto_goTypes = []interface{}{
	(*Identity)(nil),       // 0: tada.billing.api.reseller.v1.Identity
	(*Element)(nil),        // 1: tada.billing.api.reseller.v1.Element
	(*common.ArrayId)(nil), // 2: tada.common.ArrayId
}
var file_tada_billing_api_reseller_v1_reseller_entities_proto_depIdxs = []int32{
	2, // 0: tada.billing.api.reseller.v1.Identity.reseller_ids:type_name -> tada.common.ArrayId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tada_billing_api_reseller_v1_reseller_entities_proto_init() }
func file_tada_billing_api_reseller_v1_reseller_entities_proto_init() {
	if File_tada_billing_api_reseller_v1_reseller_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
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
	file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Identity_ResellerIds)(nil),
		(*Identity_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_reseller_v1_reseller_entities_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_reseller_v1_reseller_entities_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_reseller_v1_reseller_entities_proto_msgTypes,
	}.Build()
	File_tada_billing_api_reseller_v1_reseller_entities_proto = out.File
	file_tada_billing_api_reseller_v1_reseller_entities_proto_rawDesc = nil
	file_tada_billing_api_reseller_v1_reseller_entities_proto_goTypes = nil
	file_tada_billing_api_reseller_v1_reseller_entities_proto_depIdxs = nil
}
