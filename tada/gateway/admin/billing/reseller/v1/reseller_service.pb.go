// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/gateway/admin/billing/reseller/v1/reseller_service.proto

package reseller

import (
	v1 "github.com/deniskelin/tada-proto/tada/billing/api/reseller/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tada_gateway_admin_billing_reseller_v1_reseller_service_proto protoreflect.FileDescriptor

var file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x26, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x05, 0x0a, 0x0f,
	0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x84, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x22, 0x11, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x2b, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0xd1, 0x01, 0x0a, 0x26, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0xaa, 0x02, 0x26, 0x54, 0x61,
	0x64, 0x61, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5c, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_goTypes = []interface{}{
	(*v1.CreateRequest)(nil),   // 0: tada.billing.api.reseller.v1.CreateRequest
	(*v1.UpdateRequest)(nil),   // 1: tada.billing.api.reseller.v1.UpdateRequest
	(*v1.GetRequest)(nil),      // 2: tada.billing.api.reseller.v1.GetRequest
	(*v1.GetListRequest)(nil),  // 3: tada.billing.api.reseller.v1.GetListRequest
	(*v1.DeleteRequest)(nil),   // 4: tada.billing.api.reseller.v1.DeleteRequest
	(*v1.CreateResponse)(nil),  // 5: tada.billing.api.reseller.v1.CreateResponse
	(*v1.UpdateResponse)(nil),  // 6: tada.billing.api.reseller.v1.UpdateResponse
	(*v1.GetListResponse)(nil), // 7: tada.billing.api.reseller.v1.GetListResponse
	(*v1.DeleteResponse)(nil),  // 8: tada.billing.api.reseller.v1.DeleteResponse
}
var file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_depIdxs = []int32{
	0, // 0: tada.gateway.admin.billing.reseller.v1.ResellerService.Create:input_type -> tada.billing.api.reseller.v1.CreateRequest
	1, // 1: tada.gateway.admin.billing.reseller.v1.ResellerService.Update:input_type -> tada.billing.api.reseller.v1.UpdateRequest
	2, // 2: tada.gateway.admin.billing.reseller.v1.ResellerService.Get:input_type -> tada.billing.api.reseller.v1.GetRequest
	3, // 3: tada.gateway.admin.billing.reseller.v1.ResellerService.GetList:input_type -> tada.billing.api.reseller.v1.GetListRequest
	4, // 4: tada.gateway.admin.billing.reseller.v1.ResellerService.Delete:input_type -> tada.billing.api.reseller.v1.DeleteRequest
	5, // 5: tada.gateway.admin.billing.reseller.v1.ResellerService.Create:output_type -> tada.billing.api.reseller.v1.CreateResponse
	6, // 6: tada.gateway.admin.billing.reseller.v1.ResellerService.Update:output_type -> tada.billing.api.reseller.v1.UpdateResponse
	7, // 7: tada.gateway.admin.billing.reseller.v1.ResellerService.Get:output_type -> tada.billing.api.reseller.v1.GetListResponse
	7, // 8: tada.gateway.admin.billing.reseller.v1.ResellerService.GetList:output_type -> tada.billing.api.reseller.v1.GetListResponse
	8, // 9: tada.gateway.admin.billing.reseller.v1.ResellerService.Delete:output_type -> tada.billing.api.reseller.v1.DeleteResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_init() }
func file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_init() {
	if File_tada_gateway_admin_billing_reseller_v1_reseller_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_goTypes,
		DependencyIndexes: file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_depIdxs,
	}.Build()
	File_tada_gateway_admin_billing_reseller_v1_reseller_service_proto = out.File
	file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_rawDesc = nil
	file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_goTypes = nil
	file_tada_gateway_admin_billing_reseller_v1_reseller_service_proto_depIdxs = nil
}
