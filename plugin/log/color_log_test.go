package log

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	buf_pool := make([]byte, 10)
	buf := make([]byte, 5)
	// bytes.ToByte(binary.LittleEndian, 1)

	// buf_pool = bytes.ToByte(binary.LittleEndian, uint16(1))

	fmt.Println(buf_pool, buf)
	fmt.Println(len(buf_pool))
}
