// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.tl.proto

package auth

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TLConstructor int32

const (
	CRC32_UNKNOWN               TLConstructor = 0
	CRC32_auth_exportLoginToken TLConstructor = -1210022402
	CRC32_auth_importLoginToken TLConstructor = -1783866140
	CRC32_auth_acceptLoginToken TLConstructor = -392909491
)

var TLConstructor_name = map[int32]string{
	0:           "CRC32_UNKNOWN",
	-1210022402: "CRC32_auth_exportLoginToken",
	-1783866140: "CRC32_auth_importLoginToken",
	-392909491:  "CRC32_auth_acceptLoginToken",
}

var TLConstructor_value = map[string]int32{
	"CRC32_UNKNOWN":               0,
	"CRC32_auth_exportLoginToken": -1210022402,
	"CRC32_auth_importLoginToken": -1783866140,
	"CRC32_auth_acceptLoginToken": -392909491,
}

func (x TLConstructor) String() string {
	return proto.EnumName(TLConstructor_name, int32(x))
}

func (TLConstructor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d91d231cbf9bbae, []int{0}
}

//--------------------------------------------------------------------------------------------
// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
type TLAuthExportLoginToken struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=auth.TLConstructor" json:"constructor,omitempty"`
	ApiId                int32         `protobuf:"varint,3,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	ApiHash              string        `protobuf:"bytes,4,opt,name=api_hash,json=apiHash,proto3" json:"api_hash,omitempty"`
	ExceptIds            []int64       `protobuf:"varint,5,rep,packed,name=except_ids,json=exceptIds,proto3" json:"except_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLAuthExportLoginToken) Reset()         { *m = TLAuthExportLoginToken{} }
func (m *TLAuthExportLoginToken) String() string { return proto.CompactTextString(m) }
func (*TLAuthExportLoginToken) ProtoMessage()    {}
func (*TLAuthExportLoginToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d91d231cbf9bbae, []int{0}
}
func (m *TLAuthExportLoginToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLAuthExportLoginToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLAuthExportLoginToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLAuthExportLoginToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLAuthExportLoginToken.Merge(m, src)
}
func (m *TLAuthExportLoginToken) XXX_Size() int {
	return m.Size()
}
func (m *TLAuthExportLoginToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TLAuthExportLoginToken.DiscardUnknown(m)
}

var xxx_messageInfo_TLAuthExportLoginToken proto.InternalMessageInfo

func (m *TLAuthExportLoginToken) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLAuthExportLoginToken) GetApiId() int32 {
	if m != nil {
		return m.ApiId
	}
	return 0
}

func (m *TLAuthExportLoginToken) GetApiHash() string {
	if m != nil {
		return m.ApiHash
	}
	return ""
}

func (m *TLAuthExportLoginToken) GetExceptIds() []int64 {
	if m != nil {
		return m.ExceptIds
	}
	return nil
}

//--------------------------------------------------------------------------------------------
// auth.importLoginToken token:bytes = auth.LoginToken;
type TLAuthImportLoginToken struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=auth.TLConstructor" json:"constructor,omitempty"`
	Token                []byte        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLAuthImportLoginToken) Reset()         { *m = TLAuthImportLoginToken{} }
func (m *TLAuthImportLoginToken) String() string { return proto.CompactTextString(m) }
func (*TLAuthImportLoginToken) ProtoMessage()    {}
func (*TLAuthImportLoginToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d91d231cbf9bbae, []int{1}
}
func (m *TLAuthImportLoginToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLAuthImportLoginToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLAuthImportLoginToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLAuthImportLoginToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLAuthImportLoginToken.Merge(m, src)
}
func (m *TLAuthImportLoginToken) XXX_Size() int {
	return m.Size()
}
func (m *TLAuthImportLoginToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TLAuthImportLoginToken.DiscardUnknown(m)
}

var xxx_messageInfo_TLAuthImportLoginToken proto.InternalMessageInfo

func (m *TLAuthImportLoginToken) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLAuthImportLoginToken) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

