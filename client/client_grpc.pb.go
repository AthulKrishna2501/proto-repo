// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: client/client.proto

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClientService_GetMasterOfCeremony_FullMethodName = "/client.ClientService/GetMasterOfCeremony"
	ClientService_HandleStripeEvent_FullMethodName   = "/client.ClientService/HandleStripeEvent"
	ClientService_VerifyPayment_FullMethodName       = "/client.ClientService/VerifyPayment"
	ClientService_ClientDashboard_FullMethodName     = "/client.ClientService/ClientDashboard"
	ClientService_CreateEvent_FullMethodName         = "/client.ClientService/CreateEvent"
)

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	GetMasterOfCeremony(ctx context.Context, in *MasterOfCeremonyRequest, opts ...grpc.CallOption) (*MasterOfCeremonyResponse, error)
	HandleStripeEvent(ctx context.Context, in *StripeWebhookRequest, opts ...grpc.CallOption) (*StripeWebhookResponse, error)
	VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error)
	ClientDashboard(ctx context.Context, in *LandingPageRequest, opts ...grpc.CallOption) (*LandingPageResponse, error)
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) GetMasterOfCeremony(ctx context.Context, in *MasterOfCeremonyRequest, opts ...grpc.CallOption) (*MasterOfCeremonyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterOfCeremonyResponse)
	err := c.cc.Invoke(ctx, ClientService_GetMasterOfCeremony_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) HandleStripeEvent(ctx context.Context, in *StripeWebhookRequest, opts ...grpc.CallOption) (*StripeWebhookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StripeWebhookResponse)
	err := c.cc.Invoke(ctx, ClientService_HandleStripeEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPaymentResponse)
	err := c.cc.Invoke(ctx, ClientService_VerifyPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) ClientDashboard(ctx context.Context, in *LandingPageRequest, opts ...grpc.CallOption) (*LandingPageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LandingPageResponse)
	err := c.cc.Invoke(ctx, ClientService_ClientDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, ClientService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility.
type ClientServiceServer interface {
	GetMasterOfCeremony(context.Context, *MasterOfCeremonyRequest) (*MasterOfCeremonyResponse, error)
	HandleStripeEvent(context.Context, *StripeWebhookRequest) (*StripeWebhookResponse, error)
	VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error)
	ClientDashboard(context.Context, *LandingPageRequest) (*LandingPageResponse, error)
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientServiceServer struct{}

func (UnimplementedClientServiceServer) GetMasterOfCeremony(context.Context, *MasterOfCeremonyRequest) (*MasterOfCeremonyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMasterOfCeremony not implemented")
}
func (UnimplementedClientServiceServer) HandleStripeEvent(context.Context, *StripeWebhookRequest) (*StripeWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleStripeEvent not implemented")
}
func (UnimplementedClientServiceServer) VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedClientServiceServer) ClientDashboard(context.Context, *LandingPageRequest) (*LandingPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientDashboard not implemented")
}
func (UnimplementedClientServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}
func (UnimplementedClientServiceServer) testEmbeddedByValue()                       {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_GetMasterOfCeremony_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterOfCeremonyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetMasterOfCeremony(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_GetMasterOfCeremony_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetMasterOfCeremony(ctx, req.(*MasterOfCeremonyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_HandleStripeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StripeWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).HandleStripeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_HandleStripeEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).HandleStripeEvent(ctx, req.(*StripeWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).VerifyPayment(ctx, req.(*VerifyPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_ClientDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LandingPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).ClientDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_ClientDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).ClientDashboard(ctx, req.(*LandingPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMasterOfCeremony",
			Handler:    _ClientService_GetMasterOfCeremony_Handler,
		},
		{
			MethodName: "HandleStripeEvent",
			Handler:    _ClientService_HandleStripeEvent_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _ClientService_VerifyPayment_Handler,
		},
		{
			MethodName: "ClientDashboard",
			Handler:    _ClientService_ClientDashboard_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _ClientService_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/client.proto",
}
