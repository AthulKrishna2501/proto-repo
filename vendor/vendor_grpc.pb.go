// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: vendor/vendor.proto

package vendor

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
	VendorSevice_RequestCategory_FullMethodName       = "/vendor.VendorSevice/RequestCategory"
	VendorSevice_ListCategory_FullMethodName          = "/vendor.VendorSevice/ListCategory"
	VendorSevice_ApproveRejectCategory_FullMethodName = "/vendor.VendorSevice/ApproveRejectCategory"
	VendorSevice_VendorProfile_FullMethodName         = "/vendor.VendorSevice/VendorProfile"
	VendorSevice_UpdateProfile_FullMethodName         = "/vendor.VendorSevice/UpdateProfile"
	VendorSevice_CreateService_FullMethodName         = "/vendor.VendorSevice/CreateService"
	VendorSevice_UpdateService_FullMethodName         = "/vendor.VendorSevice/UpdateService"
	VendorSevice_DeleteService_FullMethodName         = "/vendor.VendorSevice/DeleteService"
	VendorSevice_ChangePassword_FullMethodName        = "/vendor.VendorSevice/ChangePassword"
)

// VendorSeviceClient is the client API for VendorSevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VendorSeviceClient interface {
	RequestCategory(ctx context.Context, in *RequestCategoryRequest, opts ...grpc.CallOption) (*RequestCategoryResponse, error)
	ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryResponse, error)
	ApproveRejectCategory(ctx context.Context, in *ApproveRejectCategoryRequest, opts ...grpc.CallOption) (*ApproveRejectCategoryResponse, error)
	VendorProfile(ctx context.Context, in *VendorProfileRequest, opts ...grpc.CallOption) (*VendorProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateVendorProfileRequest, opts ...grpc.CallOption) (*UpdateVendorProfileResponse, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
}

type vendorSeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorSeviceClient(cc grpc.ClientConnInterface) VendorSeviceClient {
	return &vendorSeviceClient{cc}
}

func (c *vendorSeviceClient) RequestCategory(ctx context.Context, in *RequestCategoryRequest, opts ...grpc.CallOption) (*RequestCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestCategoryResponse)
	err := c.cc.Invoke(ctx, VendorSevice_RequestCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoryResponse)
	err := c.cc.Invoke(ctx, VendorSevice_ListCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ApproveRejectCategory(ctx context.Context, in *ApproveRejectCategoryRequest, opts ...grpc.CallOption) (*ApproveRejectCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApproveRejectCategoryResponse)
	err := c.cc.Invoke(ctx, VendorSevice_ApproveRejectCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) VendorProfile(ctx context.Context, in *VendorProfileRequest, opts ...grpc.CallOption) (*VendorProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VendorProfileResponse)
	err := c.cc.Invoke(ctx, VendorSevice_VendorProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) UpdateProfile(ctx context.Context, in *UpdateVendorProfileRequest, opts ...grpc.CallOption) (*UpdateVendorProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVendorProfileResponse)
	err := c.cc.Invoke(ctx, VendorSevice_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, VendorSevice_CreateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, VendorSevice_UpdateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, VendorSevice_DeleteService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, VendorSevice_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorSeviceServer is the server API for VendorSevice service.
// All implementations must embed UnimplementedVendorSeviceServer
// for forward compatibility.
type VendorSeviceServer interface {
	RequestCategory(context.Context, *RequestCategoryRequest) (*RequestCategoryResponse, error)
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryResponse, error)
	ApproveRejectCategory(context.Context, *ApproveRejectCategoryRequest) (*ApproveRejectCategoryResponse, error)
	VendorProfile(context.Context, *VendorProfileRequest) (*VendorProfileResponse, error)
	UpdateProfile(context.Context, *UpdateVendorProfileRequest) (*UpdateVendorProfileResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	mustEmbedUnimplementedVendorSeviceServer()
}

// UnimplementedVendorSeviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVendorSeviceServer struct{}

func (UnimplementedVendorSeviceServer) RequestCategory(context.Context, *RequestCategoryRequest) (*RequestCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestCategory not implemented")
}
func (UnimplementedVendorSeviceServer) ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedVendorSeviceServer) ApproveRejectCategory(context.Context, *ApproveRejectCategoryRequest) (*ApproveRejectCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveRejectCategory not implemented")
}
func (UnimplementedVendorSeviceServer) VendorProfile(context.Context, *VendorProfileRequest) (*VendorProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorProfile not implemented")
}
func (UnimplementedVendorSeviceServer) UpdateProfile(context.Context, *UpdateVendorProfileRequest) (*UpdateVendorProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedVendorSeviceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedVendorSeviceServer) UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedVendorSeviceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedVendorSeviceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedVendorSeviceServer) mustEmbedUnimplementedVendorSeviceServer() {}
func (UnimplementedVendorSeviceServer) testEmbeddedByValue()                      {}

// UnsafeVendorSeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorSeviceServer will
// result in compilation errors.
type UnsafeVendorSeviceServer interface {
	mustEmbedUnimplementedVendorSeviceServer()
}

func RegisterVendorSeviceServer(s grpc.ServiceRegistrar, srv VendorSeviceServer) {
	// If the following call pancis, it indicates UnimplementedVendorSeviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VendorSevice_ServiceDesc, srv)
}

func _VendorSevice_RequestCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).RequestCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_RequestCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).RequestCategory(ctx, req.(*RequestCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_ListCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).ListCategory(ctx, req.(*ListCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_ApproveRejectCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRejectCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).ApproveRejectCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_ApproveRejectCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).ApproveRejectCategory(ctx, req.(*ApproveRejectCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_VendorProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).VendorProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_VendorProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).VendorProfile(ctx, req.(*VendorProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVendorProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).UpdateProfile(ctx, req.(*UpdateVendorProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_UpdateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorSevice_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VendorSevice_ServiceDesc is the grpc.ServiceDesc for VendorSevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendorSevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vendor.VendorSevice",
	HandlerType: (*VendorSeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestCategory",
			Handler:    _VendorSevice_RequestCategory_Handler,
		},
		{
			MethodName: "ListCategory",
			Handler:    _VendorSevice_ListCategory_Handler,
		},
		{
			MethodName: "ApproveRejectCategory",
			Handler:    _VendorSevice_ApproveRejectCategory_Handler,
		},
		{
			MethodName: "VendorProfile",
			Handler:    _VendorSevice_VendorProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _VendorSevice_UpdateProfile_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _VendorSevice_CreateService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _VendorSevice_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _VendorSevice_DeleteService_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _VendorSevice_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendor/vendor.proto",
}
