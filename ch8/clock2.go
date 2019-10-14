// 支持多接的时钟服务
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 连接中断等情况
			continue
		}
		go handleConn(conn) // 只修改了这行，一次处理一个连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 客户端断开等情况
		}
		time.Sleep(1 * time.Second)
	}
}
