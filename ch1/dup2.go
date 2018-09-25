// Dup1查找重复的行
// 从标准输入或文件列表中读取

package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err) // %v为变量的自然形式
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	
	for line,n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines ...
func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[f.Name()+":"+input.Text()]++
	}
}
