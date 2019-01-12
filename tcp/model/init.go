package model

import (
	"github.com/jinzhu/gorm"
)

var (
	JekoDB *gorm.DB
)

var (
	// 服务类型：http/tcp
	ServeType string

	// 内网IP
	IP string

	// tcp监听的ip:port
	TCPPort string

	// http监听的ip:port
	HTTPPort string
)
