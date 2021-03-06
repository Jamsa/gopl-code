// 使用包的示例，使用了tempconv包中的两个go文件中定义的类型和方法
package main

import (
	"fmt"
	"os"
	"strconv"
	
	"gopl.io/ch2/tempconv"
)

// 包初始化时按变量声明的顺序依次初始化。包上含有多个源文件时按发送给编译器的顺序初始化，Go语言构建工具会先将.go文件按文件名硫，然后依次调用编译器编译。
// 每个文件都可以包含多个init初始化函数。这个函数不能被调用或引用。它会被按导入声明的顺序执行。每个包只会被初始化一次，main包是最后被初始化的。
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
