// 将多个goroutines连接起来，构成了pipeline
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 主goroutine中输出结果
	for {
		fmt.Println(<-squares)
	}
}
