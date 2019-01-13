package tcp

import (
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/JekoServer/tcp/router"
	"github.com/Justyer/jie"
	protocol "github.com/Justyer/jie/protocol/ptl_2_2_4"
)

// TCPServe : TCP服务
func TCPServe() {
	j := jie.New()

	j.SetProtocol(protocol.NewProtocol())

	j.SetRouter(protocol.NewRouter())

	router.CreateJekoRouter(j)

	j.ListenAndInnerServe(model.TCPPort)
}
