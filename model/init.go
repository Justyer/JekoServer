package model

import (
	"github.com/jinzhu/gorm"
)

// 数据库
var (
	JekoDB *gorm.DB
)

// 命令
var (
	// 子命令类型
	TaskType string

	// 服务类型：http/tcp
	ServeType string

	// flag -t 类型
	Flag_T_Type string
)

// 网络
var (
	// 内网IP
	IP string

	// tcp监听的ip:port
	TCPPort string

	// http监听的ip:port
	HTTPPort string
)
