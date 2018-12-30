package ptl

import (
	"github.com/Justyer/lingo/bytes"
)

type DataPack_2_2_4 struct {
	OgnByte []byte
	MsgType uint16
	MsgCmd  uint16
	DataLen uint32
	Data    []byte
}

func NewDataPack_2_2_4() *DataPack_2_2_4 {
	return &DataPack_2_2_4{}
}

// 判断字节数组是否完整
// 这个方法只能判断这个字节数组是否可以解析，不能保证一定包含这个协议
func DP_2_2_4_IsWhole(buf []byte) bool {
	buf_len := len(buf)
	// 如果字节长度大于8，才可能是完整的数据包（不是一定）
	if buf_len < 8 {
		return false
	}
	var len_i uint32
	bytes.ByteToForLE(buf[4:8], &len_i)
	if buf_len < int(len_i)+8 {
		return false
	}
	return true
}

// 解析数据包
func (self *DataPack_2_2_4) Parse() {
	var type_b, cmd_b, len_b []byte

	for i := 0; i < 8; i++ {
		switch i {
		case 0, 1:
			type_b = append(type_b, self.OgnByte[i])
		case 2, 3:
			cmd_b = append(cmd_b, self.OgnByte[i])
		case 4, 5, 6, 7:
			len_b = append(len_b, self.OgnByte[i])
		}
	}

	// 解析大分类和小分类
	bytes.ByteToForLE(type_b, &self.MsgType)
	bytes.ByteToForLE(cmd_b, &self.MsgCmd)
	bytes.ByteToForLE(len_b, &self.DataLen)

	self.Data = self.OgnByte[8 : self.DataLen+8]
}
