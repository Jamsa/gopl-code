//channel实现倒计时
package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Launch!")
}

func main() {
	fmt.Println("Commencing coutdown.")

	tick := time.Tick(1 * time.Second) // 会周期性的向tick发送事件，每个事件一个时间戳
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j := <-tick
		fmt.Println(j)

	}
	launch()
}
