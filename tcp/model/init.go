package model

import (
	"net"

	"github.com/jinzhu/gorm"
)

var (
	JekoDB *gorm.DB
)

var (
	// 服务类型：http/tcp
	ServeType string

	// tcp监听的ip:port
	TCPIPort string

	// http监听的ip:port
	HTTPIPort string
)

var (
	LinkConns []net.Conn
)

var (
// 存储用户的所有信息的汇总,key为username
// GlobalAcountInfo map[string]user.AcountInfo
)
