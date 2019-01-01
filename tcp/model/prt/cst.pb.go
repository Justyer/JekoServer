// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cst.proto

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
	proto.RegisterEnum("prt.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("prt.MsgCmd", MsgCmd_name, MsgCmd_value)
}

func init() { proto.RegisterFile("cst.proto", fileDescriptor_b8ab54014d815c68) }

var fileDescriptor_b8ab54014d815c68 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0x40, 0xfb, 0x9b, 0x36, 0xf3, 0x7d, 0xd5, 0xe9, 0xd4, 0x9f, 0x77, 0xe8, 0x85, 0x37, 0xe2,
	0x03, 0x48, 0x91, 0x22, 0xb4, 0x17, 0x46, 0xbd, 0x0e, 0xdb, 0x66, 0x09, 0x45, 0xd2, 0x1d, 0x77,
	0xc6, 0x8b, 0x3c, 0x9f, 0x2f, 0x26, 0x3b, 0x82, 0xd8, 0x78, 0x13, 0xc8, 0xe1, 0xcc, 0xd9, 0x61,
	0x17, 0xf2, 0xbd, 0xe8, 0x0d, 0xc7, 0xa0, 0x81, 0x86, 0x1c, 0x75, 0x79, 0x07, 0x93, 0xad, 0xd4,
	0x2f, 0x2d, 0x7b, 0xca, 0x61, 0xbc, 0x09, 0xf5, 0xe1, 0x88, 0x3d, 0x9a, 0xc2, 0xa8, 0x08, 0xa1,
	0xc1, 0x7e, 0x82, 0x85, 0x77, 0x55, 0x8b, 0x03, 0x02, 0xc8, 0x56, 0xa1, 0xd9, 0x39, 0xc5, 0xe1,
	0xf2, 0x73, 0x00, 0xd9, 0x56, 0xea, 0x55, 0x53, 0x11, 0xc1, 0x99, 0x8d, 0x95, 0xf6, 0x2d, 0xfc,
	0x3b, 0xf6, 0x68, 0x01, 0xe7, 0x27, 0x4c, 0x18, 0xfb, 0x74, 0x09, 0xf3, 0x14, 0x2d, 0x9f, 0x3e,
	0x7c, 0x6c, 0x37, 0x07, 0xd1, 0xe4, 0x0e, 0xe8, 0x0a, 0xa8, 0x8b, 0x85, 0x71, 0x48, 0x73, 0x98,
	0x19, 0x5f, 0x7b, 0x7d, 0xb4, 0xec, 0x28, 0x1d, 0xf5, 0x1b, 0x09, 0xe3, 0xf8, 0x67, 0xfc, 0xe1,
	0xa8, 0x3e, 0xda, 0xaa, 0xc9, 0xcd, 0xe8, 0x1a, 0x16, 0x7f, 0xb8, 0x30, 0x4e, 0xe8, 0x02, 0xd0,
	0x7e, 0xcb, 0x57, 0xae, 0xa3, 0xab, 0x7c, 0xd2, 0xa7, 0xb6, 0xdc, 0x29, 0x15, 0xc6, 0xdc, 0xea,
	0x86, 0x9f, 0xd5, 0x45, 0x5d, 0xbb, 0xc6, 0x74, 0xb0, 0x7a, 0x97, 0x0b, 0xe3, 0xbf, 0x54, 0xff,
	0xbe, 0xa4, 0xf2, 0x5e, 0xd5, 0xed, 0xdf, 0x92, 0xfe, 0x3f, 0xd5, 0x3b, 0x54, 0x18, 0x67, 0xbb,
	0xcc, 0x1e, 0xe2, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xea, 0x83, 0x65, 0x95, 0x01, 0x00,
	0x00,
}