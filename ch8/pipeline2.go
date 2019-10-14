// 将多个goroutines连接起来，构成了pipeline
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		//更复杂的应用时，可以在defer中关闭channels
		close(naturals)
	}()

	// Squarer
	go func() {
		// range可以直接在goroutine中迭代
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// 主goroutine中输出结果
	for x := range squares {
		fmt.Println(x)
	}
}
