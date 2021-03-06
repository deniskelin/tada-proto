// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: tada/gateway/admin/tariff_service.proto

package admin

import (
	billing "github.com/deniskelin/tada-proto/tada/api/billing"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// UpdateTariffRequest request on update tariff
type UpdateTariffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TariffId        int64                  `protobuf:"varint,1,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
	IsDefaultTariff *bool                  `protobuf:"varint,2,opt,name=is_default_tariff,json=isDefaultTariff,proto3,oneof" json:"is_default_tariff,omitempty"`
	Status          *billing.TariffStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=tada.api.billing.TariffStatus,oneof" json:"status,omitempty"`
	CloseDate       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=close_date,json=closeDate,proto3,oneof" json:"close_date,omitempty"`
}

func (x *UpdateTariffRequest) Reset() {
	*x = UpdateTariffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_gateway_admin_tariff_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTariffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTariffRequest) ProtoMessage() {}

func (x *UpdateTariffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tada_gateway_admin_tariff_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTariffRequest.ProtoReflect.Descriptor instead.
func (*UpdateTariffRequest) Descriptor() ([]byte, []int) {
	return file_tada_gateway_admin_tariff_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTariffRequest) GetTariffId() int64 {
	if x != nil {
		return x.TariffId
	}
	return 0
}

func (x *UpdateTariffRequest) GetIsDefaultTariff() bool {
	if x != nil && x.IsDefaultTariff != nil {
		return *x.IsDefaultTariff
	}
	return false
}

func (x *UpdateTariffRequest) GetStatus() billing.TariffStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return billing.TariffStatus(0)
}

func (x *UpdateTariffRequest) GetCloseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseDate
	}
	return nil
}

// UpdateTariffResponse response on update tariff
type UpdateTariffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TariffId             int64                  `protobuf:"varint,1,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
	TariffName           string                 `protobuf:"bytes,2,opt,name=tariff_name,json=tariffName,proto3" json:"tariff_name,omitempty"`
	FreeWorkplaces       int32                  `protobuf:"varint,3,opt,name=free_workplaces,json=freeWorkplaces,proto3" json:"free_workplaces,omitempty"`
	MinTariffWorkplaces  int32                  `protobuf:"varint,4,opt,name=min_tariff_workplaces,json=minTariffWorkplaces,proto3" json:"min_tariff_workplaces,omitempty"`
	MinStepWorkplaces    int32                  `protobuf:"varint,5,opt,name=min_step_workplaces,json=minStepWorkplaces,proto3" json:"min_step_workplaces,omitempty"`
	DiskSpaceQuotaMb     float64                `protobuf:"fixed64,6,opt,name=disk_space_quota_mb,json=diskSpaceQuotaMb,proto3" json:"disk_space_quota_mb,omitempty"`
	IsBillingFree        bool                   `protobuf:"varint,7,opt,name=is_billing_free,json=isBillingFree,proto3" json:"is_billing_free,omitempty"`
	IsBillingFullTime    bool                   `protobuf:"varint,8,opt,name=is_billing_full_time,json=isBillingFullTime,proto3" json:"is_billing_full_time,omitempty"`
	PeriodDays           int32                  `protobuf:"varint,9,opt,name=period_days,json=periodDays,proto3" json:"period_days,omitempty"`
	CostWorkplace        float64                `protobuf:"fixed64,10,opt,name=cost_workplace,json=costWorkplace,proto3" json:"cost_workplace,omitempty"`
	Currency             string                 `protobuf:"bytes,11,opt,name=currency,proto3" json:"currency,omitempty"`
	IsRecalcChangeTariff bool                   `protobuf:"varint,12,opt,name=is_recalc_change_tariff,json=isRecalcChangeTariff,proto3" json:"is_recalc_change_tariff,omitempty"`
	MaxVoiceUser         int32                  `protobuf:"varint,13,opt,name=max_voice_user,json=maxVoiceUser,proto3" json:"max_voice_user,omitempty"`
	MaxVideoUser         int32                  `protobuf:"varint,14,opt,name=max_video_user,json=maxVideoUser,proto3" json:"max_video_user,omitempty"`
	VideoCallBitrate     int32                  `protobuf:"varint,15,opt,name=video_call_bitrate,json=videoCallBitrate,proto3" json:"video_call_bitrate,omitempty"`
	IsDefaultTariff      bool                   `protobuf:"varint,16,opt,name=is_default_tariff,json=isDefaultTariff,proto3" json:"is_default_tariff,omitempty"`
	OpenDate             *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	CloseDate            *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=close_date,json=closeDate,proto3" json:"close_date,omitempty"`
	Status               billing.TariffStatus   `protobuf:"varint,19,opt,name=status,proto3,enum=tada.api.billing.TariffStatus" json:"status,omitempty"`
}

func (x *UpdateTariffResponse) Reset() {
	*x = UpdateTariffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_gateway_admin_tariff_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTariffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTariffResponse) ProtoMessage() {}

