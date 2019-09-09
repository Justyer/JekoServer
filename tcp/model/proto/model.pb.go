// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

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

type RoomInfo struct {
	ID                   int32       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserList             []*UserInfo `protobuf:"bytes,2,rep,name=UserList,proto3" json:"UserList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *RoomInfo) GetUserList() []*UserInfo {
	if m != nil {
		return m.UserList
	}
	return nil
}

type UserInfo struct {
	UserID               int32    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	IconURL              string   `protobuf:"bytes,3,opt,name=IconURL,proto3" json:"IconURL,omitempty"`
	Weapon               *Weapon  `protobuf:"bytes,4,opt,name=Weapon,proto3" json:"Weapon,omitempty"`
	AttributeNum         int32    `protobuf:"varint,5,opt,name=AttributeNum,proto3" json:"AttributeNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetIconURL() string {
	if m != nil {
		return m.IconURL
	}
	return ""
}

func (m *UserInfo) GetWeapon() *Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func (m *UserInfo) GetAttributeNum() int32 {
	if m != nil {
		return m.AttributeNum
	}
	return 0
}

type PlayerInit struct {
	UserID               int32       `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName             string      `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	HeroSN               int32       `protobuf:"varint,3,opt,name=heroSN,proto3" json:"heroSN,omitempty"`
	Position             *Vector3DTO `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Direct               *Vector3DTO `protobuf:"bytes,5,opt,name=direct,proto3" json:"direct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlayerInit) Reset()         { *m = PlayerInit{} }
func (m *PlayerInit) String() string { return proto.CompactTextString(m) }
func (*PlayerInit) ProtoMessage()    {}
func (*PlayerInit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *PlayerInit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerInit.Unmarshal(m, b)
}
func (m *PlayerInit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerInit.Marshal(b, m, deterministic)
}
func (m *PlayerInit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInit.Merge(m, src)
}
func (m *PlayerInit) XXX_Size() int {
	return xxx_messageInfo_PlayerInit.Size(m)
}
func (m *PlayerInit) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInit.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInit proto.InternalMessageInfo

func (m *PlayerInit) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *PlayerInit) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PlayerInit) GetHeroSN() int32 {
	if m != nil {
		return m.HeroSN
	}
	return 0
}

func (m *PlayerInit) GetPosition() *Vector3DTO {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PlayerInit) GetDirect() *Vector3DTO {
	if m != nil {
		return m.Direct
	}
	return nil
}

type Weapon struct {
	ID                   int32              `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SN                   int32              `protobuf:"varint,2,opt,name=SN,proto3" json:"SN,omitempty"`
	WeaponExtraAttrList  []*WeaponExtraAttr `protobuf:"bytes,3,rep,name=WeaponExtraAttrList,proto3" json:"WeaponExtraAttrList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Weapon) Reset()         { *m = Weapon{} }
func (m *Weapon) String() string { return proto.CompactTextString(m) }
func (*Weapon) ProtoMessage()    {}
func (*Weapon) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *Weapon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weapon.Unmarshal(m, b)
}
func (m *Weapon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weapon.Marshal(b, m, deterministic)
}
func (m *Weapon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weapon.Merge(m, src)
}
func (m *Weapon) XXX_Size() int {
	return xxx_messageInfo_Weapon.Size(m)
}
func (m *Weapon) XXX_DiscardUnknown() {
	xxx_messageInfo_Weapon.DiscardUnknown(m)
}

var xxx_messageInfo_Weapon proto.InternalMessageInfo

func (m *Weapon) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Weapon) GetSN() int32 {
	if m != nil {
		return m.SN
	}
	return 0
}

func (m *Weapon) GetWeaponExtraAttrList() []*WeaponExtraAttr {
	if m != nil {
		return m.WeaponExtraAttrList
	}
	return nil
}

