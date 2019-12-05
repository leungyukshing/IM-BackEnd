// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im-entities.proto

package im_entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterRequest struct {
	Username             *string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password             *string  `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email                *string  `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type RegisternResponse struct {
	Code                 *string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message              *string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisternResponse) Reset()         { *m = RegisternResponse{} }
func (m *RegisternResponse) String() string { return proto.CompactTextString(m) }
func (*RegisternResponse) ProtoMessage()    {}
func (*RegisternResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{1}
}

func (m *RegisternResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisternResponse.Unmarshal(m, b)
}
func (m *RegisternResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisternResponse.Marshal(b, m, deterministic)
}
func (m *RegisternResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisternResponse.Merge(m, src)
}
func (m *RegisternResponse) XXX_Size() int {
	return xxx_messageInfo_RegisternResponse.Size(m)
}
func (m *RegisternResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisternResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisternResponse proto.InternalMessageInfo

func (m *RegisternResponse) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *RegisternResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type User struct {
	Userid               *string  `protobuf:"bytes,1,opt,name=userid" json:"userid,omitempty"`
	Username             *string  `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email                *string  `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type LoginRequest struct {
	Email                *string  `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password             *string  `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type LoginResponse struct {
	Code                 *string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message              *string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *LoginResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Chat struct {
	Chatid               *string  `protobuf:"bytes,1,opt,name=chatid" json:"chatid,omitempty"`
	Userid               *string  `protobuf:"bytes,2,opt,name=userid" json:"userid,omitempty"`
	Chatname             *string  `protobuf:"bytes,3,opt,name=chatname" json:"chatname,omitempty"`
	CreateTime           *string  `protobuf:"bytes,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	LastUpdateTime       *string  `protobuf:"bytes,5,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{5}
}

func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (m *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(m, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetChatid() string {
	if m != nil && m.Chatid != nil {
		return *m.Chatid
	}
	return ""
}

func (m *Chat) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *Chat) GetChatname() string {
	if m != nil && m.Chatname != nil {
		return *m.Chatname
	}
	return ""
}

func (m *Chat) GetCreateTime() string {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return ""
}

func (m *Chat) GetLastUpdateTime() string {
	if m != nil && m.LastUpdateTime != nil {
		return *m.LastUpdateTime
	}
	return ""
}

type GetChatListRequest struct {
	Userid               *string  `protobuf:"bytes,1,opt,name=userid" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChatListRequest) Reset()         { *m = GetChatListRequest{} }
func (m *GetChatListRequest) String() string { return proto.CompactTextString(m) }
func (*GetChatListRequest) ProtoMessage()    {}
func (*GetChatListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{6}
}

func (m *GetChatListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChatListRequest.Unmarshal(m, b)
}
func (m *GetChatListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChatListRequest.Marshal(b, m, deterministic)
}
func (m *GetChatListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChatListRequest.Merge(m, src)
}
func (m *GetChatListRequest) XXX_Size() int {
	return xxx_messageInfo_GetChatListRequest.Size(m)
}
func (m *GetChatListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChatListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChatListRequest proto.InternalMessageInfo

func (m *GetChatListRequest) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

type GetChatListResponse struct {
	Code                 *string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Chat                 []*Chat  `protobuf:"bytes,2,rep,name=chat" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChatListResponse) Reset()         { *m = GetChatListResponse{} }
func (m *GetChatListResponse) String() string { return proto.CompactTextString(m) }
func (*GetChatListResponse) ProtoMessage()    {}
func (*GetChatListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{7}
}

func (m *GetChatListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChatListResponse.Unmarshal(m, b)
}
func (m *GetChatListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChatListResponse.Marshal(b, m, deterministic)
}
func (m *GetChatListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChatListResponse.Merge(m, src)
}
func (m *GetChatListResponse) XXX_Size() int {
	return xxx_messageInfo_GetChatListResponse.Size(m)
}
func (m *GetChatListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChatListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetChatListResponse proto.InternalMessageInfo

func (m *GetChatListResponse) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *GetChatListResponse) GetChat() []*Chat {
	if m != nil {
		return m.Chat
	}
	return nil
}

type GetContactListRequest struct {
	Userid               *string  `protobuf:"bytes,1,opt,name=userid" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContactListRequest) Reset()         { *m = GetContactListRequest{} }
func (m *GetContactListRequest) String() string { return proto.CompactTextString(m) }
func (*GetContactListRequest) ProtoMessage()    {}
func (*GetContactListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{8}
}

func (m *GetContactListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContactListRequest.Unmarshal(m, b)
}
func (m *GetContactListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContactListRequest.Marshal(b, m, deterministic)
}
func (m *GetContactListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContactListRequest.Merge(m, src)
}
func (m *GetContactListRequest) XXX_Size() int {
	return xxx_messageInfo_GetContactListRequest.Size(m)
}
func (m *GetContactListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContactListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContactListRequest proto.InternalMessageInfo

func (m *GetContactListRequest) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

type GetContactListResponse struct {
	Code                 *string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	User                 []*User  `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContactListResponse) Reset()         { *m = GetContactListResponse{} }
func (m *GetContactListResponse) String() string { return proto.CompactTextString(m) }
func (*GetContactListResponse) ProtoMessage()    {}
func (*GetContactListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3017226e81765a8, []int{9}
}

func (m *GetContactListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContactListResponse.Unmarshal(m, b)
}
func (m *GetContactListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContactListResponse.Marshal(b, m, deterministic)
}
func (m *GetContactListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContactListResponse.Merge(m, src)
}
func (m *GetContactListResponse) XXX_Size() int {
	return xxx_messageInfo_GetContactListResponse.Size(m)
}
func (m *GetContactListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContactListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetContactListResponse proto.InternalMessageInfo

func (m *GetContactListResponse) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *GetContactListResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisternResponse)(nil), "RegisternResponse")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterType((*GetChatListRequest)(nil), "GetChatListRequest")
	proto.RegisterType((*GetChatListResponse)(nil), "GetChatListResponse")
	proto.RegisterType((*GetContactListRequest)(nil), "GetContactListRequest")
	proto.RegisterType((*GetContactListResponse)(nil), "GetContactListResponse")
}

func init() { proto.RegisterFile("im-entities.proto", fileDescriptor_b3017226e81765a8) }

var fileDescriptor_b3017226e81765a8 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0x25, 0x69, 0xfa, 0x7d, 0x3a, 0xf5, 0xaf, 0xab, 0x96, 0xd4, 0x1b, 0xcb, 0x5e, 0xf5, 0x42,
	0x2b, 0xf8, 0x04, 0x8a, 0x42, 0x6f, 0x7a, 0x21, 0xc1, 0x82, 0x77, 0x65, 0x49, 0x86, 0x76, 0xa1,
	0xf9, 0x31, 0x33, 0xc5, 0xa7, 0xf1, 0x5d, 0x65, 0x37, 0xd9, 0xf4, 0x07, 0x2d, 0xe2, 0xdd, 0x9e,
	0x99, 0xb3, 0x33, 0xe7, 0x9c, 0x81, 0xae, 0x4e, 0x6f, 0x31, 0x63, 0xcd, 0x1a, 0x69, 0x54, 0x94,
	0x39, 0xe7, 0x72, 0x06, 0xa7, 0x11, 0xce, 0x35, 0x31, 0x96, 0x11, 0xbe, 0xaf, 0x90, 0x58, 0x5c,
	0xc1, 0xc1, 0x8a, 0xb0, 0xcc, 0x54, 0x8a, 0xa1, 0x37, 0xf0, 0x86, 0x87, 0x51, 0x83, 0x4d, 0xaf,
	0x50, 0x44, 0x1f, 0x79, 0x99, 0x84, 0x7e, 0xd5, 0x73, 0x58, 0x5c, 0x40, 0x1b, 0x53, 0xa5, 0x97,
	0x61, 0xcb, 0x36, 0x2a, 0x20, 0x1f, 0xa1, 0xeb, 0x16, 0x64, 0x11, 0x52, 0x91, 0x67, 0x84, 0x42,
	0x40, 0x10, 0xe7, 0x89, 0x1b, 0x6f, 0xdf, 0x22, 0x84, 0xff, 0x29, 0x12, 0xa9, 0x39, 0xd6, 0x93,
	0x1d, 0x94, 0x2f, 0x10, 0x4c, 0x09, 0x4b, 0xd1, 0x83, 0x7f, 0x46, 0x88, 0x4e, 0xea, 0x7f, 0x35,
	0xda, 0x12, 0xec, 0xef, 0x08, 0xfe, 0x5e, 0xd4, 0x03, 0x1c, 0x4d, 0xf2, 0xb9, 0xce, 0x9c, 0xe5,
	0x86, 0xe5, 0x6d, 0xb0, 0xf6, 0x99, 0x95, 0x6f, 0x70, 0x5c, 0x4f, 0xf8, 0x8b, 0x25, 0xd1, 0x87,
	0xc0, 0x48, 0xb4, 0xaa, 0x3a, 0xf7, 0xed, 0x91, 0xf1, 0x17, 0xd9, 0x92, 0xfc, 0xf4, 0x20, 0x78,
	0x5a, 0x28, 0x36, 0x76, 0xe3, 0x85, 0xe2, 0xb5, 0xdd, 0x0a, 0x6d, 0xc4, 0xe0, 0xef, 0xc6, 0x60,
	0x18, 0x36, 0x86, 0xca, 0x6d, 0x83, 0xc5, 0x35, 0x74, 0xe2, 0x12, 0x15, 0xe3, 0x8c, 0x75, 0x8a,
	0x61, 0x60, 0xdb, 0x50, 0x95, 0x5e, 0x75, 0x8a, 0x62, 0x08, 0x67, 0x4b, 0x45, 0x3c, 0x5b, 0x15,
	0x49, 0xc3, 0x6a, 0x5b, 0xd6, 0x89, 0xa9, 0x4f, 0x6d, 0xd9, 0x30, 0xe5, 0x0d, 0x88, 0x31, 0xb2,
	0x51, 0x38, 0xd1, 0xc4, 0x2e, 0xc1, 0x1f, 0x6e, 0x23, 0x9f, 0xe1, 0x7c, 0x8b, 0xbd, 0x27, 0xad,
	0x3e, 0x04, 0x46, 0x6f, 0xe8, 0x0f, 0x5a, 0x36, 0x13, 0xf3, 0x29, 0xb2, 0x25, 0x79, 0x07, 0x97,
	0x66, 0x4a, 0x9e, 0xb1, 0x8a, 0x7f, 0xb5, 0x76, 0x0c, 0xbd, 0xdd, 0x0f, 0xfb, 0x37, 0xdb, 0x6b,
	0xb8, 0xcd, 0xeb, 0x6b, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x63, 0x86, 0x04, 0x0b, 0x33, 0x03,
	0x00, 0x00,
}