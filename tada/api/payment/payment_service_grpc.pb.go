// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/api/payment/payment_service.proto

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	// CreatePaymentLink Создать платежную ссылку с параметрами
	CreatePaymentLink(ctx context.Context, in *CreatePaymentLinkRequest, opts ...grpc.CallOption) (*CreatePaymentLinkResponse, error)
	// Провести платеж вручную (Manual Payment с использованием общей логики с ProcessPayment)
	MakeManualPayment(ctx context.Context, in *ManualPaymentRequest, opts ...grpc.CallOption) (*ManualPaymentResponse, error)
	//  Инициировать автоплатеж по определенному пользователю
	MakeAutoPayment(ctx context.Context, in *MakeAutoPaymentRequest, opts ...grpc.CallOption) (*MakeAutoPaymentResponse, error)
	// ProcessPayment Обработать платеж (запрос от юкассы и пр. к нам)
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentRequest, error)
	// Привязка карты к пользователю для авто платежа
	BindCard(ctx context.Context, in *BindCardRequest, opts ...grpc.CallOption) (*BindCardResponse, error)
	// GetUserInfo Получить информацию о пользователе (подключен ли автоплатеж и тд) - под вопросом!
	GetUsersCardList(ctx context.Context, in *GetUsersCardListRequest, opts ...grpc.CallOption) (*GetUsersCardListResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) CreatePaymentLink(ctx context.Context, in *CreatePaymentLinkRequest, opts ...grpc.CallOption) (*CreatePaymentLinkResponse, error) {
	out := new(CreatePaymentLinkResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/CreatePaymentLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) MakeManualPayment(ctx context.Context, in *ManualPaymentRequest, opts ...grpc.CallOption) (*ManualPaymentResponse, error) {
	out := new(ManualPaymentResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/MakeManualPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) MakeAutoPayment(ctx context.Context, in *MakeAutoPaymentRequest, opts ...grpc.CallOption) (*MakeAutoPaymentResponse, error) {
	out := new(MakeAutoPaymentResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/MakeAutoPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentRequest, error) {
	out := new(ProcessPaymentRequest)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/ProcessPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) BindCard(ctx context.Context, in *BindCardRequest, opts ...grpc.CallOption) (*BindCardResponse, error) {
	out := new(BindCardResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/BindCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetUsersCardList(ctx context.Context, in *GetUsersCardListRequest, opts ...grpc.CallOption) (*GetUsersCardListResponse, error) {
	out := new(GetUsersCardListResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Payment/GetUsersCardList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	// CreatePaymentLink Создать платежную ссылку с параметрами
	CreatePaymentLink(context.Context, *CreatePaymentLinkRequest) (*CreatePaymentLinkResponse, error)
	// Провести платеж вручную (Manual Payment с использованием общей логики с ProcessPayment)
	MakeManualPayment(context.Context, *ManualPaymentRequest) (*ManualPaymentResponse, error)
	//  Инициировать автоплатеж по определенному пользователю
	MakeAutoPayment(context.Context, *MakeAutoPaymentRequest) (*MakeAutoPaymentResponse, error)
	// ProcessPayment Обработать платеж (запрос от юкассы и пр. к нам)
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentRequest, error)
	// Привязка карты к пользователю для авто платежа
	BindCard(context.Context, *BindCardRequest) (*BindCardResponse, error)
	// GetUserInfo Получить информацию о пользователе (подключен ли автоплатеж и тд) - под вопросом!
	GetUsersCardList(context.Context, *GetUsersCardListRequest) (*GetUsersCardListResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) CreatePaymentLink(context.Context, *CreatePaymentLinkRequest) (*CreatePaymentLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaymentLink not implemented")
}
func (UnimplementedPaymentServer) MakeManualPayment(context.Context, *ManualPaymentRequest) (*ManualPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeManualPayment not implemented")
}
func (UnimplementedPaymentServer) MakeAutoPayment(context.Context, *MakeAutoPaymentRequest) (*MakeAutoPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAutoPayment not implemented")
}
func (UnimplementedPaymentServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServer) BindCard(context.Context, *BindCardRequest) (*BindCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindCard not implemented")
}
func (UnimplementedPaymentServer) GetUsersCardList(context.Context, *GetUsersCardListRequest) (*GetUsersCardListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersCardList not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_CreatePaymentLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).CreatePaymentLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/CreatePaymentLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).CreatePaymentLink(ctx, req.(*CreatePaymentLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_MakeManualPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).MakeManualPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/MakeManualPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).MakeManualPayment(ctx, req.(*ManualPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_MakeAutoPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeAutoPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).MakeAutoPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/MakeAutoPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).MakeAutoPayment(ctx, req.(*MakeAutoPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/ProcessPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_BindCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).BindCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/BindCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).BindCard(ctx, req.(*BindCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetUsersCardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersCardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetUsersCardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Payment/GetUsersCardList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetUsersCardList(ctx, req.(*GetUsersCardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.payment.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePaymentLink",
			Handler:    _Payment_CreatePaymentLink_Handler,
		},
		{
			MethodName: "MakeManualPayment",
			Handler:    _Payment_MakeManualPayment_Handler,
		},
		{
			MethodName: "MakeAutoPayment",
			Handler:    _Payment_MakeAutoPayment_Handler,
		},
		{
			MethodName: "ProcessPayment",
			Handler:    _Payment_ProcessPayment_Handler,
		},
		{
			MethodName: "BindCard",
			Handler:    _Payment_BindCard_Handler,
		},
		{
			MethodName: "GetUsersCardList",
			Handler:    _Payment_GetUsersCardList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/payment/payment_service.proto",
}
