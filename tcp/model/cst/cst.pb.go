// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cst.proto

package cst

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

type MsgType int32

const (
	MsgType_Login  MsgType = 0
	MsgType_Room   MsgType = 1
	MsgType_Ready  MsgType = 2
	MsgType_Combat MsgType = 3
)

var MsgType_name = map[int32]string{
	0: "Login",
	1: "Room",
	2: "Ready",
	3: "Combat",
}

var MsgType_value = map[string]int32{
	"Login":  0,
	"Room":   1,
	"Ready":  2,
	"Combat": 3,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8ab54014d815c68, []int{0}
}

type MsgCmd int32

const (
	// 登录
	MsgCmd_Login_LoginReq  MsgCmd = 0
	MsgCmd_Login_LoginResp MsgCmd = 1
	// 查询房间列表
	MsgCmd_Room_QueryListReq  MsgCmd = 2
	MsgCmd_Room_QueryListResp MsgCmd = 3
	// 进入房间
	MsgCmd_Room_GetInReq  MsgCmd = 4
	MsgCmd_Room_GetInResp MsgCmd = 5
	// 进入准备场景
	MsgCmd_Room_EnterReadyReq  MsgCmd = 6
	MsgCmd_Room_EnterReadyResp MsgCmd = 7
	// 升级
	MsgCmd_Ready_UpgradeReq  MsgCmd = 8
	MsgCmd_Ready_UpgradeResp MsgCmd = 9
	// 开始游戏
	MsgCmd_Ready_StartGameReq  MsgCmd = 10
	MsgCmd_Ready_StartGameResp MsgCmd = 11
	// 攻击
	MsgCmd_Combat_AttackReq  MsgCmd = 12
	MsgCmd_Combat_AttackResp MsgCmd = 13
)

var MsgCmd_name = map[int32]string{
	0:  "Login_LoginReq",
	1:  "Login_LoginResp",
	2:  "Room_QueryListReq",
	3:  "Room_QueryListResp",
	4:  "Room_GetInReq",
	5:  "Room_GetInResp",
	6:  "Room_EnterReadyReq",
	7:  "Room_EnterReadyResp",
	8:  "Ready_UpgradeReq",
	9:  "Ready_UpgradeResp",
	10: "Ready_StartGameReq",
	11: "Ready_StartGameResp",
	12: "Combat_AttackReq",
	13: "Combat_AttackResp",
}

var MsgCmd_value = map[string]int32{
	"Login_LoginReq":      0,
	"Login_LoginResp":     1,
	"Room_QueryListReq":   2,
	"Room_QueryListResp":  3,
	"Room_GetInReq":       4,
	"Room_GetInResp":      5,
	"Room_EnterReadyReq":  6,
	"Room_EnterReadyResp": 7,
	"Ready_UpgradeReq":    8,
	"Ready_UpgradeResp":   9,
	"Ready_StartGameReq":  10,
	"Ready_StartGameResp": 11,
	"Combat_AttackReq":    12,
	"Combat_AttackResp":   13,
}

func (x MsgCmd) String() string {
	return proto.EnumName(MsgCmd_name, int32(x))
}

func (MsgCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8ab54014d815c68, []int{1}
}

func init() {
	proto.RegisterEnum("MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("MsgCmd", MsgCmd_name, MsgCmd_value)
}

func init() { proto.RegisterFile("cst.proto", fileDescriptor_b8ab54014d815c68) }

var fileDescriptor_b8ab54014d815c68 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x80, 0xa1, 0x40, 0xa1, 0xa3, 0xe8, 0x30, 0xf8, 0xf3, 0x0e, 0x1c, 0xbc, 0x18, 0x1f, 0xc0,
	0x10, 0x43, 0x4c, 0xe0, 0x60, 0xd5, 0x73, 0xb3, 0xd0, 0x4d, 0x43, 0x4c, 0xd9, 0xb1, 0x33, 0x1e,
	0xfa, 0x7c, 0xbe, 0x98, 0xd9, 0x31, 0x31, 0x52, 0x2e, 0x4d, 0xfa, 0xe5, 0x9b, 0x6f, 0x27, 0xbb,
	0x90, 0xed, 0x44, 0xef, 0xb8, 0x09, 0x1a, 0x16, 0x0f, 0x30, 0xde, 0x48, 0xf5, 0xd6, 0xb2, 0xa7,
	0x0c, 0x46, 0xeb, 0x50, 0xed, 0x0f, 0xd8, 0xa3, 0x09, 0x0c, 0xf3, 0x10, 0x6a, 0xec, 0x47, 0x98,
	0x7b, 0x57, 0xb6, 0x98, 0x10, 0x40, 0xba, 0x0c, 0xf5, 0xd6, 0x29, 0x0e, 0x16, 0xdf, 0x09, 0xa4,
	0x1b, 0xa9, 0x96, 0x75, 0x49, 0x04, 0x17, 0x36, 0x56, 0xd8, 0x37, 0xf7, 0x9f, 0xd8, 0xa3, 0x39,
	0x5c, 0x1e, 0x31, 0x61, 0xec, 0xd3, 0x35, 0xcc, 0x62, 0xb4, 0x78, 0xf9, 0xf2, 0x4d, 0xbb, 0xde,
	0x8b, 0x46, 0x37, 0xa1, 0x1b, 0xa0, 0x2e, 0x16, 0xc6, 0x01, 0xcd, 0x60, 0x6a, 0x7c, 0xe5, 0xf5,
	0xd9, 0xb2, 0xc3, 0x78, 0xd4, 0x7f, 0x24, 0x8c, 0xa3, 0xbf, 0xf1, 0xa7, 0x83, 0xfa, 0xc6, 0x56,
	0x8d, 0x6e, 0x4a, 0xb7, 0x30, 0x3f, 0xe1, 0xc2, 0x38, 0xa6, 0x2b, 0x40, 0xfb, 0x2d, 0xde, 0xb9,
	0x6a, 0x5c, 0xe9, 0xa3, 0x3e, 0xb1, 0xe5, 0x8e, 0xa9, 0x30, 0x66, 0x56, 0x37, 0xfc, 0xaa, 0xae,
	0xd1, 0x95, 0xab, 0x4d, 0x07, 0xab, 0x77, 0xb9, 0x30, 0x9e, 0xc5, 0xfa, 0xef, 0x25, 0x15, 0x8f,
	0xaa, 0x6e, 0xf7, 0x11, 0xf5, 0xf3, 0x58, 0xef, 0x50, 0x61, 0x9c, 0x6e, 0x53, 0x7b, 0x83, 0xfb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xd4, 0x9a, 0x60, 0x90, 0x01, 0x00, 0x00,
}
