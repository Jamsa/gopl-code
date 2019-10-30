// 下面的示例演示的是发射倒计时
// 它使用了tick和abort两个channel分别作为倒计时器和中止信号
// 为了从两个channel中获取事件，我们需要使用select语句实现多路复用

// channel的零值是nil。对nil的channel发送和接收操作会永远阻塞。
// 在select语句中操作nil的channel将永远不会select到。

// select会等待case中有能执行的case时去执行。当条件满足时，select才会通信并执行case后的语句；
// 这个时候其它通信不会执行。一个没有任何case的select语句会永远地等待下去。
// 如果多个case同时就绪，select会随机地选择一个执行，这样来保证每个channel都 平等的被select的机会。
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launch!")
}

func main() {
	fmt.Println("Commencing coutdown. Press return to abort.")

	// time.Tick 函数将定期向它返回的channel发送消息。当函数返回时它将继续存活，造成goroutine泄漏，
	// 它适用于程序整个生命周期都需要的计数器
	tick := time.Tick(1 * time.Second)

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取一个byte
		abort <- struct{}{}            //如果读取到了就发送消息
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case x := <-tick:
			fmt.Println(x)
		case <-abort: //这种写法忽略channel中读取的值
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}
