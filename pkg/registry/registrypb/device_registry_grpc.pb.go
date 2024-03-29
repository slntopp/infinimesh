// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registrypb

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

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetByFingerprint(ctx context.Context, in *GetByFingerprintRequest, opts ...grpc.CallOption) (*GetByFingerprintResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	AssignOwnerDevices(ctx context.Context, in *OwnershipRequestDevices, opts ...grpc.CallOption) (*OwnershipResponseDevices, error)
	RemoveOwnerDevices(ctx context.Context, in *OwnershipRequestDevices, opts ...grpc.CallOption) (*OwnershipResponseDevices, error)
	GetDeviceStatus(ctx context.Context, in *GetDeviceStatusRequest, opts ...grpc.CallOption) (*GetDeviceStatusResponse, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetByFingerprint(ctx context.Context, in *GetByFingerprintRequest, opts ...grpc.CallOption) (*GetByFingerprintResponse, error) {
	out := new(GetByFingerprintResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/GetByFingerprint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) AssignOwnerDevices(ctx context.Context, in *OwnershipRequestDevices, opts ...grpc.CallOption) (*OwnershipResponseDevices, error) {
	out := new(OwnershipResponseDevices)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/AssignOwnerDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) RemoveOwnerDevices(ctx context.Context, in *OwnershipRequestDevices, opts ...grpc.CallOption) (*OwnershipResponseDevices, error) {
	out := new(OwnershipResponseDevices)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/RemoveOwnerDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDeviceStatus(ctx context.Context, in *GetDeviceStatusRequest, opts ...grpc.CallOption) (*GetDeviceStatusResponse, error) {
	out := new(GetDeviceStatusResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/GetDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
// All implementations must embed UnimplementedDevicesServer
// for forward compatibility
type DevicesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetByFingerprint(context.Context, *GetByFingerprintRequest) (*GetByFingerprintResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListDevicesRequest) (*ListResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	AssignOwnerDevices(context.Context, *OwnershipRequestDevices) (*OwnershipResponseDevices, error)
	RemoveOwnerDevices(context.Context, *OwnershipRequestDevices) (*OwnershipResponseDevices, error)
	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)
	mustEmbedUnimplementedDevicesServer()
}

// UnimplementedDevicesServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (UnimplementedDevicesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDevicesServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDevicesServer) GetByFingerprint(context.Context, *GetByFingerprintRequest) (*GetByFingerprintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByFingerprint not implemented")
}
func (UnimplementedDevicesServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDevicesServer) List(context.Context, *ListDevicesRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDevicesServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDevicesServer) AssignOwnerDevices(context.Context, *OwnershipRequestDevices) (*OwnershipResponseDevices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignOwnerDevices not implemented")
}
func (UnimplementedDevicesServer) RemoveOwnerDevices(context.Context, *OwnershipRequestDevices) (*OwnershipResponseDevices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOwnerDevices not implemented")
}
func (UnimplementedDevicesServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceStatus not implemented")
}
func (UnimplementedDevicesServer) mustEmbedUnimplementedDevicesServer() {}

// UnsafeDevicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesServer will
// result in compilation errors.
type UnsafeDevicesServer interface {
	mustEmbedUnimplementedDevicesServer()
}

func RegisterDevicesServer(s grpc.ServiceRegistrar, srv DevicesServer) {
	s.RegisterService(&Devices_ServiceDesc, srv)
}

func _Devices_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetByFingerprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByFingerprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetByFingerprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/GetByFingerprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetByFingerprint(ctx, req.(*GetByFingerprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).List(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_AssignOwnerDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnershipRequestDevices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).AssignOwnerDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/AssignOwnerDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).AssignOwnerDevices(ctx, req.(*OwnershipRequestDevices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_RemoveOwnerDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnershipRequestDevices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).RemoveOwnerDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/RemoveOwnerDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).RemoveOwnerDevices(ctx, req.(*OwnershipRequestDevices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/GetDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDeviceStatus(ctx, req.(*GetDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Devices_ServiceDesc is the grpc.ServiceDesc for Devices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Devices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.deviceregistry.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Devices_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Devices_Update_Handler,
		},
		{
			MethodName: "GetByFingerprint",
			Handler:    _Devices_GetByFingerprint_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Devices_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Devices_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Devices_Delete_Handler,
		},
		{
			MethodName: "AssignOwnerDevices",
			Handler:    _Devices_AssignOwnerDevices_Handler,
		},
		{
			MethodName: "RemoveOwnerDevices",
			Handler:    _Devices_RemoveOwnerDevices_Handler,
		},
		{
			MethodName: "GetDeviceStatus",
			Handler:    _Devices_GetDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/registry/registrypb/device_registry.proto",
}
