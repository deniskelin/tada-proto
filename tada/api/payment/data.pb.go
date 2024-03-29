// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tada/api/payment/data.proto

package payment

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

type PaymentMethod int32

const (
	// Неизвестный метод оплаты
	PaymentMethod_PAYMENT_METHOD_UNKNOWN PaymentMethod = 0
	// Платеж банковской картой
	PaymentMethod_PAYMENT_METHOD_BANK_CARD PaymentMethod = 1
	// Платеж со счета мобильного номера
	PaymentMethod_PAYMENT_METHOD_MOBILE_BALANCE PaymentMethod = 2
	// Платеж наличными
	PaymentMethod_PAYMENT_METHOD_CASH PaymentMethod = 3
	// Рассрочка
	PaymentMethod_PAYMENT_METHOD_INSTALLMENTS PaymentMethod = 4
	// Платеж через Apple Pay
	PaymentMethod_PAYMENT_METHOD_APPLE_PAY PaymentMethod = 5
	// Платеж через GooglePay
	PaymentMethod_PAYMENT_METHOD_GOOGLE_PAY PaymentMethod = 6
	// Платеж через выставление счета для оплаты через банк
	PaymentMethod_PAYMENT_METHOD_INVOICE PaymentMethod = 7
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "PAYMENT_METHOD_UNKNOWN",
		1: "PAYMENT_METHOD_BANK_CARD",
		2: "PAYMENT_METHOD_MOBILE_BALANCE",
		3: "PAYMENT_METHOD_CASH",
		4: "PAYMENT_METHOD_INSTALLMENTS",
		5: "PAYMENT_METHOD_APPLE_PAY",
		6: "PAYMENT_METHOD_GOOGLE_PAY",
		7: "PAYMENT_METHOD_INVOICE",
	}
	PaymentMethod_value = map[string]int32{
		"PAYMENT_METHOD_UNKNOWN":        0,
		"PAYMENT_METHOD_BANK_CARD":      1,
		"PAYMENT_METHOD_MOBILE_BALANCE": 2,
		"PAYMENT_METHOD_CASH":           3,
		"PAYMENT_METHOD_INSTALLMENTS":   4,
		"PAYMENT_METHOD_APPLE_PAY":      5,
		"PAYMENT_METHOD_GOOGLE_PAY":     6,
		"PAYMENT_METHOD_INVOICE":        7,
	}
)

func (x PaymentMethod) Enum() *PaymentMethod {
	p := new(PaymentMethod)
	*p = x
	return p
}

func (x PaymentMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_tada_api_payment_data_proto_enumTypes[0].Descriptor()
}

func (PaymentMethod) Type() protoreflect.EnumType {
	return &file_tada_api_payment_data_proto_enumTypes[0]
}

func (x PaymentMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod.Descriptor instead.
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return file_tada_api_payment_data_proto_rawDescGZIP(), []int{0}
}

type CardType int32

const (
	// Неизвестный тип карты
	CardType_CARD_TYPE_UNKNOWN          CardType = 0
	CardType_CARD_TYPE_MASTER_CARD      CardType = 1
	CardType_CARD_TYPE_VISA             CardType = 2
	CardType_CARD_TYPE_MIR              CardType = 3
	CardType_CARD_TYPE_UNION_PAY        CardType = 4
	CardType_CARD_TYPE_JCB              CardType = 5
	CardType_CARD_TYPE_AMERICAN_EXPRESS CardType = 6
	CardType_CARD_TYPE_DINERS_CLUB      CardType = 7
)

// Enum value maps for CardType.
var (
	CardType_name = map[int32]string{
		0: "CARD_TYPE_UNKNOWN",
		1: "CARD_TYPE_MASTER_CARD",
		2: "CARD_TYPE_VISA",
		3: "CARD_TYPE_MIR",
		4: "CARD_TYPE_UNION_PAY",
		5: "CARD_TYPE_JCB",
		6: "CARD_TYPE_AMERICAN_EXPRESS",
		7: "CARD_TYPE_DINERS_CLUB",
	}
	CardType_value = map[string]int32{
		"CARD_TYPE_UNKNOWN":          0,
		"CARD_TYPE_MASTER_CARD":      1,
		"CARD_TYPE_VISA":             2,
		"CARD_TYPE_MIR":              3,
		"CARD_TYPE_UNION_PAY":        4,
		"CARD_TYPE_JCB":              5,
		"CARD_TYPE_AMERICAN_EXPRESS": 6,
		"CARD_TYPE_DINERS_CLUB":      7,
	}
)

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}

func (x CardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardType) Descriptor() protoreflect.EnumDescriptor {
	return file_tada_api_payment_data_proto_enumTypes[1].Descriptor()
}

func (CardType) Type() protoreflect.EnumType {
	return &file_tada_api_payment_data_proto_enumTypes[1]
}

