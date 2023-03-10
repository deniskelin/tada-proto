// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/tariff/v1/tariff_entities.proto

package tariff

import (
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

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NomenclatureName         string                 `protobuf:"bytes,2,opt,name=nomenclature_name,json=nomenclatureName,proto3" json:"nomenclature_name,omitempty"`
	Description              string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Benefit                  string                 `protobuf:"bytes,4,opt,name=benefit,proto3" json:"benefit,omitempty"`
	Currency                 string                 `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	CostWorkplace            float64                `protobuf:"fixed64,10,opt,name=cost_workplace,json=costWorkplace,proto3" json:"cost_workplace,omitempty"`
	DiskSpaceQuotaMb         float64                `protobuf:"fixed64,11,opt,name=disk_space_quota_mb,json=diskSpaceQuotaMb,proto3" json:"disk_space_quota_mb,omitempty"`
	Id                       uint64                 `protobuf:"varint,17,opt,name=id,proto3" json:"id,omitempty"`
	Priority                 uint32                 `protobuf:"varint,23,opt,name=priority,proto3" json:"priority,omitempty"`
	FreeWorkplaceCount       uint32                 `protobuf:"varint,24,opt,name=free_workplace_count,json=freeWorkplaceCount,proto3" json:"free_workplace_count,omitempty"`
	MaxWorkplaceCount        *uint32                `protobuf:"varint,25,opt,name=max_workplace_count,json=maxWorkplaceCount,proto3,oneof" json:"max_workplace_count,omitempty"`
	MinWorkplaceCount        *uint32                `protobuf:"varint,26,opt,name=min_workplace_count,json=minWorkplaceCount,proto3,oneof" json:"min_workplace_count,omitempty"`
	StepIncreasingWorkplaces uint32                 `protobuf:"varint,27,opt,name=step_increasing_workplaces,json=stepIncreasingWorkplaces,proto3" json:"step_increasing_workplaces,omitempty"`
	PeriodDays               uint32                 `protobuf:"varint,28,opt,name=period_days,json=periodDays,proto3" json:"period_days,omitempty"`
	MaxVoiceUser             uint32                 `protobuf:"varint,29,opt,name=max_voice_user,json=maxVoiceUser,proto3" json:"max_voice_user,omitempty"`
	MaxVideoUser             uint32                 `protobuf:"varint,30,opt,name=max_video_user,json=maxVideoUser,proto3" json:"max_video_user,omitempty"`
	VideoCallBitrate         uint32                 `protobuf:"varint,31,opt,name=video_call_bitrate,json=videoCallBitrate,proto3" json:"video_call_bitrate,omitempty"`
	VideoSharingBitrate      uint32                 `protobuf:"varint,32,opt,name=video_sharing_bitrate,json=videoSharingBitrate,proto3" json:"video_sharing_bitrate,omitempty"`
	IsPublic                 bool                   `protobuf:"varint,38,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	IsFree                   bool                   `protobuf:"varint,39,opt,name=is_free,json=isFree,proto3" json:"is_free,omitempty"`
	IsDefault                bool                   `protobuf:"varint,40,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	OpenDate                 *timestamppb.Timestamp `protobuf:"bytes,46,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	CloseDate                *timestamppb.Timestamp `protobuf:"bytes,47,opt,name=close_date,json=closeDate,proto3,oneof" json:"close_date,omitempty"`
	Status                   TariffStatus           `protobuf:"varint,61,opt,name=status,proto3,enum=tada.billing.api.tariff.v1.TariffStatus" json:"status,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes[0]
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
	return file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Element) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Element) GetNomenclatureName() string {
	if x != nil {
		return x.NomenclatureName
	}
	return ""
}

func (x *Element) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Element) GetBenefit() string {
	if x != nil {
		return x.Benefit
	}
	return ""
}

func (x *Element) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Element) GetCostWorkplace() float64 {
	if x != nil {
		return x.CostWorkplace
	}
	return 0
}

func (x *Element) GetDiskSpaceQuotaMb() float64 {
	if x != nil {
		return x.DiskSpaceQuotaMb
	}
	return 0
}

func (x *Element) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Element) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Element) GetFreeWorkplaceCount() uint32 {
	if x != nil {
		return x.FreeWorkplaceCount
	}
	return 0
}

func (x *Element) GetMaxWorkplaceCount() uint32 {
	if x != nil && x.MaxWorkplaceCount != nil {
		return *x.MaxWorkplaceCount
	}
	return 0
}

func (x *Element) GetMinWorkplaceCount() uint32 {
	if x != nil && x.MinWorkplaceCount != nil {
		return *x.MinWorkplaceCount
	}
	return 0
}

func (x *Element) GetStepIncreasingWorkplaces() uint32 {
	if x != nil {
		return x.StepIncreasingWorkplaces
	}
	return 0
}

func (x *Element) GetPeriodDays() uint32 {
	if x != nil {
		return x.PeriodDays
	}
	return 0
}

func (x *Element) GetMaxVoiceUser() uint32 {
	if x != nil {
		return x.MaxVoiceUser
	}
	return 0
}

func (x *Element) GetMaxVideoUser() uint32 {
	if x != nil {
		return x.MaxVideoUser
	}
	return 0
}

func (x *Element) GetVideoCallBitrate() uint32 {
	if x != nil {
		return x.VideoCallBitrate
	}
	return 0
}

func (x *Element) GetVideoSharingBitrate() uint32 {
	if x != nil {
		return x.VideoSharingBitrate
	}
	return 0
}

func (x *Element) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Element) GetIsFree() bool {
	if x != nil {
		return x.IsFree
	}
	return false
}

func (x *Element) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *Element) GetOpenDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OpenDate
	}
	return nil
}

func (x *Element) GetCloseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseDate
	}
	return nil
}

func (x *Element) GetStatus() TariffStatus {
	if x != nil {
		return x.Status
	}
	return TariffStatus_TARIFF_STATUS_UNSPECIFIED
}

var File_tada_billing_api_tariff_v1_tariff_entities_proto protoreflect.FileDescriptor

var file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x08, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x6e, 0x6f, 0x6d, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x6f, 0x6d, 0x65, 0x6e,
	0x63, 0x6c, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x73,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x6d,
	0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x4d, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x66, 0x72, 0x65, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x57, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13,
	0x6d, 0x69, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x11, 0x6d, 0x69, 0x6e,
	0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x3c, 0x0a, 0x1a, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x73, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x44, 0x61, 0x79, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x61, 0x78, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x61, 0x6c, 0x6c, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x26, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x46, 0x72, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x2e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0a,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x11, 0x4a, 0x04, 0x08, 0x12, 0x10, 0x17, 0x4a,
	0x04, 0x08, 0x21, 0x10, 0x26, 0x4a, 0x04, 0x08, 0x29, 0x10, 0x2e, 0x4a, 0x04, 0x08, 0x30, 0x10,
	0x3d, 0x4a, 0x04, 0x08, 0x3e, 0x10, 0x47, 0x42, 0x9f, 0x01, 0x0a, 0x1a, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61,
	0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0xaa,
	0x02, 0x1a, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x54,
	0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescOnce sync.Once
	file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescData = file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDesc
)

func file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescGZIP() []byte {
	file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescData)
	})
	return file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDescData
}

var file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tada_billing_api_tariff_v1_tariff_entities_proto_goTypes = []interface{}{
	(*Element)(nil),               // 0: tada.billing.api.tariff.v1.Element
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(TariffStatus)(0),             // 2: tada.billing.api.tariff.v1.TariffStatus
}
var file_tada_billing_api_tariff_v1_tariff_entities_proto_depIdxs = []int32{
	1, // 0: tada.billing.api.tariff.v1.Element.open_date:type_name -> google.protobuf.Timestamp
	1, // 1: tada.billing.api.tariff.v1.Element.close_date:type_name -> google.protobuf.Timestamp
	2, // 2: tada.billing.api.tariff.v1.Element.status:type_name -> tada.billing.api.tariff.v1.TariffStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tada_billing_api_tariff_v1_tariff_entities_proto_init() }
func file_tada_billing_api_tariff_v1_tariff_entities_proto_init() {
	if File_tada_billing_api_tariff_v1_tariff_entities_proto != nil {
		return
	}
	file_tada_billing_api_tariff_v1_tariff_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_tariff_v1_tariff_entities_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_tariff_v1_tariff_entities_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_tariff_v1_tariff_entities_proto_msgTypes,
	}.Build()
	File_tada_billing_api_tariff_v1_tariff_entities_proto = out.File
	file_tada_billing_api_tariff_v1_tariff_entities_proto_rawDesc = nil
	file_tada_billing_api_tariff_v1_tariff_entities_proto_goTypes = nil
	file_tada_billing_api_tariff_v1_tariff_entities_proto_depIdxs = nil
}