func (x *UpdateTariffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tada_gateway_admin_tariff_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTariffResponse.ProtoReflect.Descriptor instead.
func (*UpdateTariffResponse) Descriptor() ([]byte, []int) {
	return file_tada_gateway_admin_tariff_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTariffResponse) GetTariffId() int64 {
	if x != nil {
		return x.TariffId
	}
	return 0
}

func (x *UpdateTariffResponse) GetTariffName() string {
	if x != nil {
		return x.TariffName
	}
	return ""
}

func (x *UpdateTariffResponse) GetFreeWorkplaces() int32 {
	if x != nil {
		return x.FreeWorkplaces
	}
	return 0
}

func (x *UpdateTariffResponse) GetMinTariffWorkplaces() int32 {
	if x != nil {
		return x.MinTariffWorkplaces
	}
	return 0
}

func (x *UpdateTariffResponse) GetMinStepWorkplaces() int32 {
	if x != nil {
		return x.MinStepWorkplaces
	}
	return 0
}

func (x *UpdateTariffResponse) GetDiskSpaceQuotaMb() float64 {
	if x != nil {
		return x.DiskSpaceQuotaMb
	}
	return 0
}

func (x *UpdateTariffResponse) GetIsBillingFree() bool {
	if x != nil {
		return x.IsBillingFree
	}
	return false
}

func (x *UpdateTariffResponse) GetIsBillingFullTime() bool {
	if x != nil {
		return x.IsBillingFullTime
	}
	return false
}

func (x *UpdateTariffResponse) GetPeriodDays() int32 {
	if x != nil {
		return x.PeriodDays
	}
	return 0
}

func (x *UpdateTariffResponse) GetCostWorkplace() float64 {
	if x != nil {
		return x.CostWorkplace
	}
	return 0
}

func (x *UpdateTariffResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *UpdateTariffResponse) GetIsRecalcChangeTariff() bool {
	if x != nil {
		return x.IsRecalcChangeTariff
	}
	return false
}

func (x *UpdateTariffResponse) GetMaxVoiceUser() int32 {
	if x != nil {
		return x.MaxVoiceUser
	}
	return 0
}

func (x *UpdateTariffResponse) GetMaxVideoUser() int32 {
	if x != nil {
		return x.MaxVideoUser
	}
	return 0
}

func (x *UpdateTariffResponse) GetVideoCallBitrate() int32 {
	if x != nil {
		return x.VideoCallBitrate
	}
	return 0
}

func (x *UpdateTariffResponse) GetIsDefaultTariff() bool {
	if x != nil {
		return x.IsDefaultTariff
	}
	return false
}

func (x *UpdateTariffResponse) GetOpenDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OpenDate
	}
	return nil
}

func (x *UpdateTariffResponse) GetCloseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseDate
	}
	return nil
}

func (x *UpdateTariffResponse) GetStatus() billing.TariffStatus {
	if x != nil {
		return x.Status
	}
	return billing.TariffStatus(0)
}

var File_tada_gateway_admin_tariff_service_proto protoreflect.FileDescriptor

var file_tada_gateway_admin_tariff_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x61, 0x64, 0x61, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x74, 0x61, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0xd6, 0x06, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x72, 0x65, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x69,
	0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70,
	0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x6d,
	0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x4d, 0x62, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x46, 0x72, 0x65,
	0x65, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x69, 0x73, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x73,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x61, 0x6c, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x52, 0x65, 0x63, 0x61, 0x6c,
	0x63, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x24, 0x0a,
	0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x61, 0x6c, 0x6c,
	0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xd1, 0x08, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x93, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x79, 0x49, 0x64, 0x12, 0x26, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x26, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0xa4, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x34,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0xa9, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2d, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x2f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x12, 0x25, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x1a, 0x25, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x24, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x61,
	0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x2a, 0x24, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x2f, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x41, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x2e, 0x74, 0x61, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x41, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x41, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x32, 0x2b, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x41, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x93, 0x01,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x27,
	0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x32, 0x25, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x3a, 0x01, 0x2a, 0x42, 0x95, 0x01, 0x0a, 0x12, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x15, 0x54, 0x61, 0x64, 0x61,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0xaa, 0x02, 0x12, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x12, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tada_gateway_admin_tariff_service_proto_rawDescOnce sync.Once
	file_tada_gateway_admin_tariff_service_proto_rawDescData = file_tada_gateway_admin_tariff_service_proto_rawDesc
)

func file_tada_gateway_admin_tariff_service_proto_rawDescGZIP() []byte {
	file_tada_gateway_admin_tariff_service_proto_rawDescOnce.Do(func() {
		file_tada_gateway_admin_tariff_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_gateway_admin_tariff_service_proto_rawDescData)
	})
	return file_tada_gateway_admin_tariff_service_proto_rawDescData
}

