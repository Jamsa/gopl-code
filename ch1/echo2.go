// Echo2 显示它的命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 使用range 产生索引和值对, 从Args[1]开始忽略程序本身
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	
	fmt.Println(s)

	s, sep = "", ""
	for idx, arg := range os.Args {
		s += sep + fmt.Sprintf("%d:", idx) + arg
		sep = " "
	}
	
	fmt.Println(s)
}
