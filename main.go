package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/Justyer/lingo/util"
)

var (
	ip  string
	all []net.Conn
	chl chan int

	blood = 300
)

func init() {
	var err error
	ip, err = util.LocalIP()
	if err != nil {
		log.Fatal(err)
	}

	chl = make(chan int, 1000)
}

func deal() {
	for m := range chl {
		// dela
		fmt.Println("fuck:", m)
	}
}

func attack(w http.ResponseWriter, r *http.Request) {
	blood--
	fmt.Fprintf(w, "[http-> monster blood surplus]: %d\n", blood)
}

func resume(w http.ResponseWriter, r *http.Request) {
	blood++
	fmt.Fprintf(w, "[http-> monster blood add]: %d\n", blood)
}

func main() {
	go func() {
		http.HandleFunc("/attack", attack)
		http.HandleFunc("/resume", resume)
		http.ListenAndServe(fmt.Sprintf("%s:9596", ip), nil)
	}()

	go deal()

	fmt.Printf("--------------------------------\n")
	fmt.Printf("Starting the server on: %s\n", ip)
	fmt.Printf("--------------------------------\n\n")
	// 创建 listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:9595", ip))
	if err != nil {
		fmt.Printf("[Error listening]: %s", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("[Error accepting]: %s\n", err.Error())
			return // 终止程序
		}
		fmt.Printf("[link-add-iport]: (%s)\n", conn.RemoteAddr().String())
		all = append(all, conn)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		l, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("[Error reading]: %s, on (%s)\n", err.Error(), conn.RemoteAddr().String())
			return //终止程序
		}
		resp := buf[:l]
		fmt.Printf("[Received data]: %v\n", string(resp))
		// for i := 0; i < l-1; i++ {
		// 	resp[i], resp[l-1-i] = resp[l-1-i], resp[i]
		// }
		blood--
		chl <- blood
		fmt.Printf("[tcp-> monster blood surplus]: %d\n", blood)
		for i := 0; i < len(all); i++ {
			all[i].Write([]byte(strconv.Itoa(blood)))
		}
	}
}
