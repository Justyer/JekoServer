package tcp

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/Justyer/JekoServer/model/auth"

	"github.com/Justyer/JekoServer/extension/log"
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/lingo/bytes"
)

// 开始一个TCP服务
func TCPServe() {
	// 启动监听
	listener, err := net.Listen("tcp", model.TCPIPort)
	defer listener.Close()
	if err != nil {
		log.Err("[Error listening]: %s", err.Error())
		return
	}
	log.Info("[listener]: (%s)\n", model.TCPIPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Err("[Error accepting]: %s", err.Error())
			os.Exit(1)
		}

		log.Info("[link-start-iport]: (%s)", conn.RemoteAddr().String())
		model.LinkConns = append(model.LinkConns, conn)
		go GoLink(conn)
	}
}

// 每个client连接后的处理
func GoLink(conn net.Conn) {
	defer conn.Close()

	var buf_pool []byte

	for {
		buf := make([]byte, 50)
		l, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Err("[link-close-iport]: (%s)", conn.RemoteAddr().String())
				return
			}
			log.Err("[Error reading]: %s, on (%s)", err.Error(), conn.RemoteAddr().String())
			return
		}
		// 去掉多余的0
		buf = buf[:l]

		buf_pool = bytes.Extend(buf_pool, buf)

		// 如果字节长度大于8，才可能是完整的数据包（不是一定）
		if l > 8 {
			log.Info("[1if_len]: %d", l)

			type_byte, cmd_byte, len_byte := ChaiByte_2_2_4(buf_pool, l)

			var dl uint32
			bytes.ByteToForLE(len_byte, &dl)
			// 如果缓冲区接收的字节长度大于等于数据包长度，处理该数据
			if l >= int(dl)+8 {
				log.Tx("[datapack]: %v", buf_pool[:dl+8])

				dp := DataPack{}
				dp.MsgType = type_byte
				dp.MsgCmd = cmd_byte
				dp.LenByte = len_byte
				dp.Len = int(dl)
				dp.Data = buf_pool[8 : dl+8]

				Cate1(dp, conn)

				buf_pool = buf[dl+8 : l]
			} else {
				fmt.Println("2if", dl)
			}
		} else {
			fmt.Println("1if", l)
		}

	}
}

type DataPack struct {
	MsgType []byte
	MsgCmd  []byte
	LenByte []byte
	Len     int
	Data    []byte
}

func ChaiByte_2_2_4(buf []byte, l int) ([]byte, []byte, []byte) {
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
	return big_cate_byte, sml_cate_byte, len_cate_byte
}

func Cate1(dp DataPack, conn net.Conn) {

	// 解析大分类和小分类
	var type_short uint16
	bytes.ByteToForLE(dp.MsgType, type_short)
	var cmd_short uint16
	bytes.ByteToForLE(dp.MsgCmd, cmd_short)

	// 解析数据
	var req auth.LoginReq
	proto.Unmarshal(dp.Data, &req)
	log.Tx("[recv]: %v %v %d %s", type_short, cmd_short, dp.Len, req.String())

	// 封装要相应的数据
	var resp auth.LoginResp
	resp.Code = 0
	resp.UserName = "zxy"
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, dp.MsgType)
	resp_final_byte = bytes.Extend(resp_final_byte, dp.MsgCmd)
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	_, err = conn.Write(resp_final_byte)
	if err != nil {
		log.Err("write err:", err)
	}
	log.Tx("resp_final_byte: %v", resp_final_byte)
}