func (x CardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardType.Descriptor instead.
func (CardType) EnumDescriptor() ([]byte, []int) {
	return file_tada_api_payment_data_proto_rawDescGZIP(), []int{1}
}

type PaymentStatus int32

const (
	// Неизвестный статус платежа
	PaymentStatus_PAYMENT_STATUS_UNKNOWN PaymentStatus = 0
	// Платеж создан и ожидает действий от пользователя. Из этого статуса платеж может перейти в PAYMENT_STATUS_SUCCEEDED, PAYMENT_STATUS_WAITING_FOR_CAPTURE
	// (при двухстадийной оплате) или PAYMENT_STATUS_CANCELED (если что-то пошло не так).
	PaymentStatus_PAYMENT_STATUS_PENDING PaymentStatus = 1
	// Платеж оплачен, деньги авторизованы и ожидают списания. Из этого статуса платеж может перейти в PAYMENT_STATUS_SUCCEEDED
	// (если вы списали оплату) или PAYMENT_STATUS_CANCELED (если вы отменили платеж или что-то пошло не так).
	PaymentStatus_PAYMENT_STATUS_WAITING_FOR_CAPTURE PaymentStatus = 2
	// Платеж успешно завершен, деньги будут перечислены на ваш расчетный счет.
	// Это финальный и неизменяемый статус.
	PaymentStatus_PAYMENT_STATUS_SUCCEEDED PaymentStatus = 3
	// Платеж отменен. Вы увидите этот статус, если вы отменили платеж самостоятельно, истекло время на принятие платежа
	// или платеж был отклонен Платежной Системой (напр. Юкасса) или платежным провайдером. Это финальный и неизменяемый статус.
	PaymentStatus_PAYMENT_STATUS_CANCELED PaymentStatus = 4
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PAYMENT_STATUS_UNKNOWN",
		1: "PAYMENT_STATUS_PENDING",
		2: "PAYMENT_STATUS_WAITING_FOR_CAPTURE",
		3: "PAYMENT_STATUS_SUCCEEDED",
		4: "PAYMENT_STATUS_CANCELED",
	}
	PaymentStatus_value = map[string]int32{
		"PAYMENT_STATUS_UNKNOWN":             0,
		"PAYMENT_STATUS_PENDING":             1,
		"PAYMENT_STATUS_WAITING_FOR_CAPTURE": 2,
		"PAYMENT_STATUS_SUCCEEDED":           3,
		"PAYMENT_STATUS_CANCELED":            4,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tada_api_payment_data_proto_enumTypes[2].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_tada_api_payment_data_proto_enumTypes[2]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_tada_api_payment_data_proto_rawDescGZIP(), []int{2}
}

type Currency int32

const (
	// Неизвестная валюта
	Currency_CURRENCY_UNKNOWN Currency = 0
	// Российский Рубль
	Currency_CURRENCY_RUB Currency = 1
	// Доллар США
	Currency_CURRENCY_USD Currency = 2
	// Евро
	Currency_CURRENCY_EUR Currency = 3
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "CURRENCY_UNKNOWN",
		1: "CURRENCY_RUB",
		2: "CURRENCY_USD",
		3: "CURRENCY_EUR",
	}
	Currency_value = map[string]int32{
		"CURRENCY_UNKNOWN": 0,
		"CURRENCY_RUB":     1,
		"CURRENCY_USD":     2,
		"CURRENCY_EUR":     3,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_tada_api_payment_data_proto_enumTypes[3].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_tada_api_payment_data_proto_enumTypes[3]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_tada_api_payment_data_proto_rawDescGZIP(), []int{3}
}

var File_tada_api_payment_data_proto protoreflect.FileDescriptor

var file_tada_api_payment_data_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74,
	0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2a,
	0xff, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x50,
	0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4d, 0x4f,
	0x42, 0x49, 0x4c, 0x45, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45,
	0x5f, 0x50, 0x41, 0x59, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f,
	0x50, 0x41, 0x59, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x10,
	0x07, 0x2a, 0xca, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49,
	0x53, 0x41, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x49, 0x52, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x59, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x43,
	0x42, 0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x41, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x49, 0x4e, 0x45, 0x52, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x42, 0x10, 0x07, 0x2a, 0xaa,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x55, 0x52, 0x45, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b,
	0x0a, 0x17, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x56, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x52, 0x55, 0x42, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x53, 0x44, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x45, 0x55,
	0x52, 0x10, 0x03, 0x42, 0x90, 0x01, 0x0a, 0x10, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x54, 0x61, 0x64, 0x61, 0x41, 0x50,
	0x49, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x2f, 0x74, 0x64, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0xf8, 0x01, 0x01,
	0xaa, 0x02, 0x10, 0x74, 0x61, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0xca, 0x02, 0x10, 0x54, 0x61, 0x64, 0x61, 0x5c, 0x41, 0x50, 0x49, 0x5c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tada_api_payment_data_proto_rawDescOnce sync.Once
	file_tada_api_payment_data_proto_rawDescData = file_tada_api_payment_data_proto_rawDesc
)

func file_tada_api_payment_data_proto_rawDescGZIP() []byte {
	file_tada_api_payment_data_proto_rawDescOnce.Do(func() {
		file_tada_api_payment_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_tada_api_payment_data_proto_rawDescData)
	})
	return file_tada_api_payment_data_proto_rawDescData
}

var file_tada_api_payment_data_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_tada_api_payment_data_proto_goTypes = []interface{}{
	(PaymentMethod)(0), // 0: tada.api.payment.PaymentMethod
	(CardType)(0),      // 1: tada.api.payment.CardType
	(PaymentStatus)(0), // 2: tada.api.payment.PaymentStatus
	(Currency)(0),      // 3: tada.api.payment.Currency
}
var file_tada_api_payment_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tada_api_payment_data_proto_init() }
func file_tada_api_payment_data_proto_init() {
	if File_tada_api_payment_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tada_api_payment_data_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tada_api_payment_data_proto_goTypes,
		DependencyIndexes: file_tada_api_payment_data_proto_depIdxs,
		EnumInfos:         file_tada_api_payment_data_proto_enumTypes,
	}.Build()
	File_tada_api_payment_data_proto = out.File
	file_tada_api_payment_data_proto_rawDesc = nil
	file_tada_api_payment_data_proto_goTypes = nil
	file_tada_api_payment_data_proto_depIdxs = nil
}
