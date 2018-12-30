package msg

import (
	"net"

	"github.com/Justyer/lingo/bytes"
)

type Message struct {
	Conn    net.Conn
	BufPool []byte
}

func NewMsg() *Message {
	return &Message{}
}

func (self *Message) Read(bs int) error {
	buf := make([]byte, bs)
	l, err := self.Conn.Read(buf)
	if err != nil {
		return err
	}

	// 去掉多余的0字节
	buf = buf[:l]

	self.BufPool = bytes.Extend(self.BufPool, buf)

	return nil
}

func (self *Message) HeadCut(l int) {
	self.BufPool = self.BufPool[l:]
}
