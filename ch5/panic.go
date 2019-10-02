/*
* panic异常
* panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer机制）。
* 随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。
* 调用panic函数可以产生panic异常
*
 */
package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) //除零
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//控制堆栈输出
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	defer printStack()
	f(3)
}
