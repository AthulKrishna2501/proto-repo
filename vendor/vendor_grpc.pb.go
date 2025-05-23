// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vendor

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
	GetVendorDashboard(ctx context.Context, in *GetVendorDashboardRequest, opts ...grpc.CallOption) (*GetVendorDashboardResponse, error)
	GetVendorServices(ctx context.Context, in *GetVendorServicesRequest, opts ...grpc.CallOption) (*GetVendorServicesResponse, error)
	GetBookingRequests(ctx context.Context, in *GetBookingRequestsRequest, opts ...grpc.CallOption) (*GetBookingRequestsResponse, error)
	ApproveBooking(ctx context.Context, in *ApproveBookingRequest, opts ...grpc.CallOption) (*ApproveBookingResponse, error)
	GetVendorWallet(ctx context.Context, in *GetVendorWalletRequest, opts ...grpc.CallOption) (*GetVendorWalletResponse, error)
	GetVendorTransactions(ctx context.Context, in *ViewVendorTransactionsRequest, opts ...grpc.CallOption) (*ViewVendorTransactionResponse, error)
}

type vendorSeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorSeviceClient(cc grpc.ClientConnInterface) VendorSeviceClient {
	return &vendorSeviceClient{cc}
}

func (c *vendorSeviceClient) RequestCategory(ctx context.Context, in *RequestCategoryRequest, opts ...grpc.CallOption) (*RequestCategoryResponse, error) {
	out := new(RequestCategoryResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/RequestCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryResponse, error) {
	out := new(ListCategoryResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/ListCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ApproveRejectCategory(ctx context.Context, in *ApproveRejectCategoryRequest, opts ...grpc.CallOption) (*ApproveRejectCategoryResponse, error) {
	out := new(ApproveRejectCategoryResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/ApproveRejectCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) VendorProfile(ctx context.Context, in *VendorProfileRequest, opts ...grpc.CallOption) (*VendorProfileResponse, error) {
	out := new(VendorProfileResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/VendorProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) UpdateProfile(ctx context.Context, in *UpdateVendorProfileRequest, opts ...grpc.CallOption) (*UpdateVendorProfileResponse, error) {
	out := new(UpdateVendorProfileResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) GetVendorDashboard(ctx context.Context, in *GetVendorDashboardRequest, opts ...grpc.CallOption) (*GetVendorDashboardResponse, error) {
	out := new(GetVendorDashboardResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/GetVendorDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) GetVendorServices(ctx context.Context, in *GetVendorServicesRequest, opts ...grpc.CallOption) (*GetVendorServicesResponse, error) {
	out := new(GetVendorServicesResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/GetVendorServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) GetBookingRequests(ctx context.Context, in *GetBookingRequestsRequest, opts ...grpc.CallOption) (*GetBookingRequestsResponse, error) {
	out := new(GetBookingRequestsResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/GetBookingRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) ApproveBooking(ctx context.Context, in *ApproveBookingRequest, opts ...grpc.CallOption) (*ApproveBookingResponse, error) {
	out := new(ApproveBookingResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/ApproveBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) GetVendorWallet(ctx context.Context, in *GetVendorWalletRequest, opts ...grpc.CallOption) (*GetVendorWalletResponse, error) {
	out := new(GetVendorWalletResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/GetVendorWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorSeviceClient) GetVendorTransactions(ctx context.Context, in *ViewVendorTransactionsRequest, opts ...grpc.CallOption) (*ViewVendorTransactionResponse, error) {
	out := new(ViewVendorTransactionResponse)
	err := c.cc.Invoke(ctx, "/vendor.VendorSevice/GetVendorTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorSeviceServer is the server API for VendorSevice service.
// All implementations must embed UnimplementedVendorSeviceServer
// for forward compatibility
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
	GetVendorDashboard(context.Context, *GetVendorDashboardRequest) (*GetVendorDashboardResponse, error)
	GetVendorServices(context.Context, *GetVendorServicesRequest) (*GetVendorServicesResponse, error)
	GetBookingRequests(context.Context, *GetBookingRequestsRequest) (*GetBookingRequestsResponse, error)
	ApproveBooking(context.Context, *ApproveBookingRequest) (*ApproveBookingResponse, error)
	GetVendorWallet(context.Context, *GetVendorWalletRequest) (*GetVendorWalletResponse, error)
	GetVendorTransactions(context.Context, *ViewVendorTransactionsRequest) (*ViewVendorTransactionResponse, error)
	mustEmbedUnimplementedVendorSeviceServer()
}

// UnimplementedVendorSeviceServer must be embedded to have forward compatible implementations.
type UnimplementedVendorSeviceServer struct {
}

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
func (UnimplementedVendorSeviceServer) GetVendorDashboard(context.Context, *GetVendorDashboardRequest) (*GetVendorDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendorDashboard not implemented")
}
func (UnimplementedVendorSeviceServer) GetVendorServices(context.Context, *GetVendorServicesRequest) (*GetVendorServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendorServices not implemented")
}
func (UnimplementedVendorSeviceServer) GetBookingRequests(context.Context, *GetBookingRequestsRequest) (*GetBookingRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingRequests not implemented")
}
func (UnimplementedVendorSeviceServer) ApproveBooking(context.Context, *ApproveBookingRequest) (*ApproveBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveBooking not implemented")
}
func (UnimplementedVendorSeviceServer) GetVendorWallet(context.Context, *GetVendorWalletRequest) (*GetVendorWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendorWallet not implemented")
}
func (UnimplementedVendorSeviceServer) GetVendorTransactions(context.Context, *ViewVendorTransactionsRequest) (*ViewVendorTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendorTransactions not implemented")
}
func (UnimplementedVendorSeviceServer) mustEmbedUnimplementedVendorSeviceServer() {}

// UnsafeVendorSeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorSeviceServer will
// result in compilation errors.
type UnsafeVendorSeviceServer interface {
	mustEmbedUnimplementedVendorSeviceServer()
}

func RegisterVendorSeviceServer(s grpc.ServiceRegistrar, srv VendorSeviceServer) {
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
		FullMethod: "/vendor.VendorSevice/RequestCategory",
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
		FullMethod: "/vendor.VendorSevice/ListCategory",
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
		FullMethod: "/vendor.VendorSevice/ApproveRejectCategory",
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
		FullMethod: "/vendor.VendorSevice/VendorProfile",
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
		FullMethod: "/vendor.VendorSevice/UpdateProfile",
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
		FullMethod: "/vendor.VendorSevice/CreateService",
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
		FullMethod: "/vendor.VendorSevice/UpdateService",
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
		FullMethod: "/vendor.VendorSevice/DeleteService",
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
		FullMethod: "/vendor.VendorSevice/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_GetVendorDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVendorDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).GetVendorDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/GetVendorDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).GetVendorDashboard(ctx, req.(*GetVendorDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_GetVendorServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVendorServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).GetVendorServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/GetVendorServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).GetVendorServices(ctx, req.(*GetVendorServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_GetBookingRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).GetBookingRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/GetBookingRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).GetBookingRequests(ctx, req.(*GetBookingRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_ApproveBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).ApproveBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/ApproveBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).ApproveBooking(ctx, req.(*ApproveBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_GetVendorWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVendorWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).GetVendorWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/GetVendorWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).GetVendorWallet(ctx, req.(*GetVendorWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorSevice_GetVendorTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewVendorTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorSeviceServer).GetVendorTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendor.VendorSevice/GetVendorTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorSeviceServer).GetVendorTransactions(ctx, req.(*ViewVendorTransactionsRequest))
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
		{
			MethodName: "GetVendorDashboard",
			Handler:    _VendorSevice_GetVendorDashboard_Handler,
		},
		{
			MethodName: "GetVendorServices",
			Handler:    _VendorSevice_GetVendorServices_Handler,
		},
		{
			MethodName: "GetBookingRequests",
			Handler:    _VendorSevice_GetBookingRequests_Handler,
		},
		{
			MethodName: "ApproveBooking",
			Handler:    _VendorSevice_ApproveBooking_Handler,
		},
		{
			MethodName: "GetVendorWallet",
			Handler:    _VendorSevice_GetVendorWallet_Handler,
		},
		{
			MethodName: "GetVendorTransactions",
			Handler:    _VendorSevice_GetVendorTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendor/vendor.proto",
}
