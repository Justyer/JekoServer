package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"unsafe"

	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.100:9595")
	if err != nil {
		log.Fatal(err)
	}

	var req prt.LoginReq
	req.MAC = "1234"
	data_byte, _ := proto.Marshal(&req)
	data_len := len(data_byte)
	len_byte := Uint322Byte(uint32(data_len))

	cate1_byte := []byte{
		byte(uint16(0)),
		byte(uint16(0) >> 8),
	}
	cate2_byte := []byte{
		byte(uint16(0)),
		byte(uint16(0) >> 8),
	}

	final_byte := BytesCombine(cate1_byte, cate2_byte, len_byte, data_byte)
	// final_byte := BytesCombine(cate1_byte, cate2_byte)
	conn.Write(final_byte)
	// for {
	// 	var x int
	// 	fmt.Scanf("%d", &x)
	// 	if x == 1 {
	// 		conn.Write(final_byte)
	// 	}
	// }

	//读取到buffer
	buf := make([]byte, 20)
	//如果服务端没有把数据传递过来，那么客户端阻塞，直到服务端向其中写入了数据。
	l, _ := conn.Read(buf)
	fmt.Println(buf)

	var big_cate_byte, sml_cate_byte, len_cate_byte []byte

	for i := 0; i < l; i++ {
		switch i {
		case 0, 1:
			big_cate_byte = append(big_cate_byte, buf[i])
		case 2, 3:
			sml_cate_byte = append(sml_cate_byte, buf[i])
		case 4, 5, 6, 7:
			len_cate_byte = append(len_cate_byte, buf[i])
		}
	}

	var dl uint32
	dl_buf := bytes.NewBuffer(len_cate_byte)
	binary.Read(dl_buf, binary.LittleEndian, &dl)
	data := buf[8 : dl+8]
	fmt.Println("dl_buf", dl_buf.Bytes(), data)

	var resp prt.LoginResp
	proto.Unmarshal(data, &resp)
	fmt.Printf("%#v", resp)
}
func Uint322Byte(data uint32) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp uint32 = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 4) & data) >> (index * 4))
	}
	return ret
}

func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
