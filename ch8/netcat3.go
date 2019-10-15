//使用不带缓存的channel的版本
//无缓存channels的发送操作将导致发送者goroutine阻塞，直到国一个goroutine
//在相同的channels上执行接收操作，当发送的值通过channels成功传输之后，两个
//goroutine才能继续执行后面的语句。反之如果接收操作先发生，那么接收者goroutine
//也将阻塞，直到另一个goroutine在相同的channels上执行发送操作。
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //通知主goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等侍后台groutine结束
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//带缓存的channels在达到最大容量时发送端将阻塞。
//channels有可能出现因为没有人接收而被永远卡，这种情况被称为goroutine泄漏。
//泄漏的goroutines将不会被回收。因此需要确保每个不再需要的goroutine能正常退出。
//泄漏的场景：使用多个goroutines向无缓存的channels发送数据，如果程序在接收第一个
//数据后就返回，那么发送速度慢的goroutines就都会被阻塞，因为他们没有机会发送数据。
//这些goroutine也不会被销毁。
