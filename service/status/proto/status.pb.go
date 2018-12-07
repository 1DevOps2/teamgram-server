// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package status

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// //////////////////////////////////////////////////////////////////////////////////////
type Int32 struct {
	V                    int32    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{0}
}
func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (dst *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(dst, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetV() int32 {
	if m != nil {
		return m.V
	}
	return 0
}

type Int32List struct {
	Vlist                []int32  `protobuf:"varint,1,rep,packed,name=vlist,proto3" json:"vlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32List) Reset()         { *m = Int32List{} }
func (m *Int32List) String() string { return proto.CompactTextString(m) }
func (*Int32List) ProtoMessage()    {}
func (*Int32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{1}
}
func (m *Int32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32List.Unmarshal(m, b)
}
func (m *Int32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32List.Marshal(b, m, deterministic)
}
func (dst *Int32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32List.Merge(dst, src)
}
func (m *Int32List) XXX_Size() int {
	return xxx_messageInfo_Int32List.Size(m)
}
func (m *Int32List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32List.DiscardUnknown(m)
}

var xxx_messageInfo_Int32List proto.InternalMessageInfo

func (m *Int32List) GetVlist() []int32 {
	if m != nil {
		return m.Vlist
	}
	return nil
}

type Int64List struct {
	Vlist                []int64  `protobuf:"varint,1,rep,packed,name=vlist,proto3" json:"vlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64List) Reset()         { *m = Int64List{} }
func (m *Int64List) String() string { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()    {}
func (*Int64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{2}
}
func (m *Int64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64List.Unmarshal(m, b)
}
func (m *Int64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64List.Marshal(b, m, deterministic)
}
func (dst *Int64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64List.Merge(dst, src)
}
func (m *Int64List) XXX_Size() int {
	return xxx_messageInfo_Int64List.Size(m)
}
func (m *Int64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64List.DiscardUnknown(m)
}

var xxx_messageInfo_Int64List proto.InternalMessageInfo

func (m *Int64List) GetVlist() []int64 {
	if m != nil {
		return m.Vlist
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{3}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

// //////////////////////////////////////////////////////////////////////////////////////
type SessionEntry struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ServerId             int32    `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	AuthKeyId            int64    `protobuf:"varint,3,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	Expired              int64    `protobuf:"varint,4,opt,name=expired,proto3" json:"expired,omitempty"`
	Layer                int32    `protobuf:"varint,5,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionEntry) Reset()         { *m = SessionEntry{} }
func (m *SessionEntry) String() string { return proto.CompactTextString(m) }
func (*SessionEntry) ProtoMessage()    {}
func (*SessionEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{4}
}
func (m *SessionEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionEntry.Unmarshal(m, b)
}
func (m *SessionEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionEntry.Marshal(b, m, deterministic)
}
func (dst *SessionEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionEntry.Merge(dst, src)
}
func (m *SessionEntry) XXX_Size() int {
	return xxx_messageInfo_SessionEntry.Size(m)
}
func (m *SessionEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SessionEntry proto.InternalMessageInfo

func (m *SessionEntry) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SessionEntry) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *SessionEntry) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *SessionEntry) GetExpired() int64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *SessionEntry) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

type SessionEntryList struct {
	Sessions             []*SessionEntry `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	PushSessions         []*SessionEntry `protobuf:"bytes,2,rep,name=push_sessions,json=pushSessions,proto3" json:"push_sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SessionEntryList) Reset()         { *m = SessionEntryList{} }
func (m *SessionEntryList) String() string { return proto.CompactTextString(m) }
func (*SessionEntryList) ProtoMessage()    {}
func (*SessionEntryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{5}
}
func (m *SessionEntryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionEntryList.Unmarshal(m, b)
}
func (m *SessionEntryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionEntryList.Marshal(b, m, deterministic)
}
func (dst *SessionEntryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionEntryList.Merge(dst, src)
}
func (m *SessionEntryList) XXX_Size() int {
	return xxx_messageInfo_SessionEntryList.Size(m)
}
func (m *SessionEntryList) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionEntryList.DiscardUnknown(m)
}

var xxx_messageInfo_SessionEntryList proto.InternalMessageInfo

func (m *SessionEntryList) GetSessions() []*SessionEntry {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func (m *SessionEntryList) GetPushSessions() []*SessionEntry {
	if m != nil {
		return m.PushSessions
	}
	return nil
}

type UsersSessionEntryList struct {
	UsersSessions        map[int32]*SessionEntryList `protobuf:"bytes,1,rep,name=users_sessions,json=usersSessions,proto3" json:"users_sessions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UsersSessionEntryList) Reset()         { *m = UsersSessionEntryList{} }
func (m *UsersSessionEntryList) String() string { return proto.CompactTextString(m) }
func (*UsersSessionEntryList) ProtoMessage()    {}
func (*UsersSessionEntryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_3b4897f4a6c6f9d6, []int{6}
}
func (m *UsersSessionEntryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersSessionEntryList.Unmarshal(m, b)
}
func (m *UsersSessionEntryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersSessionEntryList.Marshal(b, m, deterministic)
}
func (dst *UsersSessionEntryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersSessionEntryList.Merge(dst, src)
}
func (m *UsersSessionEntryList) XXX_Size() int {
	return xxx_messageInfo_UsersSessionEntryList.Size(m)
}
func (m *UsersSessionEntryList) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersSessionEntryList.DiscardUnknown(m)
}

var xxx_messageInfo_UsersSessionEntryList proto.InternalMessageInfo

func (m *UsersSessionEntryList) GetUsersSessions() map[int32]*SessionEntryList {
	if m != nil {
		return m.UsersSessions
	}
	return nil
}

func init() {
	proto.RegisterType((*Int32)(nil), "status.Int32")
	proto.RegisterType((*Int32List)(nil), "status.Int32List")
	proto.RegisterType((*Int64List)(nil), "status.Int64List")
	proto.RegisterType((*Void)(nil), "status.Void")
	proto.RegisterType((*SessionEntry)(nil), "status.SessionEntry")
	proto.RegisterType((*SessionEntryList)(nil), "status.SessionEntryList")
	proto.RegisterType((*UsersSessionEntryList)(nil), "status.UsersSessionEntryList")
	proto.RegisterMapType((map[int32]*SessionEntryList)(nil), "status.UsersSessionEntryList.UsersSessionsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCStatusClient is the client API for RPCStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCStatusClient interface {
	SetSessionOnline(ctx context.Context, in *SessionEntry, opts ...grpc.CallOption) (*Void, error)
	SetSessionOffline(ctx context.Context, in *SessionEntry, opts ...grpc.CallOption) (*Void, error)
	GetUserOnlineSessions(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SessionEntryList, error)
	GetUsersOnlineSessionsList(ctx context.Context, in *Int32List, opts ...grpc.CallOption) (*UsersSessionEntryList, error)
}

type rPCStatusClient struct {
	cc *grpc.ClientConn
}

func NewRPCStatusClient(cc *grpc.ClientConn) RPCStatusClient {
	return &rPCStatusClient{cc}
}

func (c *rPCStatusClient) SetSessionOnline(ctx context.Context, in *SessionEntry, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/status.RPCStatus/SetSessionOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCStatusClient) SetSessionOffline(ctx context.Context, in *SessionEntry, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/status.RPCStatus/SetSessionOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCStatusClient) GetUserOnlineSessions(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SessionEntryList, error) {
	out := new(SessionEntryList)
	err := c.cc.Invoke(ctx, "/status.RPCStatus/GetUserOnlineSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCStatusClient) GetUsersOnlineSessionsList(ctx context.Context, in *Int32List, opts ...grpc.CallOption) (*UsersSessionEntryList, error) {
	out := new(UsersSessionEntryList)
	err := c.cc.Invoke(ctx, "/status.RPCStatus/GetUsersOnlineSessionsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCStatusServer is the server API for RPCStatus service.
type RPCStatusServer interface {
	SetSessionOnline(context.Context, *SessionEntry) (*Void, error)
	SetSessionOffline(context.Context, *SessionEntry) (*Void, error)
	GetUserOnlineSessions(context.Context, *Int32) (*SessionEntryList, error)
	GetUsersOnlineSessionsList(context.Context, *Int32List) (*UsersSessionEntryList, error)
}

func RegisterRPCStatusServer(s *grpc.Server, srv RPCStatusServer) {
	s.RegisterService(&_RPCStatus_serviceDesc, srv)
}

func _RPCStatus_SetSessionOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCStatusServer).SetSessionOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.RPCStatus/SetSessionOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCStatusServer).SetSessionOnline(ctx, req.(*SessionEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCStatus_SetSessionOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCStatusServer).SetSessionOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.RPCStatus/SetSessionOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCStatusServer).SetSessionOffline(ctx, req.(*SessionEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCStatus_GetUserOnlineSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCStatusServer).GetUserOnlineSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.RPCStatus/GetUserOnlineSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCStatusServer).GetUserOnlineSessions(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCStatus_GetUsersOnlineSessionsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCStatusServer).GetUsersOnlineSessionsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/status.RPCStatus/GetUsersOnlineSessionsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCStatusServer).GetUsersOnlineSessionsList(ctx, req.(*Int32List))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "status.RPCStatus",
	HandlerType: (*RPCStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSessionOnline",
			Handler:    _RPCStatus_SetSessionOnline_Handler,
		},
		{
			MethodName: "SetSessionOffline",
			Handler:    _RPCStatus_SetSessionOffline_Handler,
		},
		{
			MethodName: "GetUserOnlineSessions",
			Handler:    _RPCStatus_GetUserOnlineSessions_Handler,
		},
		{
			MethodName: "GetUsersOnlineSessionsList",
			Handler:    _RPCStatus_GetUsersOnlineSessionsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "status.proto",
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_status_3b4897f4a6c6f9d6) }

var fileDescriptor_status_3b4897f4a6c6f9d6 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x50,
	0x10, 0xe5, 0x26, 0x9b, 0xee, 0x76, 0x9a, 0xca, 0xee, 0xb0, 0xc5, 0x10, 0x51, 0x6a, 0x9e, 0x0a,
	0x42, 0x58, 0xba, 0xb2, 0x7e, 0x3c, 0xc9, 0x8a, 0x68, 0x51, 0xb4, 0x64, 0x51, 0xc1, 0x97, 0x92,
	0x35, 0x77, 0xd9, 0xcb, 0xc6, 0xb4, 0x64, 0x6e, 0x82, 0x79, 0xf2, 0x1f, 0xf8, 0xe4, 0xcf, 0xf2,
	0x47, 0xc9, 0xfd, 0xc8, 0x92, 0xda, 0x5a, 0x7c, 0xcb, 0x9c, 0x39, 0x67, 0x66, 0xce, 0xdc, 0x09,
	0xf8, 0x24, 0x53, 0x59, 0x51, 0xbc, 0x2a, 0x97, 0x72, 0x89, 0x3d, 0x13, 0x45, 0x23, 0xf0, 0x66,
	0x85, 0x3c, 0x9d, 0xa2, 0x0f, 0xac, 0x0e, 0xd8, 0x98, 0x4d, 0xbc, 0x84, 0xd5, 0xd1, 0x43, 0xe8,
	0x6b, 0xf8, 0x9d, 0x20, 0x89, 0xc7, 0xe0, 0xd5, 0xb9, 0x20, 0x19, 0xb0, 0xb1, 0x3b, 0xf1, 0x12,
	0x13, 0x58, 0xca, 0xd9, 0xe3, 0x4d, 0x8a, 0xdb, 0x52, 0x7a, 0xb0, 0xf7, 0x69, 0x29, 0xb2, 0xe8,
	0x17, 0x03, 0xff, 0x82, 0x13, 0x89, 0x65, 0xf1, 0xaa, 0x90, 0x65, 0x83, 0x77, 0x61, 0xbf, 0x22,
	0x5e, 0x2e, 0x44, 0x66, 0x5b, 0xf6, 0x54, 0x38, 0xcb, 0xf0, 0x1e, 0xf4, 0x89, 0x97, 0xb5, 0x49,
	0x39, 0x3a, 0x75, 0x60, 0x80, 0x59, 0x86, 0x0f, 0x60, 0x90, 0x56, 0xf2, 0x7a, 0x71, 0xc3, 0x1b,
	0x95, 0x76, 0xc7, 0x6c, 0xe2, 0x26, 0x7d, 0x05, 0xbd, 0xe5, 0xcd, 0x2c, 0xc3, 0x00, 0xf6, 0xf9,
	0xf7, 0x95, 0x28, 0x79, 0x16, 0xec, 0xe9, 0x5c, 0x1b, 0xaa, 0xf1, 0xf2, 0xb4, 0xe1, 0x65, 0xe0,
	0xe9, 0x92, 0x26, 0x88, 0x7e, 0xc0, 0x61, 0x77, 0x2a, 0x6d, 0xe4, 0x04, 0x0e, 0xc8, 0x60, 0xa4,
	0xbd, 0x0c, 0xa6, 0xc7, 0xb1, 0x5d, 0x5c, 0x97, 0x9b, 0xdc, 0xb2, 0xf0, 0x19, 0x0c, 0x57, 0x15,
	0x5d, 0x2f, 0x6e, 0x65, 0xce, 0x0e, 0x99, 0xaf, 0xa8, 0x16, 0xa1, 0xe8, 0x37, 0x83, 0xd1, 0x47,
	0xe2, 0x25, 0x6d, 0x8c, 0xf1, 0x19, 0xee, 0xa8, 0x8d, 0xd0, 0xe2, 0xaf, 0x61, 0x4e, 0xda, 0xaa,
	0x5b, 0x65, 0x6b, 0x28, 0x99, 0x8e, 0xc3, 0xaa, 0x8b, 0x85, 0x5f, 0x00, 0x37, 0x49, 0x78, 0x08,
	0xee, 0x0d, 0x6f, 0xec, 0x5b, 0xa8, 0x4f, 0x8c, 0xc1, 0xab, 0xd3, 0xbc, 0xe2, 0xfa, 0x11, 0x06,
	0xd3, 0x60, 0x9b, 0x1b, 0xd5, 0x32, 0x31, 0xb4, 0xe7, 0xce, 0x53, 0x36, 0xfd, 0xe9, 0x40, 0x3f,
	0x99, 0xbf, 0xbc, 0xd0, 0x4c, 0x3c, 0x53, 0xdb, 0x95, 0x96, 0xff, 0xa1, 0xc8, 0x45, 0xc1, 0x71,
	0xeb, 0x52, 0x42, 0xbf, 0x45, 0xd5, 0xb1, 0xe0, 0x13, 0x38, 0xea, 0xe8, 0xae, 0xae, 0xfe, 0x5b,
	0xf8, 0x02, 0x46, 0xaf, 0xb9, 0x54, 0xee, 0x4c, 0xb7, 0xd6, 0x22, 0x0e, 0x5b, 0x9a, 0x3e, 0xe9,
	0xf0, 0x9f, 0x5e, 0xf0, 0x3d, 0x84, 0xb6, 0x02, 0xad, 0x97, 0xd0, 0xd9, 0xa3, 0xb5, 0x32, 0x0a,
	0x0a, 0xef, 0xef, 0x7c, 0x8e, 0xf3, 0x47, 0x10, 0x88, 0x6f, 0x71, 0xc1, 0x2f, 0xab, 0x3c, 0x8d,
	0xd5, 0x19, 0x8b, 0xaf, 0xdc, 0x2a, 0xce, 0x07, 0x66, 0x4d, 0x73, 0xf5, 0x37, 0xbe, 0x71, 0xe6,
	0xec, 0xb2, 0xa7, 0x7f, 0xcc, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xcc, 0x49, 0xa5,
	0xa8, 0x03, 0x00, 0x00,
}