// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/personal_account/counterparty/v1/counterparty_entities.proto

package counterparty

import (
	common "github.com/deniskelin/tada-proto/tada/common"
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

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identity:
	//	*Identity_CounterpartyIds
	//	*Identity_AccountingDectionaryCode
	Identity isIdentity_Identity `protobuf_oneof:"identity"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[0]
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
	return file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescGZIP(), []int{0}
}

func (m *Identity) GetIdentity() isIdentity_Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (x *Identity) GetCounterpartyIds() *common.ArrayId {
	if x, ok := x.GetIdentity().(*Identity_CounterpartyIds); ok {
		return x.CounterpartyIds
	}
	return nil
}

func (x *Identity) GetAccountingDectionaryCode() string {
	if x, ok := x.GetIdentity().(*Identity_AccountingDectionaryCode); ok {
		return x.AccountingDectionaryCode
	}
	return ""
}

type isIdentity_Identity interface {
	isIdentity_Identity()
}

type Identity_CounterpartyIds struct {
	CounterpartyIds *common.ArrayId `protobuf:"bytes,1,opt,name=counterparty_ids,json=counterpartyIds,proto3,oneof"`
}

type Identity_AccountingDectionaryCode struct {
	AccountingDectionaryCode string `protobuf:"bytes,2,opt,name=accounting_dectionary_code,json=accountingDectionaryCode,proto3,oneof"`
}

func (*Identity_CounterpartyIds) isIdentity_Identity() {}

func (*Identity_AccountingDectionaryCode) isIdentity_Identity() {}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                                uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PersonalAccountId                 uint64                 `protobuf:"varint,2,opt,name=personal_account_id,json=personalAccountId,proto3" json:"personal_account_id,omitempty"`
	ElectronicDocumentManagementId    *uint64                `protobuf:"varint,3,opt,name=electronic_document_management_id,json=electronicDocumentManagementId,proto3,oneof" json:"electronic_document_management_id,omitempty"`
	FullName                          string                 `protobuf:"bytes,9,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	TaxpayerIdentificationNumber      string                 `protobuf:"bytes,10,opt,name=taxpayer_identification_number,json=taxpayerIdentificationNumber,proto3" json:"taxpayer_identification_number,omitempty"`
	LegalAddress                      string                 `protobuf:"bytes,11,opt,name=legal_address,json=legalAddress,proto3" json:"legal_address,omitempty"`
	PhysicalAddress                   string                 `protobuf:"bytes,12,opt,name=physical_address,json=physicalAddress,proto3" json:"physical_address,omitempty"`
	AccountingDectionaryCode          *string                `protobuf:"bytes,13,opt,name=accounting_dectionary_code,json=accountingDectionaryCode,proto3,oneof" json:"accounting_dectionary_code,omitempty"`
	ClassifierOfIndustrialEnterprises *string                `protobuf:"bytes,14,opt,name=classifier_of_industrial_enterprises,json=classifierOfIndustrialEnterprises,proto3,oneof" json:"classifier_of_industrial_enterprises,omitempty"`
	CreatedAt                         *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CounterpartyType                  CounterpartyType       `protobuf:"varint,21,opt,name=counterparty_type,json=counterpartyType,proto3,enum=tada.billing.api.personal_account.counterparty.v1.CounterpartyType" json:"counterparty_type,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[1]
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
	return file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Element) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Element) GetPersonalAccountId() uint64 {
	if x != nil {
		return x.PersonalAccountId
	}
	return 0
}

func (x *Element) GetElectronicDocumentManagementId() uint64 {
	if x != nil && x.ElectronicDocumentManagementId != nil {
		return *x.ElectronicDocumentManagementId
	}
	return 0
}

func (x *Element) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Element) GetTaxpayerIdentificationNumber() string {
	if x != nil {
		return x.TaxpayerIdentificationNumber
	}
	return ""
}

func (x *Element) GetLegalAddress() string {
	if x != nil {
		return x.LegalAddress
	}
	return ""
}

func (x *Element) GetPhysicalAddress() string {
	if x != nil {
		return x.PhysicalAddress
	}
	return ""
}

func (x *Element) GetAccountingDectionaryCode() string {
	if x != nil && x.AccountingDectionaryCode != nil {
		return *x.AccountingDectionaryCode
	}
	return ""
}

func (x *Element) GetClassifierOfIndustrialEnterprises() string {
	if x != nil && x.ClassifierOfIndustrialEnterprises != nil {
		return *x.ClassifierOfIndustrialEnterprises
	}
	return ""
}

func (x *Element) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Element) GetCounterpartyType() CounterpartyType {
	if x != nil {
		return x.CounterpartyType
	}
	return CounterpartyType_COUNTERPARTY_TYPE_UNSPECIFIED
}

var File_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto protoreflect.FileDescriptor

var file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x74,
	0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x49, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x73, 0x12, 0x3e, 0x0a, 0x1a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x18, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x92, 0x06, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x4e, 0x0a, 0x21, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x5f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x1e,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x1e, 0x74, 0x61, 0x78, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x74, 0x61, 0x78, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x65, 0x67,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x1a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x18, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x24, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x21, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x4f, 0x66, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x70, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x24, 0x0a, 0x22, 0x5f, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x27, 0x0a, 0x25, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x6f,
	0x66, 0x5f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x09, 0x4a, 0x04,
	0x08, 0x0f, 0x10, 0x10, 0x4a, 0x04, 0x08, 0x11, 0x10, 0x15, 0x42, 0xff, 0x01, 0x0a, 0x31, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e,
	0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0xaa, 0x02, 0x30, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x30, 0x54, 0x61, 0x64, 0x61,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescOnce sync.Once
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescData = file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDesc
)

func file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescGZIP() []byte {
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescData)
	})
	return file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDescData
}

var file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_goTypes = []interface{}{
	(*Identity)(nil),              // 0: tada.billing.api.personal_account.counterparty.v1.Identity
	(*Element)(nil),               // 1: tada.billing.api.personal_account.counterparty.v1.Element
	(*common.ArrayId)(nil),        // 2: tada.common.ArrayId
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(CounterpartyType)(0),         // 4: tada.billing.api.personal_account.counterparty.v1.CounterpartyType
}
var file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_depIdxs = []int32{
	2, // 0: tada.billing.api.personal_account.counterparty.v1.Identity.counterparty_ids:type_name -> tada.common.ArrayId
	3, // 1: tada.billing.api.personal_account.counterparty.v1.Element.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: tada.billing.api.personal_account.counterparty.v1.Element.counterparty_type:type_name -> tada.billing.api.personal_account.counterparty.v1.CounterpartyType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_init()
}
func file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_init() {
	if File_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto != nil {
		return
	}
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Identity_CounterpartyIds)(nil),
		(*Identity_AccountingDectionaryCode)(nil),
	}
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_msgTypes,
	}.Build()
	File_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto = out.File
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_rawDesc = nil
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_goTypes = nil
	file_tada_billing_api_personal_account_counterparty_v1_counterparty_entities_proto_depIdxs = nil
}
