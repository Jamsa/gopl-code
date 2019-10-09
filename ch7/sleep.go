/*
* flag.Value接口使用示例
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

// 用 -period 控制休眠时间，默认值为1秒，参数帮助信息为sleep period
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
