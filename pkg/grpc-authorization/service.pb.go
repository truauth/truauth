// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package grpcAuthorization

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

type AuthorizeUserRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ClientID             string   `protobuf:"bytes,2,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeUserRequest) Reset()         { *m = AuthorizeUserRequest{} }
func (m *AuthorizeUserRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeUserRequest) ProtoMessage()    {}
func (*AuthorizeUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *AuthorizeUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeUserRequest.Unmarshal(m, b)
}
func (m *AuthorizeUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeUserRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizeUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeUserRequest.Merge(m, src)
}
func (m *AuthorizeUserRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeUserRequest.Size(m)
}
func (m *AuthorizeUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeUserRequest proto.InternalMessageInfo

func (m *AuthorizeUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AuthorizeUserRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type SuccessResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessResponse) Reset()         { *m = SuccessResponse{} }
func (m *SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse) ProtoMessage()    {}
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse.Unmarshal(m, b)
}
func (m *SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse.Merge(m, src)
}
func (m *SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse.Size(m)
}
func (m *SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse proto.InternalMessageInfo

func (m *SuccessResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*AuthorizeUserRequest)(nil), "grpcAuthorization.AuthorizeUserRequest")
	proto.RegisterType((*SuccessResponse)(nil), "grpcAuthorization.SuccessResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0x2f, 0x2a, 0x48, 0x76,
	0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x53, 0xf2, 0xe2, 0x12,
	0x81, 0x09, 0xa4, 0x86, 0x16, 0xa7, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89,
	0x71, 0xb1, 0x81, 0xb8, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90,
	0x14, 0x17, 0x87, 0x73, 0x4e, 0x66, 0x6a, 0x5e, 0x89, 0xa7, 0x8b, 0x04, 0x13, 0x58, 0x06, 0xce,
	0x57, 0xd2, 0xe6, 0xe2, 0x0f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x0e, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe0, 0x62, 0x87, 0x0a, 0x81, 0xcd, 0xe1, 0x08, 0x82, 0x71, 0x8d,
	0x2e, 0x31, 0x72, 0xf1, 0xa2, 0x38, 0x45, 0x28, 0x0e, 0x21, 0x00, 0x76, 0x8a, 0x90, 0xba, 0x1e,
	0x86, 0x7b, 0xf5, 0xb0, 0x39, 0x56, 0x4a, 0x09, 0x8b, 0x42, 0x34, 0x97, 0x28, 0x31, 0x08, 0x25,
	0x72, 0x09, 0x78, 0x16, 0x83, 0xb4, 0xc1, 0xcd, 0x48, 0xa1, 0xb2, 0x15, 0x49, 0x6c, 0xe0, 0x70,
	0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xbd, 0x9a, 0x06, 0x78, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	AuthorizeUser(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	IsUserAuthorized(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) AuthorizeUser(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/grpcAuthorization.Authorization/AuthorizeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IsUserAuthorized(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/grpcAuthorization.Authorization/IsUserAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	AuthorizeUser(context.Context, *AuthorizeUserRequest) (*SuccessResponse, error)
	IsUserAuthorized(context.Context, *AuthorizeUserRequest) (*SuccessResponse, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) AuthorizeUser(ctx context.Context, req *AuthorizeUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeUser not implemented")
}
func (*UnimplementedAuthorizationServer) IsUserAuthorized(ctx context.Context, req *AuthorizeUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserAuthorized not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcAuthorization.Authorization/AuthorizeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).AuthorizeUser(ctx, req.(*AuthorizeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IsUserAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsUserAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcAuthorization.Authorization/IsUserAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsUserAuthorized(ctx, req.(*AuthorizeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcAuthorization.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeUser",
			Handler:    _Authorization_AuthorizeUser_Handler,
		},
		{
			MethodName: "IsUserAuthorized",
			Handler:    _Authorization_IsUserAuthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
