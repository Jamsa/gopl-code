// 聊天室示例
// 客户端可以用nc 连接 9000端口
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string //发送的消息的channel

var (
	entering = make(chan client) //客户端进入
	leaving  = make(chan client) //客户端离开
	messages = make(chan string) //所有来自客户端的消息
)

func broadcaster() {
	// 与多个goroutine相关的这个变量，不需要锁，只为它被限定在 broadcaster 内
	clients := make(map[client]bool) //连接的客户端
	for {
		select {
		case msg := <-messages:
			//将消息向所有客户端广播
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "你是 " + who
	messages <- who + " 已经进入"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	//注意：忽略了input.Err()

	leaving <- ch
	messages <- who + " 已经离开"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) //注意：忽略了网络错误
	}
}
