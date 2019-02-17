// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/node/nodepb/namespace.proto

package nodepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Namespace struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{0}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateNamespaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceRequest) Reset()         { *m = CreateNamespaceRequest{} }
func (m *CreateNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceRequest) ProtoMessage()    {}
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{1}
}

func (m *CreateNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceRequest.Unmarshal(m, b)
}
func (m *CreateNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceRequest.Merge(m, src)
}
func (m *CreateNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceRequest.Size(m)
}
func (m *CreateNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceRequest proto.InternalMessageInfo

func (m *CreateNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateNamespaceResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceResponse) Reset()         { *m = CreateNamespaceResponse{} }
func (m *CreateNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceResponse) ProtoMessage()    {}
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{2}
}

func (m *CreateNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceResponse.Unmarshal(m, b)
}
func (m *CreateNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceResponse.Merge(m, src)
}
func (m *CreateNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceResponse.Size(m)
}
func (m *CreateNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceResponse proto.InternalMessageInfo

func (m *CreateNamespaceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetNamespaceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNamespaceRequest) Reset()         { *m = GetNamespaceRequest{} }
func (m *GetNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetNamespaceRequest) ProtoMessage()    {}
func (*GetNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{3}
}

func (m *GetNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNamespaceRequest.Unmarshal(m, b)
}
func (m *GetNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *GetNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNamespaceRequest.Merge(m, src)
}
func (m *GetNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_GetNamespaceRequest.Size(m)
}
func (m *GetNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNamespaceRequest proto.InternalMessageInfo

func (m *GetNamespaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetNamespaceResponse struct {
	Namespace            *Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetNamespaceResponse) Reset()         { *m = GetNamespaceResponse{} }
func (m *GetNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*GetNamespaceResponse) ProtoMessage()    {}
func (*GetNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{4}
}

func (m *GetNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNamespaceResponse.Unmarshal(m, b)
}
func (m *GetNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *GetNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNamespaceResponse.Merge(m, src)
}
func (m *GetNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_GetNamespaceResponse.Size(m)
}
func (m *GetNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNamespaceResponse proto.InternalMessageInfo

func (m *GetNamespaceResponse) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type ListNamespacesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNamespacesRequest) Reset()         { *m = ListNamespacesRequest{} }
func (m *ListNamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesRequest) ProtoMessage()    {}
func (*ListNamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{5}
}

func (m *ListNamespacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesRequest.Unmarshal(m, b)
}
func (m *ListNamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesRequest.Marshal(b, m, deterministic)
}
func (m *ListNamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesRequest.Merge(m, src)
}
func (m *ListNamespacesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesRequest.Size(m)
}
func (m *ListNamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesRequest proto.InternalMessageInfo

type ListNamespacesForAccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNamespacesForAccountRequest) Reset()         { *m = ListNamespacesForAccountRequest{} }
func (m *ListNamespacesForAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesForAccountRequest) ProtoMessage()    {}
func (*ListNamespacesForAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{6}
}

func (m *ListNamespacesForAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Unmarshal(m, b)
}
func (m *ListNamespacesForAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Marshal(b, m, deterministic)
}
func (m *ListNamespacesForAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesForAccountRequest.Merge(m, src)
}
func (m *ListNamespacesForAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Size(m)
}
func (m *ListNamespacesForAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesForAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesForAccountRequest proto.InternalMessageInfo

func (m *ListNamespacesForAccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type ListNamespacesResponse struct {
	Namespaces           []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListNamespacesResponse) Reset()         { *m = ListNamespacesResponse{} }
func (m *ListNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesResponse) ProtoMessage()    {}
func (*ListNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{7}
}

func (m *ListNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesResponse.Unmarshal(m, b)
}
func (m *ListNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *ListNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesResponse.Merge(m, src)
}
func (m *ListNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesResponse.Size(m)
}
func (m *ListNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesResponse proto.InternalMessageInfo

func (m *ListNamespacesResponse) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Namespace)(nil), "infinimesh.node.Namespace")
	proto.RegisterType((*CreateNamespaceRequest)(nil), "infinimesh.node.CreateNamespaceRequest")
	proto.RegisterType((*CreateNamespaceResponse)(nil), "infinimesh.node.CreateNamespaceResponse")
	proto.RegisterType((*GetNamespaceRequest)(nil), "infinimesh.node.GetNamespaceRequest")
	proto.RegisterType((*GetNamespaceResponse)(nil), "infinimesh.node.GetNamespaceResponse")
	proto.RegisterType((*ListNamespacesRequest)(nil), "infinimesh.node.ListNamespacesRequest")
	proto.RegisterType((*ListNamespacesForAccountRequest)(nil), "infinimesh.node.ListNamespacesForAccountRequest")
	proto.RegisterType((*ListNamespacesResponse)(nil), "infinimesh.node.ListNamespacesResponse")
}

func init() { proto.RegisterFile("pkg/node/nodepb/namespace.proto", fileDescriptor_2ba94f8fc77af896) }

var fileDescriptor_2ba94f8fc77af896 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x4f, 0xc2, 0x30,
	0x14, 0xc5, 0x33, 0x30, 0x18, 0xae, 0x06, 0x4c, 0x51, 0x58, 0xf6, 0x02, 0x69, 0x44, 0x30, 0x31,
	0x9b, 0xc1, 0x17, 0xa3, 0x4f, 0x62, 0xa2, 0x2f, 0xc6, 0x10, 0xf4, 0x49, 0x9f, 0xc6, 0x56, 0xb5,
	0x31, 0xb4, 0x73, 0x1d, 0x7e, 0x2d, 0xbf, 0xa2, 0xd9, 0xe8, 0xba, 0x3f, 0x9d, 0x99, 0x2f, 0x84,
	0xb5, 0xbf, 0x7b, 0xee, 0xe9, 0x3d, 0xb9, 0x30, 0x0c, 0x3e, 0xdf, 0x1d, 0xc6, 0x7d, 0x92, 0xfc,
	0x04, 0x2b, 0x87, 0xb9, 0x6b, 0x22, 0x02, 0xd7, 0x23, 0x76, 0x10, 0xf2, 0x88, 0xa3, 0x2e, 0x65,
	0x6f, 0x94, 0xd1, 0x35, 0x11, 0x1f, 0x76, 0x8c, 0x60, 0x07, 0xda, 0x8f, 0x29, 0x83, 0x3a, 0xd0,
	0xa0, 0xbe, 0x69, 0x8c, 0x8c, 0x69, 0x7b, 0xd9, 0xa0, 0x3e, 0x42, 0xb0, 0x13, 0x0b, 0x98, 0x8d,
	0xe4, 0x24, 0xf9, 0x8f, 0xcf, 0xa0, 0x7f, 0x1b, 0x12, 0x37, 0x22, 0xaa, 0x6c, 0x49, 0xbe, 0x36,
	0x44, 0x44, 0x8a, 0x36, 0x72, 0xf4, 0x29, 0x0c, 0x34, 0x5a, 0x04, 0x9c, 0x09, 0xad, 0x19, 0x1e,
	0x43, 0xef, 0x9e, 0x44, 0x9a, 0x6a, 0x19, 0x5b, 0xc0, 0x61, 0x11, 0x93, 0x72, 0x97, 0xd0, 0x56,
	0x8f, 0x4d, 0xf0, 0xbd, 0x99, 0x65, 0x97, 0x5e, 0x6b, 0x67, 0x65, 0x19, 0x8c, 0x07, 0x70, 0xf4,
	0x40, 0x45, 0x26, 0x29, 0x64, 0x6b, 0x7c, 0x0d, 0xc3, 0xe2, 0xc5, 0x1d, 0x0f, 0x6f, 0x3c, 0x8f,
	0x6f, 0x58, 0x94, 0xba, 0x33, 0x61, 0xd7, 0xdd, 0x9e, 0x48, 0x8b, 0xe9, 0x27, 0x7e, 0x86, 0x7e,
	0x59, 0x55, 0x3a, 0xbd, 0x02, 0x50, 0xcd, 0x85, 0x69, 0x8c, 0x9a, 0x35, 0x56, 0x73, 0xf4, 0xec,
	0xa7, 0x09, 0x07, 0xea, 0xe6, 0x89, 0x84, 0xdf, 0xd4, 0x23, 0xc8, 0x87, 0x6e, 0x69, 0xc8, 0x68,
	0xa2, 0xe9, 0x55, 0x87, 0x66, 0x4d, 0xeb, 0x41, 0x69, 0xfb, 0x15, 0xf6, 0xf3, 0x83, 0x47, 0xc7,
	0x5a, 0x65, 0x45, 0x7c, 0xd6, 0xb8, 0x86, 0x92, 0xe2, 0x2e, 0x74, 0x8a, 0xd3, 0x42, 0x27, 0x5a,
	0x61, 0x65, 0x48, 0xd6, 0xa4, 0x96, 0x93, 0x2d, 0x36, 0x60, 0xfe, 0x95, 0x26, 0x3a, 0xaf, 0x11,
	0xd1, 0x82, 0xff, 0x77, 0xdb, 0xb9, 0x03, 0x3d, 0x8f, 0xaf, 0xcb, 0xf4, 0xbc, 0xa3, 0xd0, 0x45,
	0xbc, 0x98, 0x0b, 0xe3, 0xa5, 0xb5, 0x5d, 0xd9, 0x55, 0x2b, 0xd9, 0xd4, 0x8b, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb1, 0xde, 0x46, 0xca, 0xcc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	ListNamespacesForAccount(ctx context.Context, in *ListNamespacesForAccountRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type namespaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceServiceClient(cc *grpc.ClientConn) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.NamespaceService/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.NamespaceService/GetNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.NamespaceService/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) ListNamespacesForAccount(ctx context.Context, in *ListNamespacesForAccountRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.NamespaceService/ListNamespacesForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
type NamespaceServiceServer interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	ListNamespacesForAccount(context.Context, *ListNamespacesForAccountRequest) (*ListNamespacesResponse, error)
}

func RegisterNamespaceServiceServer(s *grpc.Server, srv NamespaceServiceServer) {
	s.RegisterService(&_NamespaceService_serviceDesc, srv)
}

func _NamespaceService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.NamespaceService/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.NamespaceService/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.NamespaceService/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_ListNamespacesForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesForAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).ListNamespacesForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.NamespaceService/ListNamespacesForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).ListNamespacesForAccount(ctx, req.(*ListNamespacesForAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.node.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _NamespaceService_CreateNamespace_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _NamespaceService_GetNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _NamespaceService_ListNamespaces_Handler,
		},
		{
			MethodName: "ListNamespacesForAccount",
			Handler:    _NamespaceService_ListNamespacesForAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/node/nodepb/namespace.proto",
}