var file_tada_gateway_admin_tariff_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tada_gateway_admin_tariff_service_proto_goTypes = []interface{}{
	(*UpdateTariffRequest)(nil),                  // 0: tada.gateway.admin.UpdateTariffRequest
	(*UpdateTariffResponse)(nil),                 // 1: tada.gateway.admin.UpdateTariffResponse
	(billing.TariffStatus)(0),                    // 2: tada.api.billing.TariffStatus
	(*timestamppb.Timestamp)(nil),                // 3: google.protobuf.Timestamp
	(*billing.GetTariffByIdRequest)(nil),         // 4: tada.api.billing.GetTariffByIdRequest
	(*billing.GetTariffsListRequest)(nil),        // 5: tada.api.billing.GetTariffsListRequest
	(*billing.GetActiveTariffsListRequest)(nil),  // 6: tada.api.billing.GetActiveTariffsListRequest
	(*billing.CreateTariffRequest)(nil),          // 7: tada.api.billing.CreateTariffRequest
	(*billing.CloseTariffRequest)(nil),           // 8: tada.api.billing.CloseTariffRequest
	(*billing.SetTariffAsDefaultRequest)(nil),    // 9: tada.api.billing.SetTariffAsDefaultRequest
	(*billing.GetTariffByIdResponse)(nil),        // 10: tada.api.billing.GetTariffByIdResponse
	(*billing.GetTariffsListResponse)(nil),       // 11: tada.api.billing.GetTariffsListResponse
	(*billing.GetActiveTariffsListResponse)(nil), // 12: tada.api.billing.GetActiveTariffsListResponse
	(*billing.CreateTariffResponse)(nil),         // 13: tada.api.billing.CreateTariffResponse
	(*billing.CloseTariffResponse)(nil),          // 14: tada.api.billing.CloseTariffResponse
	(*billing.SetTariffAsDefaultResponse)(nil),   // 15: tada.api.billing.SetTariffAsDefaultResponse
}
var file_tada_gateway_admin_tariff_service_proto_depIdxs = []int32{
	2,  // 0: tada.gateway.admin.UpdateTariffRequest.status:type_name -> tada.api.billing.TariffStatus
	3,  // 1: tada.gateway.admin.UpdateTariffRequest.close_date:type_name -> google.protobuf.Timestamp
	3,  // 2: tada.gateway.admin.UpdateTariffResponse.open_date:type_name -> google.protobuf.Timestamp
	3,  // 3: tada.gateway.admin.UpdateTariffResponse.close_date:type_name -> google.protobuf.Timestamp
	2,  // 4: tada.gateway.admin.UpdateTariffResponse.status:type_name -> tada.api.billing.TariffStatus
	4,  // 5: tada.gateway.admin.Tariff.GetTariffById:input_type -> tada.api.billing.GetTariffByIdRequest
	5,  // 6: tada.gateway.admin.Tariff.GetTariffsList:input_type -> tada.api.billing.GetTariffsListRequest
	6,  // 7: tada.gateway.admin.Tariff.GetActiveTariffsList:input_type -> tada.api.billing.GetActiveTariffsListRequest
	7,  // 8: tada.gateway.admin.Tariff.CreateTariff:input_type -> tada.api.billing.CreateTariffRequest
	8,  // 9: tada.gateway.admin.Tariff.CloseTariff:input_type -> tada.api.billing.CloseTariffRequest
	9,  // 10: tada.gateway.admin.Tariff.SetTariffAsDefault:input_type -> tada.api.billing.SetTariffAsDefaultRequest
	0,  // 11: tada.gateway.admin.Tariff.UpdateTariff:input_type -> tada.gateway.admin.UpdateTariffRequest
	10, // 12: tada.gateway.admin.Tariff.GetTariffById:output_type -> tada.api.billing.GetTariffByIdResponse
	11, // 13: tada.gateway.admin.Tariff.GetTariffsList:output_type -> tada.api.billing.GetTariffsListResponse
	12, // 14: tada.gateway.admin.Tariff.GetActiveTariffsList:output_type -> tada.api.billing.GetActiveTariffsListResponse
	13, // 15: tada.gateway.admin.Tariff.CreateTariff:output_type -> tada.api.billing.CreateTariffResponse
	14, // 16: tada.gateway.admin.Tariff.CloseTariff:output_type -> tada.api.billing.CloseTariffResponse
	15, // 17: tada.gateway.admin.Tariff.SetTariffAsDefault:output_type -> tada.api.billing.SetTariffAsDefaultResponse
	1,  // 18: tada.gateway.admin.Tariff.UpdateTariff:output_type -> tada.gateway.admin.UpdateTariffResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_tada_gateway_admin_tariff_service_proto_init() }
func file_tada_gateway_admin_tariff_service_proto_init() {
	if File_tada_gateway_admin_tariff_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tada_gateway_admin_tariff_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTariffRequest); i {
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
		file_tada_gateway_admin_tariff_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTariffResponse); i {
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
	file_tada_gateway_admin_tariff_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_gateway_admin_tariff_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tada_gateway_admin_tariff_service_proto_goTypes,
		DependencyIndexes: file_tada_gateway_admin_tariff_service_proto_depIdxs,
		MessageInfos:      file_tada_gateway_admin_tariff_service_proto_msgTypes,
	}.Build()
	File_tada_gateway_admin_tariff_service_proto = out.File
	file_tada_gateway_admin_tariff_service_proto_rawDesc = nil
	file_tada_gateway_admin_tariff_service_proto_goTypes = nil
	file_tada_gateway_admin_tariff_service_proto_depIdxs = nil
}