//--------------------------------------------------------------------------------------------
// auth.acceptLoginToken token:bytes = Authorization;
type TLAuthAcceptLoginToken struct {
	Constructor          TLConstructor `protobuf:"varint,1,opt,name=constructor,proto3,enum=auth.TLConstructor" json:"constructor,omitempty"`
	Token                []byte        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TLAuthAcceptLoginToken) Reset()         { *m = TLAuthAcceptLoginToken{} }
func (m *TLAuthAcceptLoginToken) String() string { return proto.CompactTextString(m) }
func (*TLAuthAcceptLoginToken) ProtoMessage()    {}
func (*TLAuthAcceptLoginToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d91d231cbf9bbae, []int{2}
}
func (m *TLAuthAcceptLoginToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLAuthAcceptLoginToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLAuthAcceptLoginToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLAuthAcceptLoginToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLAuthAcceptLoginToken.Merge(m, src)
}
func (m *TLAuthAcceptLoginToken) XXX_Size() int {
	return m.Size()
}
func (m *TLAuthAcceptLoginToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TLAuthAcceptLoginToken.DiscardUnknown(m)
}

var xxx_messageInfo_TLAuthAcceptLoginToken proto.InternalMessageInfo

func (m *TLAuthAcceptLoginToken) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return CRC32_UNKNOWN
}

func (m *TLAuthAcceptLoginToken) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterEnum("auth.TLConstructor", TLConstructor_name, TLConstructor_value)
	proto.RegisterType((*TLAuthExportLoginToken)(nil), "auth.TL_auth_exportLoginToken")
	proto.RegisterType((*TLAuthImportLoginToken)(nil), "auth.TL_auth_importLoginToken")
	proto.RegisterType((*TLAuthAcceptLoginToken)(nil), "auth.TL_auth_acceptLoginToken")
}

func init() { proto.RegisterFile("auth.tl.proto", fileDescriptor_5d91d231cbf9bbae) }

