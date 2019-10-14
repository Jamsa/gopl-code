// goroutine示例
package main

import (
	"fmt"
	"time"
)

func main() {
	// go语句创建goroutine后立即返回
	go spinner(100 * time.Millisecond)
	const n = 45

	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d)=%d\n", n, fibN)
	// 主函数退出时，所有goroutine都会被直接打断
}

func spinner(delay time.Duration) {
	// 这里的死循环在主函数退出时将被打断
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
