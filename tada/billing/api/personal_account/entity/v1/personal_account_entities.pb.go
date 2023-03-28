// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/billing/api/personal_account/entity/v1/personal_account_entities.proto

package personal_account

import (
	v1 "gitlab.tada.team/tdapis/go-genproto/tada/billing/api/tariff/v1"
	common "gitlab.tada.team/tdapis/go-genproto/tada/common"
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
	//	*Identity_PersonalAccountIds
	//	*Identity_OwnerUuid
	//	*Identity_TeamUuid
	Identity isIdentity_Identity `protobuf_oneof:"identity"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[0]
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
	return file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescGZIP(), []int{0}
}

func (m *Identity) GetIdentity() isIdentity_Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (x *Identity) GetPersonalAccountIds() *common.ArrayId {
	if x, ok := x.GetIdentity().(*Identity_PersonalAccountIds); ok {
		return x.PersonalAccountIds
	}
	return nil
}

func (x *Identity) GetOwnerUuid() string {
	if x, ok := x.GetIdentity().(*Identity_OwnerUuid); ok {
		return x.OwnerUuid
	}
	return ""
}

func (x *Identity) GetTeamUuid() string {
	if x, ok := x.GetIdentity().(*Identity_TeamUuid); ok {
		return x.TeamUuid
	}
	return ""
}

type isIdentity_Identity interface {
	isIdentity_Identity()
}

type Identity_PersonalAccountIds struct {
	PersonalAccountIds *common.ArrayId `protobuf:"bytes,1,opt,name=personal_account_ids,json=personalAccountIds,proto3,oneof"`
}

type Identity_OwnerUuid struct {
	OwnerUuid string `protobuf:"bytes,2,opt,name=owner_uuid,json=ownerUuid,proto3,oneof"`
}

type Identity_TeamUuid struct {
	TeamUuid string `protobuf:"bytes,3,opt,name=team_uuid,json=teamUuid,proto3,oneof"`
}

func (*Identity_PersonalAccountIds) isIdentity_Identity() {}

func (*Identity_OwnerUuid) isIdentity_Identity() {}

func (*Identity_TeamUuid) isIdentity_Identity() {}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId                uint64                 `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	TeamsCount             uint32                 `protobuf:"varint,8,opt,name=teams_count,json=teamsCount,proto3" json:"teams_count,omitempty"`
	WorkplaceCount         uint32                 `protobuf:"varint,9,opt,name=workplace_count,json=workplaceCount,proto3" json:"workplace_count,omitempty"`
	EmptyWorkplaceCount    uint32                 `protobuf:"varint,10,opt,name=empty_workplace_count,json=emptyWorkplaceCount,proto3" json:"empty_workplace_count,omitempty"`
	OccupiedWorkplaceCount uint32                 `protobuf:"varint,11,opt,name=occupied_workplace_count,json=occupiedWorkplaceCount,proto3" json:"occupied_workplace_count,omitempty"`
	FreeWorkplaceCount     uint32                 `protobuf:"varint,12,opt,name=free_workplace_count,json=freeWorkplaceCount,proto3" json:"free_workplace_count,omitempty"`
	PaidWorkplaceCount     uint32                 `protobuf:"varint,13,opt,name=paid_workplace_count,json=paidWorkplaceCount,proto3" json:"paid_workplace_count,omitempty"`
	OwnerUuid              string                 `protobuf:"bytes,14,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	IsBlocked              bool                   `protobuf:"varint,21,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	IsSuspended            bool                   `protobuf:"varint,22,opt,name=is_suspended,json=isSuspended,proto3" json:"is_suspended,omitempty"`
	NextBillingDate        *timestamppb.Timestamp `protobuf:"bytes,28,opt,name=next_billing_date,json=nextBillingDate,proto3" json:"next_billing_date,omitempty"`
	BlockDate              *timestamppb.Timestamp `protobuf:"bytes,29,opt,name=block_date,json=blockDate,proto3,oneof" json:"block_date,omitempty"`
	SuspendDate            *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=suspend_date,json=suspendDate,proto3,oneof" json:"suspend_date,omitempty"`
	Status                 PersonalAccountStatus  `protobuf:"varint,36,opt,name=status,proto3,enum=tada.billing.api.personal_account.entity.v1.PersonalAccountStatus" json:"status,omitempty"`
	Tariff                 *v1.Element            `protobuf:"bytes,37,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[1]
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
	return file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Element) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Element) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Element) GetTeamsCount() uint32 {
	if x != nil {
		return x.TeamsCount
	}
	return 0
}

func (x *Element) GetWorkplaceCount() uint32 {
	if x != nil {
		return x.WorkplaceCount
	}
	return 0
}

func (x *Element) GetEmptyWorkplaceCount() uint32 {
	if x != nil {
		return x.EmptyWorkplaceCount
	}
	return 0
}

func (x *Element) GetOccupiedWorkplaceCount() uint32 {
	if x != nil {
		return x.OccupiedWorkplaceCount
	}
	return 0
}

func (x *Element) GetFreeWorkplaceCount() uint32 {
	if x != nil {
		return x.FreeWorkplaceCount
	}
	return 0
}

func (x *Element) GetPaidWorkplaceCount() uint32 {
	if x != nil {
		return x.PaidWorkplaceCount
	}
	return 0
}

func (x *Element) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

func (x *Element) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *Element) GetIsSuspended() bool {
	if x != nil {
		return x.IsSuspended
	}
	return false
}

func (x *Element) GetNextBillingDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NextBillingDate
	}
	return nil
}

func (x *Element) GetBlockDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockDate
	}
	return nil
}

func (x *Element) GetSuspendDate() *timestamppb.Timestamp {
	if x != nil {
		return x.SuspendDate
	}
	return nil
}

func (x *Element) GetStatus() PersonalAccountStatus {
	if x != nil {
		return x.Status
	}
	return PersonalAccountStatus_PERSONAL_ACCOUNT_STATUS_UNSPECIFIED
}

func (x *Element) GetTariff() *v1.Element {
	if x != nil {
		return x.Tariff
	}
	return nil
}

type UserElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserUuid string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserElement) Reset() {
	*x = UserElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserElement) ProtoMessage() {}

func (x *UserElement) ProtoReflect() protoreflect.Message {
	mi := &file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserElement.ProtoReflect.Descriptor instead.
func (*UserElement) Descriptor() ([]byte, []int) {
	return file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescGZIP(), []int{2}
}

func (x *UserElement) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserElement) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

var File_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto protoreflect.FileDescriptor

var file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x74, 0x61, 0x64,
	0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x14, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x49, 0x64, 0x48, 0x00, 0x52, 0x12, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x65, 0x61, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0xd4, 0x06, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x69, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x6f, 0x63, 0x63,
	0x75, 0x70, 0x69, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x66, 0x72, 0x65, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x70, 0x61, 0x69, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53,
	0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x42, 0x0a, 0x0c, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x0b, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x24,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3b, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x0f, 0x10, 0x15, 0x4a, 0x04, 0x08, 0x17, 0x10, 0x1c, 0x4a,
	0x04, 0x08, 0x1f, 0x10, 0x24, 0x4a, 0x04, 0x08, 0x26, 0x10, 0x33, 0x22, 0x43, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x42, 0xeb, 0x01, 0x0a, 0x2b, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e,
	0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0xaa, 0x02, 0x2a, 0x54, 0x61, 0x64, 0x61, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x2a, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5c, 0x41, 0x70, 0x69, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescOnce sync.Once
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescData = file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDesc
)

func file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescGZIP() []byte {
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescOnce.Do(func() {
		file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescData)
	})
	return file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDescData
}

var file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_goTypes = []interface{}{
	(*Identity)(nil),              // 0: tada.billing.api.personal_account.entity.v1.Identity
	(*Element)(nil),               // 1: tada.billing.api.personal_account.entity.v1.Element
	(*UserElement)(nil),           // 2: tada.billing.api.personal_account.entity.v1.UserElement
	(*common.ArrayId)(nil),        // 3: tada.common.ArrayId
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(PersonalAccountStatus)(0),    // 5: tada.billing.api.personal_account.entity.v1.PersonalAccountStatus
	(*v1.Element)(nil),            // 6: tada.billing.api.tariff.v1.Element
}
var file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_depIdxs = []int32{
	3, // 0: tada.billing.api.personal_account.entity.v1.Identity.personal_account_ids:type_name -> tada.common.ArrayId
	4, // 1: tada.billing.api.personal_account.entity.v1.Element.next_billing_date:type_name -> google.protobuf.Timestamp
	4, // 2: tada.billing.api.personal_account.entity.v1.Element.block_date:type_name -> google.protobuf.Timestamp
	4, // 3: tada.billing.api.personal_account.entity.v1.Element.suspend_date:type_name -> google.protobuf.Timestamp
	5, // 4: tada.billing.api.personal_account.entity.v1.Element.status:type_name -> tada.billing.api.personal_account.entity.v1.PersonalAccountStatus
	6, // 5: tada.billing.api.personal_account.entity.v1.Element.tariff:type_name -> tada.billing.api.tariff.v1.Element
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_init() }
func file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_init() {
	if File_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto != nil {
		return
	}
	file_tada_billing_api_personal_account_entity_v1_personal_account_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserElement); i {
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
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Identity_PersonalAccountIds)(nil),
		(*Identity_OwnerUuid)(nil),
		(*Identity_TeamUuid)(nil),
	}
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_goTypes,
		DependencyIndexes: file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_depIdxs,
		MessageInfos:      file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_msgTypes,
	}.Build()
	File_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto = out.File
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_rawDesc = nil
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_goTypes = nil
	file_tada_billing_api_personal_account_entity_v1_personal_account_entities_proto_depIdxs = nil
}