type WeaponExtraAttr struct {
	AttrType             int32    `protobuf:"varint,1,opt,name=AttrType,proto3" json:"AttrType,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeaponExtraAttr) Reset()         { *m = WeaponExtraAttr{} }
func (m *WeaponExtraAttr) String() string { return proto.CompactTextString(m) }
func (*WeaponExtraAttr) ProtoMessage()    {}
func (*WeaponExtraAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

func (m *WeaponExtraAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeaponExtraAttr.Unmarshal(m, b)
}
func (m *WeaponExtraAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeaponExtraAttr.Marshal(b, m, deterministic)
}
func (m *WeaponExtraAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeaponExtraAttr.Merge(m, src)
}
func (m *WeaponExtraAttr) XXX_Size() int {
	return xxx_messageInfo_WeaponExtraAttr.Size(m)
}
func (m *WeaponExtraAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_WeaponExtraAttr.DiscardUnknown(m)
}

var xxx_messageInfo_WeaponExtraAttr proto.InternalMessageInfo

func (m *WeaponExtraAttr) GetAttrType() int32 {
	if m != nil {
		return m.AttrType
	}
	return 0
}

func (m *WeaponExtraAttr) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Vector3DTO struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3DTO) Reset()         { *m = Vector3DTO{} }
func (m *Vector3DTO) String() string { return proto.CompactTextString(m) }
func (*Vector3DTO) ProtoMessage()    {}
func (*Vector3DTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

func (m *Vector3DTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3DTO.Unmarshal(m, b)
}
func (m *Vector3DTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3DTO.Marshal(b, m, deterministic)
}
func (m *Vector3DTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3DTO.Merge(m, src)
}
func (m *Vector3DTO) XXX_Size() int {
	return xxx_messageInfo_Vector3DTO.Size(m)
}
func (m *Vector3DTO) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3DTO.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3DTO proto.InternalMessageInfo

func (m *Vector3DTO) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3DTO) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3DTO) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*RoomInfo)(nil), "prt.RoomInfo")
	proto.RegisterType((*UserInfo)(nil), "prt.UserInfo")
	proto.RegisterType((*PlayerInit)(nil), "prt.PlayerInit")
	proto.RegisterType((*Weapon)(nil), "prt.Weapon")
	proto.RegisterType((*WeaponExtraAttr)(nil), "prt.WeaponExtraAttr")
	proto.RegisterType((*Vector3DTO)(nil), "prt.Vector3DTO")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x99, 0xc9, 0x4d, 0xae, 0xf7, 0xc4, 0xab, 0x30, 0x15, 0x09, 0x5d, 0x85, 0x74, 0xd1,
	0x94, 0x82, 0x0b, 0x85, 0xee, 0x4b, 0xb5, 0x10, 0x90, 0xb4, 0x8c, 0x7f, 0xba, 0x8e, 0x3a, 0xa5,
	0x01, 0xe3, 0x84, 0x71, 0x04, 0xe3, 0xdb, 0xf4, 0x05, 0xfa, 0x8c, 0x65, 0x4e, 0x62, 0xfa, 0xcf,
	0x55, 0x57, 0xc9, 0x77, 0xbe, 0x6f, 0x26, 0xbf, 0x73, 0x4e, 0xc0, 0xcd, 0xe4, 0x4a, 0xac, 0x7b,
	0xb9, 0x92, 0x5a, 0x32, 0x2b, 0x57, 0x3a, 0x18, 0x41, 0x83, 0x4b, 0x99, 0x45, 0x9b, 0x67, 0xc9,
	0x5a, 0x40, 0xa3, 0xa1, 0x47, 0x7c, 0x12, 0xda, 0x9c, 0x46, 0x43, 0x76, 0x05, 0x8d, 0xd9, 0x56,
	0xa8, 0x71, 0xba, 0xd5, 0x1e, 0xf5, 0xad, 0xd0, 0xed, 0xff, 0xef, 0xe5, 0x4a, 0xf7, 0x4c, 0xd1,
	0x1c, 0xe0, 0xb5, 0x1d, 0xbc, 0x92, 0x32, 0x8b, 0xf7, 0x74, 0xc1, 0xc1, 0xf7, 0xe3, 0x5d, 0x95,
	0x62, 0xe7, 0x65, 0x26, 0x4e, 0x32, 0xe1, 0x51, 0x9f, 0x84, 0xff, 0x78, 0xad, 0x99, 0x07, 0x7f,
	0xa3, 0xa5, 0xdc, 0xcc, 0xf8, 0xd8, 0xb3, 0xd0, 0x3a, 0x4a, 0x76, 0x01, 0xce, 0x93, 0x48, 0x72,
	0xb9, 0xf1, 0xfe, 0xf8, 0x24, 0x74, 0xfb, 0x2e, 0x32, 0x94, 0x25, 0x5e, 0x59, 0x2c, 0x80, 0xe6,
	0xad, 0xd6, 0x2a, 0x5d, 0xec, 0xb4, 0x88, 0x77, 0x99, 0x67, 0xe3, 0x87, 0xbf, 0xd4, 0x82, 0x37,
	0x02, 0xf0, 0xb8, 0x4e, 0x0a, 0x43, 0x99, 0xea, 0x5f, 0x51, 0x76, 0xc1, 0x79, 0x11, 0x4a, 0x4e,
	0x62, 0x84, 0xb4, 0x79, 0xa5, 0xd8, 0x35, 0x34, 0x72, 0xb9, 0x4d, 0x75, 0x5a, 0x53, 0xb6, 0x91,
	0x72, 0x2e, 0x96, 0x5a, 0xaa, 0xc1, 0x70, 0xfa, 0xc0, 0xeb, 0x00, 0xbb, 0x04, 0x67, 0x95, 0x2a,
	0xb1, 0xd4, 0x48, 0x79, 0x22, 0x5a, 0xd9, 0x41, 0x7e, 0xec, 0xfc, 0xc7, 0x66, 0x5a, 0x40, 0x27,
	0x31, 0xd2, 0xd9, 0x9c, 0x4e, 0x62, 0x76, 0x0f, 0x67, 0x65, 0x72, 0xb4, 0xd7, 0x2a, 0x31, 0x5d,
	0xe3, 0xd2, 0x2c, 0x5c, 0x5a, 0xe7, 0xd3, 0xc0, 0x6a, 0x9f, 0x9f, 0x3a, 0x10, 0xdc, 0x41, 0xfb,
	0x5b, 0xd9, 0x8c, 0xc3, 0x3c, 0xa7, 0x45, 0x2e, 0x2a, 0x80, 0x5a, 0xb3, 0x0e, 0xd8, 0xf3, 0x64,
	0xbd, 0x13, 0x15, 0x49, 0x29, 0x82, 0x1b, 0x80, 0x8f, 0x66, 0x58, 0x13, 0xc8, 0x1e, 0x0f, 0x52,
	0x4e, 0xf6, 0x46, 0x15, 0x98, 0xa6, 0x9c, 0x14, 0x46, 0x1d, 0x70, 0x92, 0x94, 0x93, 0xc3, 0xc2,
	0xc1, 0xdf, 0x72, 0xf0, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x11, 0x7c, 0x1d, 0x6c, 0xa5, 0x02, 0x00,
	0x00,
}