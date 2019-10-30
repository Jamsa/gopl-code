//通过channel获取键盘输入中止执行
package main

import "os"

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取一个byte
		abort <- struct{}{}            //如果读取到了就发送消息
	}()
}