var fileDescriptor_5d91d231cbf9bbae = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x91, 0xa4, 0xa5, 0x07, 0x41, 0xc1, 0xb4, 0xc8, 0x35, 0xc2, 0x8a, 0x32, 0x45, 0x48,
	0xb5, 0xa5, 0x54, 0x8c, 0x0c, 0x90, 0x85, 0x8a, 0x28, 0x05, 0x13, 0x84, 0xc4, 0x62, 0x5d, 0x2e,
	0x87, 0x7d, 0xa2, 0xf6, 0x9d, 0xee, 0xce, 0xd0, 0x76, 0x64, 0x40, 0x48, 0xfc, 0x0b, 0x90, 0xd8,
	0x18, 0x18, 0x99, 0x58, 0x18, 0x58, 0x90, 0x10, 0x23, 0x13, 0x44, 0xfc, 0x00, 0x56, 0x84, 0xd4,
	0xa0, 0x3b, 0x27, 0x75, 0x71, 0x23, 0x26, 0xba, 0x58, 0xef, 0xbd, 0xef, 0x7b, 0xdf, 0xf7, 0x9e,
	0xdf, 0xc1, 0x06, 0xca, 0x54, 0xec, 0xa9, 0x1d, 0x8f, 0x0b, 0xa6, 0x98, 0x55, 0xd3, 0xa9, 0xb3,
	0x11, 0x51, 0x15, 0x67, 0x23, 0x0f, 0xb3, 0xc4, 0x8f, 0x58, 0xc4, 0x7c, 0x03, 0x8e, 0xb2, 0x87,
	0x26, 0x33, 0x89, 0x89, 0xf2, 0x26, 0xc7, 0x8d, 0x18, 0x8b, 0x76, 0x48, 0xc1, 0x7a, 0x22, 0x10,
	0xe7, 0x44, 0xc8, 0x19, 0xee, 0x48, 0x1c, 0x93, 0x04, 0x69, 0x17, 0xcc, 0x04, 0x09, 0xd5, 0x1e,
	0x27, 0x73, 0x6c, 0xbd, 0xc0, 0x94, 0x40, 0xa9, 0xe4, 0x4c, 0xa8, 0x19, 0xb4, 0x5a, 0x40, 0x72,
	0x2f, 0xc5, 0x79, 0xb5, 0xfd, 0x12, 0x40, 0x7b, 0xd8, 0x0f, 0xf5, 0x9c, 0x21, 0xd9, 0xd5, 0xf4,
	0x3e, 0x8b, 0x68, 0x3a, 0x64, 0x8f, 0x48, 0x6a, 0x5d, 0x85, 0x67, 0x30, 0x4b, 0xa5, 0x12, 0x19,
	0x56, 0x4c, 0xd8, 0xa0, 0x05, 0x3a, 0xe7, 0xba, 0x17, 0x3c, 0xb3, 0xe3, 0xb0, 0xdf, 0x2b, 0xa0,
	0xe0, 0x28, 0xcf, 0x5a, 0x83, 0x4b, 0x88, 0xd3, 0x90, 0x8e, 0xed, 0x6a, 0x0b, 0x74, 0xea, 0x41,
	0x1d, 0x71, 0xba, 0x35, 0xb6, 0xd6, 0xe1, 0x69, 0x5d, 0x8e, 0x91, 0x8c, 0xed, 0x5a, 0x0b, 0x74,
	0x56, 0x82, 0x65, 0xc4, 0xe9, 0x4d, 0x24, 0x63, 0xeb, 0x32, 0x84, 0x64, 0x17, 0x13, 0xae, 0x42,
	0x3a, 0x96, 0x76, 0xbd, 0x55, 0xed, 0x54, 0x83, 0x95, 0xbc, 0xb2, 0x35, 0x96, 0xed, 0xa8, 0x98,
	0x91, 0x26, 0xff, 0x67, 0xc6, 0x55, 0x58, 0x57, 0xba, 0xdf, 0x8c, 0x78, 0x36, 0xc8, 0x93, 0xa3,
	0x46, 0x08, 0x6b, 0xf7, 0x13, 0x32, 0xba, 0xf2, 0x1a, 0xc0, 0xc6, 0x5f, 0x4d, 0xd6, 0x79, 0xd8,
	0xe8, 0x05, 0xbd, 0xcd, 0x6e, 0x78, 0x6f, 0x70, 0x6b, 0xb0, 0x7d, 0x7f, 0xd0, 0xac, 0x58, 0x1d,
	0x78, 0x29, 0x2f, 0x2d, 0xbc, 0x4e, 0xf3, 0xe0, 0xc5, 0xd3, 0xf7, 0xbf, 0xa7, 0xd3, 0xe9, 0x14,
	0x94, 0x98, 0xe5, 0x7f, 0xd4, 0xfc, 0xf1, 0xee, 0xed, 0x9b, 0x5f, 0x8b, 0x98, 0xe5, 0x25, 0x9b,
	0x9f, 0xbe, 0x7e, 0xf9, 0x70, 0x60, 0x98, 0x4e, 0xed, 0xf9, 0x2b, 0xb7, 0xd2, 0x7d, 0x76, 0x0a,
	0x2e, 0x07, 0xb7, 0x7b, 0xd7, 0x33, 0x15, 0x5b, 0x77, 0xe1, 0xda, 0xe2, 0x77, 0xe2, 0xce, 0xff,
	0xc2, 0xe2, 0x49, 0x1d, 0xdb, 0x4b, 0x94, 0x79, 0x6e, 0x86, 0x17, 0x16, 0x48, 0xbb, 0x72, 0x28,
	0x7a, 0xec, 0xb0, 0x25, 0xd1, 0x32, 0xfe, 0x4f, 0xd1, 0x3b, 0x33, 0xd1, 0x63, 0x47, 0x2c, 0x89,
	0x96, 0x71, 0xe7, 0xe2, 0xa1, 0xa8, 0xde, 0x98, 0x09, 0xba, 0x8f, 0x14, 0x65, 0x69, 0xbb, 0x72,
	0x63, 0xfb, 0xe7, 0x77, 0x17, 0x7c, 0x9c, 0xb8, 0xe0, 0xf3, 0xc4, 0x05, 0xdf, 0x26, 0x2e, 0x78,
	0x70, 0x4d, 0x11, 0x94, 0x44, 0x02, 0x25, 0x1e, 0x65, 0xfe, 0x3c, 0xde, 0x90, 0x44, 0x3c, 0x26,
	0xc2, 0x47, 0x9c, 0xfb, 0x3a, 0xa4, 0x98, 0xf8, 0x23, 0xba, 0x1f, 0xce, 0x63, 0x6d, 0x6b, 0x3e,
	0xa3, 0x25, 0xe3, 0xb3, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x45, 0x42, 0xe6, 0x33, 0x04,
	0x00, 0x00,
}

