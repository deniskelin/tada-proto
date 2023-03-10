// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/tariff/v1/tariff_types.proto

package tariff

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

// Status of tariff
type TariffStatus int32

const (
	TariffStatus_TARIFF_STATUS_UNSPECIFIED TariffStatus = 0
	TariffStatus_TARIFF_STATUS_ACTIVE      TariffStatus = 1
	TariffStatus_TARIFF_STATUS_ARCHIVE     TariffStatus = 2
)

// Enum value maps for TariffStatus.
var (
	TariffStatus_name = map[int32]string{
		0: "TARIFF_STATUS_UNSPECIFIED",
		1: "TARIFF_STATUS_ACTIVE",
		2: "TARIFF_STATUS_ARCHIVE",
	}
	TariffStatus_value = map[string]int32{
		"TARIFF_STATUS_UNSPECIFIED": 0,
		"TARIFF_STATUS_ACTIVE":      1,
		"TARIFF_STATUS_ARCHIVE":     2,
	}
)

func (x TariffStatus) Enum() *TariffStatus {
	p := new(TariffStatus)
	*p = x
	return p
}

func (x TariffStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TariffStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tada_billing_api_tariff_v1_tariff_types_proto_enumTypes[0].Descriptor()
}

func (TariffStatus) Type() protoreflect.EnumType {
	return &file_tada_billing_api_tariff_v1_tariff_types_proto_enumTypes[0]
}

func (x TariffStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TariffStatus.Descriptor instead.
func (TariffStatus) EnumDescriptor() ([]byte, []int) {
	return file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescGZIP(), []int{0}
}

var File_tada_billing_api_tariff_v1_tariff_types_proto protoreflect.FileDescriptor

var file_tada_billing_api_tariff_v1_tariff_types_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2a, 0x62, 0x0a, 0x0c, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x54,
	0x41, 0x52, 0x49, 0x46, 0x46, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41,
	0x52, 0x49, 0x46, 0x46, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x52, 0x49, 0x46, 0x46, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x10, 0x02, 0x42,
	0x9f, 0x01, 0x0a, 0x1a, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65,
	0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0xaa, 0x02, 0x1a, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5c, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescOnce sync.Once
	file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescData = file_tada_billing_api_tariff_v1_tariff_types_proto_rawDesc
)

func file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescGZIP() []byte {
	file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescData)
	})
	return file_tada_billing_api_tariff_v1_tariff_types_proto_rawDescData
}

var file_tada_billing_api_tariff_v1_tariff_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tada_billing_api_tariff_v1_tariff_types_proto_goTypes = []interface{}{
	(TariffStatus)(0), // 0: tada.billing.api.tariff.v1.TariffStatus
}
var file_tada_billing_api_tariff_v1_tariff_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tada_billing_api_tariff_v1_tariff_types_proto_init() }
func file_tada_billing_api_tariff_v1_tariff_types_proto_init() {
	if File_tada_billing_api_tariff_v1_tariff_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_tariff_v1_tariff_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_tariff_v1_tariff_types_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_tariff_v1_tariff_types_proto_depIdxs,
		EnumInfos:         file_tada_billing_api_tariff_v1_tariff_types_proto_enumTypes,
	}.Build()
	File_tada_billing_api_tariff_v1_tariff_types_proto = out.File
	file_tada_billing_api_tariff_v1_tariff_types_proto_rawDesc = nil
	file_tada_billing_api_tariff_v1_tariff_types_proto_goTypes = nil
	file_tada_billing_api_tariff_v1_tariff_types_proto_depIdxs = nil
}
