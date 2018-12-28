package tcp

import (
	"fmt"
	"net"

	"github.com/Justyer/JekoServer/model"
)

// 开始一个TCP服务
func TCPServe() {
	listener, err := net.Listen("tcp", model.TCPIPort)
	if err != nil {
		fmt.Printf("[Error listening]: %s", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("[Error accepting]: %s\n", err.Error())
			return // 终止程序
		}

		tcpConn, ok := conn.(*net.TCPConn)
		if !ok {
			//如果断言失败，放弃连接
			fmt.Println("assert fail")
			continue
		}
		tcpConn.SetNoDelay(false)

		fmt.Printf("[link-add-iport]: (%s)\n", conn.RemoteAddr().String())
		model.LinkConn = append(model.LinkConn, conn)
		go JustDoIt(conn)
	}
}

// 每个client连接后的处理
func JustDoIt(conn net.Conn) {
	defer conn.Close()
	// buf_pool := make([]byte, 1024)
	// for {
	// 	buf := make([]byte, 512)
	// 	l, err := conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Printf("[Error reading]: %s, on (%s)\n", err.Error(), conn.RemoteAddr().String())
	// 		return
	// 	}

	// 	// 将本次读取的字节放到字节池里
	// 	for _, b := range buf {
	// 		buf_pool = append(buf_pool, b)
	// 	}

	// 	var datapack []byte

	// 	// 如果字节长度大于6，才可能是完整的数据包（不是一定）
	// 	if l > 6 {
	// 		var dl int
	// 		dl_buf := bytes.NewBuffer(buf[4:6])
	// 		binary.Write(dl_buf, binary.BigEndian, &dl)
	// 		// 如果缓冲区接收的字节长度大于等于数据包长度，处理该数据
	// 		if l >= dl+6 {
	// 			datapack = buf[:dl+6]

	// 			// 大分类
	// 			big_cate := datapack[0:2]
	// 			// 小分类
	// 			sml_cate := datapack[4:4]

	// 			// 数据
	// 			data := datapack[6 : dl+6]
	// 		}
	// 	}

	// }
}