func (this *TLAuthExportLoginToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&auth.TLAuthExportLoginToken{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "ApiId: "+fmt.Sprintf("%#v", this.ApiId)+",\n")
	s = append(s, "ApiHash: "+fmt.Sprintf("%#v", this.ApiHash)+",\n")
	s = append(s, "ExceptIds: "+fmt.Sprintf("%#v", this.ExceptIds)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLAuthImportLoginToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.TLAuthImportLoginToken{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TLAuthAcceptLoginToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.TLAuthAcceptLoginToken{")
	s = append(s, "Constructor: "+fmt.Sprintf("%#v", this.Constructor)+",\n")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAuthTl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCAuthClient is the client API for RPCAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCAuthClient interface {
	// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
	AuthExportLoginToken(ctx context.Context, in *TLAuthExportLoginToken, opts ...grpc.CallOption) (*mtproto.Auth_LoginToken, error)
	// auth.importLoginToken token:bytes = auth.LoginToken;
	AuthImportLoginToken(ctx context.Context, in *TLAuthImportLoginToken, opts ...grpc.CallOption) (*mtproto.Auth_LoginToken, error)
	// auth.acceptLoginToken token:bytes = Authorization;
	AuthAcceptLoginToken(ctx context.Context, in *TLAuthAcceptLoginToken, opts ...grpc.CallOption) (*mtproto.Authorization, error)
}

type rPCAuthClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthClient(cc *grpc.ClientConn) RPCAuthClient {
	return &rPCAuthClient{cc}
}

func (c *rPCAuthClient) AuthExportLoginToken(ctx context.Context, in *TLAuthExportLoginToken, opts ...grpc.CallOption) (*mtproto.Auth_LoginToken, error) {
	out := new(mtproto.Auth_LoginToken)
	err := c.cc.Invoke(ctx, "/auth.RPCAuth/auth_exportLoginToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthClient) AuthImportLoginToken(ctx context.Context, in *TLAuthImportLoginToken, opts ...grpc.CallOption) (*mtproto.Auth_LoginToken, error) {
	out := new(mtproto.Auth_LoginToken)
	err := c.cc.Invoke(ctx, "/auth.RPCAuth/auth_importLoginToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthClient) AuthAcceptLoginToken(ctx context.Context, in *TLAuthAcceptLoginToken, opts ...grpc.CallOption) (*mtproto.Authorization, error) {
	out := new(mtproto.Authorization)
	err := c.cc.Invoke(ctx, "/auth.RPCAuth/auth_acceptLoginToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCAuthServer is the server API for RPCAuth service.
type RPCAuthServer interface {
	// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
	AuthExportLoginToken(context.Context, *TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error)
	// auth.importLoginToken token:bytes = auth.LoginToken;
	AuthImportLoginToken(context.Context, *TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error)
	// auth.acceptLoginToken token:bytes = Authorization;
	AuthAcceptLoginToken(context.Context, *TLAuthAcceptLoginToken) (*mtproto.Authorization, error)
}

// UnimplementedRPCAuthServer can be embedded to have forward compatible implementations.
type UnimplementedRPCAuthServer struct {
}

func (*UnimplementedRPCAuthServer) AuthExportLoginToken(ctx context.Context, req *TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthExportLoginToken not implemented")
}
func (*UnimplementedRPCAuthServer) AuthImportLoginToken(ctx context.Context, req *TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthImportLoginToken not implemented")
}
func (*UnimplementedRPCAuthServer) AuthAcceptLoginToken(ctx context.Context, req *TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthAcceptLoginToken not implemented")
}

func RegisterRPCAuthServer(s *grpc.Server, srv RPCAuthServer) {
	s.RegisterService(&_RPCAuth_serviceDesc, srv)
}

func _RPCAuth_AuthExportLoginToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthExportLoginToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthServer).AuthExportLoginToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RPCAuth/AuthExportLoginToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthServer).AuthExportLoginToken(ctx, req.(*TLAuthExportLoginToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuth_AuthImportLoginToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthImportLoginToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthServer).AuthImportLoginToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RPCAuth/AuthImportLoginToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthServer).AuthImportLoginToken(ctx, req.(*TLAuthImportLoginToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuth_AuthAcceptLoginToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLAuthAcceptLoginToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthServer).AuthAcceptLoginToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RPCAuth/AuthAcceptLoginToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthServer).AuthAcceptLoginToken(ctx, req.(*TLAuthAcceptLoginToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.RPCAuth",
	HandlerType: (*RPCAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "auth_exportLoginToken",
			Handler:    _RPCAuth_AuthExportLoginToken_Handler,
		},
		{
			MethodName: "auth_importLoginToken",
			Handler:    _RPCAuth_AuthImportLoginToken_Handler,
		},
		{
			MethodName: "auth_acceptLoginToken",
			Handler:    _RPCAuth_AuthAcceptLoginToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.tl.proto",
}

func (m *TLAuthExportLoginToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLAuthExportLoginToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLAuthExportLoginToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ExceptIds) > 0 {
		dAtA2 := make([]byte, len(m.ExceptIds)*10)
		var j1 int
		for _, num1 := range m.ExceptIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAuthTl(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ApiHash) > 0 {
		i -= len(m.ApiHash)
		copy(dAtA[i:], m.ApiHash)
		i = encodeVarintAuthTl(dAtA, i, uint64(len(m.ApiHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.ApiId != 0 {
		i = encodeVarintAuthTl(dAtA, i, uint64(m.ApiId))
		i--
		dAtA[i] = 0x18
	}
	if m.Constructor != 0 {
		i = encodeVarintAuthTl(dAtA, i, uint64(m.Constructor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLAuthImportLoginToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLAuthImportLoginToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLAuthImportLoginToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintAuthTl(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Constructor != 0 {
		i = encodeVarintAuthTl(dAtA, i, uint64(m.Constructor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLAuthAcceptLoginToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLAuthAcceptLoginToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLAuthAcceptLoginToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintAuthTl(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Constructor != 0 {
		i = encodeVarintAuthTl(dAtA, i, uint64(m.Constructor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthTl(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthTl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TLAuthExportLoginToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovAuthTl(uint64(m.Constructor))
	}
	if m.ApiId != 0 {
		n += 1 + sovAuthTl(uint64(m.ApiId))
	}
	l = len(m.ApiHash)
	if l > 0 {
		n += 1 + l + sovAuthTl(uint64(l))
	}
	if len(m.ExceptIds) > 0 {
		l = 0
		for _, e := range m.ExceptIds {
			l += sovAuthTl(uint64(e))
		}
		n += 1 + sovAuthTl(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLAuthImportLoginToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovAuthTl(uint64(m.Constructor))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuthTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLAuthAcceptLoginToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constructor != 0 {
		n += 1 + sovAuthTl(uint64(m.Constructor))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuthTl(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAuthTl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthTl(x uint64) (n int) {
	return sovAuthTl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLAuthExportLoginToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_auth_exportLoginToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_auth_exportLoginToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= TLConstructor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiId", wireType)
			}
			m.ApiId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthTl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthTl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuthTl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExceptIds = append(m.ExceptIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuthTl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAuthTl
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAuthTl
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ExceptIds) == 0 {
					m.ExceptIds = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAuthTl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExceptIds = append(m.ExceptIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExceptIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuthTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLAuthImportLoginToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_auth_importLoginToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_auth_importLoginToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= TLConstructor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthTl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthTl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token[:0], dAtA[iNdEx:postIndex]...)
			if m.Token == nil {
				m.Token = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLAuthAcceptLoginToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthTl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_auth_acceptLoginToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_auth_acceptLoginToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constructor", wireType)
			}
			m.Constructor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Constructor |= TLConstructor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthTl
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthTl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = append(m.Token[:0], dAtA[iNdEx:postIndex]...)
			if m.Token == nil {
				m.Token = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthTl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthTl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuthTl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthTl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthTl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuthTl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthTl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthTl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthTl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthTl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthTl = fmt.Errorf("proto: unexpected end of group")
)