package serve

import (
	"log"

	"github.com/Justyer/JekoServer/http"
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/JekoServer/tcp"
)

// Run : 运行serve
func Run() {
	switch model.Flag_T_Type {
	case "tcp":
		tcp.TCPServe()
	case "http":
		http.HTTPServe()
	default:
		log.Fatal("flag:-t not found")
		// wg := &sync.WaitGroup{}
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	http.HTTPServe()
		// }()
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	tcp.TCPServe()
		// }()
		// wg.Wait()
	}
}
