package base

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/model/tool"

	ptl "github.com/Justyer/JekoServer/plugin/protocol"
)

type BaseController struct {
	DataPack *ptl.DataPack_2_2_4
	Conn     net.Conn
	Cache    *tool.Cache
}
