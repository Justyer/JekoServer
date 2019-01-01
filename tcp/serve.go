package tcp

import (
	"io"
	"net"
	"os"

	"github.com/Justyer/JekoServer/tcp/router"

	"github.com/Justyer/JekoServer/plugin/log"
	ptl "github.com/Justyer/JekoServer/plugin/protocol"
	"github.com/Justyer/JekoServer/tcp/model"
	msg "github.com/Justyer/JekoServer/tcp/model/message"
	"github.com/Justyer/JekoServer/tcp/model/tool"
	"github.com/Justyer/JekoServer/tcp/model/user"
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
	log.Info("[listener]: (%s)", model.TCPIPort)

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

	msg := msg.NewMsg()
	msg.Conn = conn

	var c tool.Cache
	c.User = &user.User{}

	for {
		if err := msg.Read(50); err != nil {
			if err == io.EOF {
				log.Err("[link-close-iport]: (%s)", msg.Conn.RemoteAddr().String())
				return
			}
			log.Err("[Error reading]: %s, on (%s)", err.Error(), msg.Conn.RemoteAddr().String())
			return
		}

		dp := ptl.NewDataPack_2_2_4()
		dp.OgnByte = msg.BufPool

		if ptl.DP_2_2_4_IsWhole(msg.BufPool) {
			dp.Parse()
			r := router.NewJekoRouter()
			r.DataPack = dp
			r.Conn = msg.Conn
			r.Cache = &c
			r.MsgType()
		}
		msg.HeadCut(int(dp.DataLen) + 8)

	}
}
