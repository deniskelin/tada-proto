// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/gateway/admin/billing/personal_account/enquiry/v1/enquiry_service.proto

package enquiry

import (
	v1 "github.com/deniskelin/tada-proto/tada/billing/api/personal_account/enquiry/v1"
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

var File_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto protoreflect.FileDescriptor

var file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8e, 0x0b, 0x0a, 0x0e, 0x45, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x65, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22,
	0x12, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x9a, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x38,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22,
	0x0f, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0xa6, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xb8, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x65, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x3b, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xba, 0x01, 0x0a,
	0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x2e,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x42, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x65,
	0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xce, 0x01, 0x0a, 0x11, 0x53, 0x65,
	0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x46, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x8e, 0x02, 0x0a, 0x36, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x71, 0x75, 0x69,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64,
	0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0xaa, 0x02, 0x35, 0x54, 0x61,
	0x64, 0x61, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x35, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5c, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_goTypes = []interface{}{
	(*v1.CreateRequest)(nil),             // 0: tada.billing.api.personal_account.enquiry.v1.CreateRequest
	(*v1.UpdateRequest)(nil),             // 1: tada.billing.api.personal_account.enquiry.v1.UpdateRequest
	(*v1.GetRequest)(nil),                // 2: tada.billing.api.personal_account.enquiry.v1.GetRequest
	(*v1.GetListRequest)(nil),            // 3: tada.billing.api.personal_account.enquiry.v1.GetListRequest
	(*v1.DeleteRequest)(nil),             // 4: tada.billing.api.personal_account.enquiry.v1.DeleteRequest
	(*v1.ChangeStatusRequest)(nil),       // 5: tada.billing.api.personal_account.enquiry.v1.ChangeStatusRequest
	(*v1.SetLifecycleDatesRequest)(nil),  // 6: tada.billing.api.personal_account.enquiry.v1.SetLifecycleDatesRequest
	(*v1.CreateResponse)(nil),            // 7: tada.billing.api.personal_account.enquiry.v1.CreateResponse
	(*v1.UpdateResponse)(nil),            // 8: tada.billing.api.personal_account.enquiry.v1.UpdateResponse
	(*v1.GetListResponse)(nil),           // 9: tada.billing.api.personal_account.enquiry.v1.GetListResponse
	(*v1.GetCountsListResponse)(nil),     // 10: tada.billing.api.personal_account.enquiry.v1.GetCountsListResponse
	(*v1.DeleteResponse)(nil),            // 11: tada.billing.api.personal_account.enquiry.v1.DeleteResponse
	(*v1.ChangeStatusResponse)(nil),      // 12: tada.billing.api.personal_account.enquiry.v1.ChangeStatusResponse
	(*v1.SetLifecycleDatesResponse)(nil), // 13: tada.billing.api.personal_account.enquiry.v1.SetLifecycleDatesResponse
}
var file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_depIdxs = []int32{
	0,  // 0: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Create:input_type -> tada.billing.api.personal_account.enquiry.v1.CreateRequest
	1,  // 1: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Update:input_type -> tada.billing.api.personal_account.enquiry.v1.UpdateRequest
	2,  // 2: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Get:input_type -> tada.billing.api.personal_account.enquiry.v1.GetRequest
	3,  // 3: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.GetList:input_type -> tada.billing.api.personal_account.enquiry.v1.GetListRequest
	3,  // 4: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.GetCountsList:input_type -> tada.billing.api.personal_account.enquiry.v1.GetListRequest
	4,  // 5: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Delete:input_type -> tada.billing.api.personal_account.enquiry.v1.DeleteRequest
	5,  // 6: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.ChangeStatus:input_type -> tada.billing.api.personal_account.enquiry.v1.ChangeStatusRequest
	6,  // 7: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.SetLifecycleDates:input_type -> tada.billing.api.personal_account.enquiry.v1.SetLifecycleDatesRequest
	7,  // 8: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Create:output_type -> tada.billing.api.personal_account.enquiry.v1.CreateResponse
	8,  // 9: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Update:output_type -> tada.billing.api.personal_account.enquiry.v1.UpdateResponse
	9,  // 10: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Get:output_type -> tada.billing.api.personal_account.enquiry.v1.GetListResponse
	9,  // 11: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.GetList:output_type -> tada.billing.api.personal_account.enquiry.v1.GetListResponse
	10, // 12: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.GetCountsList:output_type -> tada.billing.api.personal_account.enquiry.v1.GetCountsListResponse
	11, // 13: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.Delete:output_type -> tada.billing.api.personal_account.enquiry.v1.DeleteResponse
	12, // 14: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.ChangeStatus:output_type -> tada.billing.api.personal_account.enquiry.v1.ChangeStatusResponse
	13, // 15: tada.gateway.admin.billing.personal_account.enquiry.v1.EnquiryService.SetLifecycleDates:output_type -> tada.billing.api.personal_account.enquiry.v1.SetLifecycleDatesResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_init() }
func file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_init() {
	if File_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_goTypes,
		DependencyIndexes: file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_depIdxs,
	}.Build()
	File_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto = out.File
	file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_rawDesc = nil
	file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_goTypes = nil
	file_tada_gateway_admin_billing_personal_account_enquiry_v1_enquiry_service_proto_depIdxs = nil
}
