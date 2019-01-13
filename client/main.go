package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	bb "github.com/Justyer/lingo/bytes"
	"github.com/Justyer/lingo/ip"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:9595", ip.MustInnerIP()))
	if err != nil {
		log.Fatal(err)
	}

	var req prt.LoginReq
	req.MAC = "1234"
	req.UserName = "dxc"
	req.PassWord = "111"
	data_byte, _ := proto.Marshal(&req)
	data_len := len(data_byte)
	len_byte := bb.ToByteForLE(uint32(data_len))

	cate1_byte := bb.ToByteForLE(uint16(1))
	cate2_byte := bb.ToByteForLE(uint16(2))

	final_byte := bb.Merge(cate1_byte, cate2_byte, len_byte, data_byte)
	fmt.Println(cate1_byte, cate2_byte, data_len, uint32(data_len), data_byte)
	conn.Write(final_byte)
	// conn.Write(final_byte)
	// for {
	// 	var x int
	// 	fmt.Scanf("%d", &x)
	// 	if x == 1 {
	// 		conn.Write(final_byte)
	// 	}
	// }

	//读取到buffer
	buf := make([]byte, 512)
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
