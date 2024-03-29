// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tada/api/payment/invoice_service.proto

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

// InvoiceClient is the client API for Invoice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceClient interface {
	// Сформировать форму счета на оплату с занесением в базу информации о нем и о статусе (по аналогии с п1 и п2, их тоже нужно вносить в базу)
	CreateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error)
}

type invoiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceClient(cc grpc.ClientConnInterface) InvoiceClient {
	return &invoiceClient{cc}
}

func (c *invoiceClient) CreateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error) {
	out := new(GenerateInvoiceResponse)
	err := c.cc.Invoke(ctx, "/tada.api.payment.Invoice/CreateInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServer is the server API for Invoice service.
// All implementations must embed UnimplementedInvoiceServer
// for forward compatibility
type InvoiceServer interface {
	// Сформировать форму счета на оплату с занесением в базу информации о нем и о статусе (по аналогии с п1 и п2, их тоже нужно вносить в базу)
	CreateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error)
	mustEmbedUnimplementedInvoiceServer()
}

// UnimplementedInvoiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceServer struct {
}

func (UnimplementedInvoiceServer) CreateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedInvoiceServer) mustEmbedUnimplementedInvoiceServer() {}

// UnsafeInvoiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServer will
// result in compilation errors.
type UnsafeInvoiceServer interface {
	mustEmbedUnimplementedInvoiceServer()
}

func RegisterInvoiceServer(s grpc.ServiceRegistrar, srv InvoiceServer) {
	s.RegisterService(&Invoice_ServiceDesc, srv)
}

func _Invoice_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tada.api.payment.Invoice/CreateInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).CreateInvoice(ctx, req.(*GenerateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoice_ServiceDesc is the grpc.ServiceDesc for Invoice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tada.api.payment.Invoice",
	HandlerType: (*InvoiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoice",
			Handler:    _Invoice_CreateInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tada/api/payment/invoice_service.proto",
}
