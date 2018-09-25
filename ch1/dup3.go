// Dup1查找重复的行
// 从文件列表中读取，读取文件内容时，整体读取

package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup2:%v\n", err) // %v为变量的自然形式
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line] ++
		}
	}

	
	for line,n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
