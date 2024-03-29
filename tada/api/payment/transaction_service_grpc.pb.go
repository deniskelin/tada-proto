// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/api/payment/transaction_service.proto

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

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	// GetTransactionList Получить список всех транзакций с их статусами и прочей инфой а так же получить список
	// транзакций конкретного пользователя (скорее всего 5 и 6 это один метод просто передали или нет идентификатор юзера)
	GetTransactionList(ctx context.Context, in *GetTransactionListRequest, opts ...grpc.CallOption) (*GetTransactionListResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) GetTransactionList(ctx context.Context, in *GetTransactionListRequest, opts ...grpc.CallOption) (*GetTransactionListResponse, error) {
	out := new(GetTransactionListResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Transaction/GetTransactionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	// GetTransactionList Получить список всех транзакций с их статусами и прочей инфой а так же получить список
	// транзакций конкретного пользователя (скорее всего 5 и 6 это один метод просто передали или нет идентификатор юзера)
	GetTransactionList(context.Context, *GetTransactionListRequest) (*GetTransactionListResponse, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) GetTransactionList(context.Context, *GetTransactionListRequest) (*GetTransactionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionList not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_GetTransactionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetTransactionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Transaction/GetTransactionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetTransactionList(ctx, req.(*GetTransactionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.payment.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionList",
			Handler:    _Transaction_GetTransactionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/payment/transaction_service.proto",
}
