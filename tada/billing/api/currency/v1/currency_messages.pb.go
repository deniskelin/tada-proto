// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/currency/v1/currency_messages.proto

package currency

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

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsActive *bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[0]
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
	return file_tada_billing_api_currency_v1_currency_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetListRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyList []*Element `protobuf:"bytes,1,rep,name=currency_list,json=currencyList,proto3" json:"currency_list,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[1]
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
	return file_tada_billing_api_currency_v1_currency_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetListResponse) GetCurrencyList() []*Element {
	if x != nil {
		return x.CurrencyList
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyId uint64 `protobuf:"varint,1,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	IsVisible  bool   `protobuf:"varint,2,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[2]
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
	return file_tada_billing_api_currency_v1_currency_messages_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetCurrencyId() uint64 {
	if x != nil {
		return x.CurrencyId
	}
	return 0
}

func (x *UpdateRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency *Element `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[3]
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
	return file_tada_billing_api_currency_v1_currency_messages_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetCurrency() *Element {
	if x != nil {
		return x.Currency
	}
	return nil
}

var File_tada_billing_api_currency_v1_currency_messages_proto protoreflect.FileDescriptor

var file_tada_billing_api_currency_v1_currency_messages_proto_rawDesc = []byte{
	0x0a, 0x34, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x5d, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x42, 0xa9, 0x01, 0x0a, 0x1c, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0xaa,
	0x02, 0x1c, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1c, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_currency_v1_currency_messages_proto_rawDescOnce sync.Once
	file_tada_billing_api_currency_v1_currency_messages_proto_rawDescData = file_tada_billing_api_currency_v1_currency_messages_proto_rawDesc
)

func file_tada_billing_api_currency_v1_currency_messages_proto_rawDescGZIP() []byte {
	file_tada_billing_api_currency_v1_currency_messages_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_currency_v1_currency_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_currency_v1_currency_messages_proto_rawDescData)
	})
	return file_tada_billing_api_currency_v1_currency_messages_proto_rawDescData
}

var file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tada_billing_api_currency_v1_currency_messages_proto_goTypes = []interface{}{
	(*GetListRequest)(nil),  // 0: tada.billing.api.currency.v1.GetListRequest
	(*GetListResponse)(nil), // 1: tada.billing.api.currency.v1.GetListResponse
	(*UpdateRequest)(nil),   // 2: tada.billing.api.currency.v1.UpdateRequest
	(*UpdateResponse)(nil),  // 3: tada.billing.api.currency.v1.UpdateResponse
	(*Element)(nil),         // 4: tada.billing.api.currency.v1.Element
}
var file_tada_billing_api_currency_v1_currency_messages_proto_depIdxs = []int32{
	4, // 0: tada.billing.api.currency.v1.GetListResponse.currency_list:type_name -> tada.billing.api.currency.v1.Element
	4, // 1: tada.billing.api.currency.v1.UpdateResponse.currency:type_name -> tada.billing.api.currency.v1.Element
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tada_billing_api_currency_v1_currency_messages_proto_init() }
func file_tada_billing_api_currency_v1_currency_messages_proto_init() {
	if File_tada_billing_api_currency_v1_currency_messages_proto != nil {
		return
	}
	file_tada_billing_api_currency_v1_currency_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_currency_v1_currency_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_currency_v1_currency_messages_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_currency_v1_currency_messages_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_currency_v1_currency_messages_proto_msgTypes,
	}.Build()
	File_tada_billing_api_currency_v1_currency_messages_proto = out.File
	file_tada_billing_api_currency_v1_currency_messages_proto_rawDesc = nil
	file_tada_billing_api_currency_v1_currency_messages_proto_goTypes = nil
	file_tada_billing_api_currency_v1_currency_messages_proto_depIdxs = nil
}
