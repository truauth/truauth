// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package grpcIdentity

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ClientIdentityRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientIdentityRequest) Reset()         { *m = ClientIdentityRequest{} }
func (m *ClientIdentityRequest) String() string { return proto.CompactTextString(m) }
func (*ClientIdentityRequest) ProtoMessage()    {}
func (*ClientIdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ClientIdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientIdentityRequest.Unmarshal(m, b)
}
func (m *ClientIdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientIdentityRequest.Marshal(b, m, deterministic)
}
func (m *ClientIdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientIdentityRequest.Merge(m, src)
}
func (m *ClientIdentityRequest) XXX_Size() int {
	return xxx_messageInfo_ClientIdentityRequest.Size(m)
}
func (m *ClientIdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientIdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientIdentityRequest proto.InternalMessageInfo

func (m *ClientIdentityRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ClientIdentityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClientIdentity struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Secret               string   `protobuf:"bytes,3,opt,name=Secret,proto3" json:"Secret,omitempty"`
	Callback             string   `protobuf:"bytes,4,opt,name=Callback,proto3" json:"Callback,omitempty"`
	Redirect             string   `protobuf:"bytes,5,opt,name=Redirect,proto3" json:"Redirect,omitempty"`
	RelationID           string   `protobuf:"bytes,6,opt,name=RelationID,proto3" json:"RelationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientIdentity) Reset()         { *m = ClientIdentity{} }
func (m *ClientIdentity) String() string { return proto.CompactTextString(m) }
func (*ClientIdentity) ProtoMessage()    {}
func (*ClientIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ClientIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientIdentity.Unmarshal(m, b)
}
func (m *ClientIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientIdentity.Marshal(b, m, deterministic)
}
func (m *ClientIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientIdentity.Merge(m, src)
}
func (m *ClientIdentity) XXX_Size() int {
	return xxx_messageInfo_ClientIdentity.Size(m)
}
func (m *ClientIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_ClientIdentity proto.InternalMessageInfo

func (m *ClientIdentity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ClientIdentity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClientIdentity) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ClientIdentity) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *ClientIdentity) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *ClientIdentity) GetRelationID() string {
	if m != nil {
		return m.RelationID
	}
	return ""
}

type UserIdentityRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdentityRequest) Reset()         { *m = UserIdentityRequest{} }
func (m *UserIdentityRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdentityRequest) ProtoMessage()    {}
func (*UserIdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *UserIdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdentityRequest.Unmarshal(m, b)
}
func (m *UserIdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdentityRequest.Marshal(b, m, deterministic)
}
func (m *UserIdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdentityRequest.Merge(m, src)
}
func (m *UserIdentityRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdentityRequest.Size(m)
}
func (m *UserIdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdentityRequest proto.InternalMessageInfo

func (m *UserIdentityRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserIdentityRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserIdentityRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SafeUserIdentity struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SafeUserIdentity) Reset()         { *m = SafeUserIdentity{} }
func (m *SafeUserIdentity) String() string { return proto.CompactTextString(m) }
func (*SafeUserIdentity) ProtoMessage()    {}
func (*SafeUserIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *SafeUserIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SafeUserIdentity.Unmarshal(m, b)
}
func (m *SafeUserIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SafeUserIdentity.Marshal(b, m, deterministic)
}
func (m *SafeUserIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SafeUserIdentity.Merge(m, src)
}
func (m *SafeUserIdentity) XXX_Size() int {
	return xxx_messageInfo_SafeUserIdentity.Size(m)
}
func (m *SafeUserIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_SafeUserIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_SafeUserIdentity proto.InternalMessageInfo

func (m *SafeUserIdentity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SafeUserIdentity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SafeUserIdentity) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type UnsafeUserIdentity struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Type                 int32    `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsafeUserIdentity) Reset()         { *m = UnsafeUserIdentity{} }
func (m *UnsafeUserIdentity) String() string { return proto.CompactTextString(m) }
func (*UnsafeUserIdentity) ProtoMessage()    {}
func (*UnsafeUserIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *UnsafeUserIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsafeUserIdentity.Unmarshal(m, b)
}
func (m *UnsafeUserIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsafeUserIdentity.Marshal(b, m, deterministic)
}
func (m *UnsafeUserIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsafeUserIdentity.Merge(m, src)
}
func (m *UnsafeUserIdentity) XXX_Size() int {
	return xxx_messageInfo_UnsafeUserIdentity.Size(m)
}
func (m *UnsafeUserIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsafeUserIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_UnsafeUserIdentity proto.InternalMessageInfo

func (m *UnsafeUserIdentity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UnsafeUserIdentity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UnsafeUserIdentity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UnsafeUserIdentity) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type CreatedResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatedResponse) Reset()         { *m = CreatedResponse{} }
func (m *CreatedResponse) String() string { return proto.CompactTextString(m) }
func (*CreatedResponse) ProtoMessage()    {}
func (*CreatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *CreatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatedResponse.Unmarshal(m, b)
}
func (m *CreatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatedResponse.Marshal(b, m, deterministic)
}
func (m *CreatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatedResponse.Merge(m, src)
}
func (m *CreatedResponse) XXX_Size() int {
	return xxx_messageInfo_CreatedResponse.Size(m)
}
func (m *CreatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatedResponse proto.InternalMessageInfo

func (m *CreatedResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ClientIdentityRequest)(nil), "grpcIdentity.ClientIdentityRequest")
	proto.RegisterType((*ClientIdentity)(nil), "grpcIdentity.ClientIdentity")
	proto.RegisterType((*UserIdentityRequest)(nil), "grpcIdentity.UserIdentityRequest")
	proto.RegisterType((*SafeUserIdentity)(nil), "grpcIdentity.SafeUserIdentity")
	proto.RegisterType((*UnsafeUserIdentity)(nil), "grpcIdentity.UnsafeUserIdentity")
	proto.RegisterType((*CreatedResponse)(nil), "grpcIdentity.CreatedResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcb, 0xee, 0xd2, 0x40,
	0x14, 0xc6, 0xa1, 0x16, 0xac, 0x27, 0x8a, 0x66, 0xb8, 0xa4, 0x69, 0x94, 0xe0, 0xb8, 0x31, 0x31,
	0x61, 0xa1, 0x4b, 0x97, 0xc0, 0xa2, 0x1b, 0x63, 0x06, 0x89, 0x97, 0xb8, 0x19, 0xa6, 0x47, 0x32,
	0xb1, 0xb4, 0x65, 0x66, 0xc0, 0xf0, 0x3e, 0x3e, 0xa4, 0x4b, 0xd3, 0xa1, 0x6d, 0xda, 0x4a, 0xd0,
	0xf8, 0xdf, 0xcd, 0x77, 0x6e, 0xdf, 0x6f, 0x6e, 0xf0, 0x48, 0xa3, 0x3a, 0x49, 0x81, 0xf3, 0x4c,
	0xa5, 0x26, 0x25, 0x0f, 0x77, 0x2a, 0x13, 0x61, 0x84, 0x89, 0x91, 0xe6, 0x4c, 0xdf, 0xc2, 0x78,
	0x11, 0x4b, 0x4c, 0x4c, 0x19, 0x61, 0x78, 0x38, 0xa2, 0x36, 0x64, 0x00, 0x4e, 0xb8, 0xf4, 0xbb,
	0xb3, 0xee, 0xcb, 0x07, 0xcc, 0x09, 0x97, 0x84, 0x80, 0xfb, 0x8e, 0xef, 0xd1, 0x77, 0x6c, 0xc4,
	0xae, 0xe9, 0xcf, 0x2e, 0x0c, 0x9a, 0xdd, 0xff, 0xd2, 0x46, 0x26, 0xd0, 0x5f, 0xa3, 0x50, 0x68,
	0xfc, 0x7b, 0x36, 0x5a, 0x28, 0x12, 0x80, 0xb7, 0xe0, 0x71, 0xbc, 0xe5, 0xe2, 0xbb, 0xef, 0xda,
	0x4c, 0xa5, 0xf3, 0x1c, 0xc3, 0x48, 0x2a, 0x14, 0xc6, 0xef, 0x5d, 0x72, 0xa5, 0x26, 0x53, 0x00,
	0x86, 0x31, 0x37, 0x32, 0x4d, 0xc2, 0xa5, 0xdf, 0xb7, 0xd9, 0x5a, 0x84, 0x0a, 0x18, 0x6e, 0x34,
	0xaa, 0xf6, 0x0e, 0x03, 0xf0, 0xf2, 0x70, 0x92, 0xe3, 0x5d, 0x80, 0x2b, 0x4d, 0x46, 0xd0, 0x5b,
	0xed, 0xb9, 0x8c, 0x0b, 0xee, 0x8b, 0xc8, 0x3b, 0xde, 0x73, 0xad, 0x7f, 0xa4, 0x2a, 0x2a, 0xd0,
	0x2b, 0x4d, 0x3f, 0xc1, 0x93, 0x35, 0xff, 0x86, 0x75, 0xa3, 0xff, 0x70, 0x20, 0xe0, 0x7e, 0x38,
	0x67, 0x68, 0xa7, 0xf7, 0x98, 0x5d, 0xd3, 0x13, 0x90, 0x4d, 0xa2, 0xef, 0x3e, 0xfb, 0x06, 0x7d,
	0xe5, 0xeb, 0xd6, 0x7c, 0x5f, 0xc1, 0xe3, 0x85, 0x42, 0x6e, 0x30, 0x62, 0xa8, 0xb3, 0x34, 0xd1,
	0x48, 0x7c, 0xb8, 0xbf, 0x3e, 0x0a, 0x81, 0x5a, 0x5b, 0x4f, 0x8f, 0x95, 0xf2, 0xf5, 0x2f, 0x07,
	0xbc, 0x8a, 0xed, 0x0b, 0x0c, 0x57, 0xc9, 0xe1, 0x28, 0x55, 0x13, 0xf9, 0xf9, 0xbc, 0xfe, 0xf4,
	0xe6, 0x57, 0xee, 0x24, 0x98, 0x36, 0x4b, 0xda, 0x27, 0x4a, 0x3b, 0xe4, 0x33, 0x8c, 0x18, 0xee,
	0xa4, 0x36, 0xa8, 0x1a, 0xc3, 0x67, 0xad, 0xe1, 0x7f, 0x9c, 0x58, 0xf0, 0xac, 0x59, 0xd1, 0xda,
	0x1b, 0xed, 0x90, 0xaf, 0x30, 0x2e, 0xb0, 0x5b, 0x8f, 0xfa, 0x45, 0xab, 0xf3, 0xda, 0x87, 0x09,
	0x9e, 0xde, 0x2a, 0xa2, 0x1d, 0xf2, 0x11, 0x26, 0x25, 0x78, 0x6b, 0xfc, 0xcd, 0xce, 0xbf, 0x62,
	0x6f, 0xfb, 0xf6, 0x5f, 0xbf, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xa5, 0xea, 0x79, 0xe8,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	EnquireUserIdentity(ctx context.Context, in *UserIdentityRequest, opts ...grpc.CallOption) (*SafeUserIdentity, error)
	RegisterUserIdentity(ctx context.Context, in *UnsafeUserIdentity, opts ...grpc.CallOption) (*CreatedResponse, error)
	EnquireClientIdentity(ctx context.Context, in *ClientIdentityRequest, opts ...grpc.CallOption) (*ClientIdentity, error)
	RegisterClientIdentity(ctx context.Context, in *ClientIdentity, opts ...grpc.CallOption) (*CreatedResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) EnquireUserIdentity(ctx context.Context, in *UserIdentityRequest, opts ...grpc.CallOption) (*SafeUserIdentity, error) {
	out := new(SafeUserIdentity)
	err := c.cc.Invoke(ctx, "/grpcIdentity.Identity/EnquireUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) RegisterUserIdentity(ctx context.Context, in *UnsafeUserIdentity, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/grpcIdentity.Identity/RegisterUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) EnquireClientIdentity(ctx context.Context, in *ClientIdentityRequest, opts ...grpc.CallOption) (*ClientIdentity, error) {
	out := new(ClientIdentity)
	err := c.cc.Invoke(ctx, "/grpcIdentity.Identity/EnquireClientIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) RegisterClientIdentity(ctx context.Context, in *ClientIdentity, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/grpcIdentity.Identity/RegisterClientIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	EnquireUserIdentity(context.Context, *UserIdentityRequest) (*SafeUserIdentity, error)
	RegisterUserIdentity(context.Context, *UnsafeUserIdentity) (*CreatedResponse, error)
	EnquireClientIdentity(context.Context, *ClientIdentityRequest) (*ClientIdentity, error)
	RegisterClientIdentity(context.Context, *ClientIdentity) (*CreatedResponse, error)
}

// UnimplementedIdentityServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (*UnimplementedIdentityServer) EnquireUserIdentity(ctx context.Context, req *UserIdentityRequest) (*SafeUserIdentity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnquireUserIdentity not implemented")
}
func (*UnimplementedIdentityServer) RegisterUserIdentity(ctx context.Context, req *UnsafeUserIdentity) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserIdentity not implemented")
}
func (*UnimplementedIdentityServer) EnquireClientIdentity(ctx context.Context, req *ClientIdentityRequest) (*ClientIdentity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnquireClientIdentity not implemented")
}
func (*UnimplementedIdentityServer) RegisterClientIdentity(ctx context.Context, req *ClientIdentity) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClientIdentity not implemented")
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_EnquireUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).EnquireUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcIdentity.Identity/EnquireUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).EnquireUserIdentity(ctx, req.(*UserIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_RegisterUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsafeUserIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).RegisterUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcIdentity.Identity/RegisterUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).RegisterUserIdentity(ctx, req.(*UnsafeUserIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_EnquireClientIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).EnquireClientIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcIdentity.Identity/EnquireClientIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).EnquireClientIdentity(ctx, req.(*ClientIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_RegisterClientIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).RegisterClientIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcIdentity.Identity/RegisterClientIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).RegisterClientIdentity(ctx, req.(*ClientIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcIdentity.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnquireUserIdentity",
			Handler:    _Identity_EnquireUserIdentity_Handler,
		},
		{
			MethodName: "RegisterUserIdentity",
			Handler:    _Identity_RegisterUserIdentity_Handler,
		},
		{
			MethodName: "EnquireClientIdentity",
			Handler:    _Identity_EnquireClientIdentity_Handler,
		},
		{
			MethodName: "RegisterClientIdentity",
			Handler:    _Identity_RegisterClientIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}