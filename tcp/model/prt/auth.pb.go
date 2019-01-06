// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package prt

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

type LoginReq struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	MAC                  string   `protobuf:"bytes,2,opt,name=MAC,proto3" json:"MAC,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	PassWord             string   `protobuf:"bytes,4,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *LoginReq) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *LoginReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginReq) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

type LoginResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	UserID               int32    `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResp) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *LoginResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "prt.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "prt.LoginResp")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x28, 0x2a, 0x51, 0xca, 0xe0, 0xe2, 0xf0,
	0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0x09, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x7d, 0x1d, 0x9d, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x29, 0x2e, 0x8e, 0xd0, 0xe2, 0xd4, 0x22,
	0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0xb0, 0x30, 0x9c, 0x0f, 0x92, 0x0b, 0x48, 0x2c, 0x2e, 0x0e,
	0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x81, 0xc8, 0xc1, 0xf8, 0x4a, 0xc1, 0x5c, 0x9c, 0x50, 0x9b, 0x8a,
	0x0b, 0x40, 0x56, 0x39, 0xe7, 0xa7, 0xc0, 0xad, 0x02, 0xb1, 0x85, 0xc4, 0xb8, 0xd8, 0x40, 0x06,
	0x79, 0xba, 0x80, 0x6d, 0x63, 0x0d, 0x82, 0xf2, 0xf0, 0x59, 0x98, 0xc4, 0x06, 0xf6, 0x8a, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xaf, 0x18, 0xec, 0xd8, 0x00, 0x00, 0x00,
}
