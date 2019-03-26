// 作用域
package main

import (
	"fmt"
)

func f() int {
	return 0
}

func g(x int) int {
	return x
}

var g1 = "g"

func main() {
	f1 := "f"
	fmt.Println(f1)				// 局部变量
	fmt.Println(g1)				// 包级变量
	//fmt.Println(h)				// 未定义
	test()

	if x := f(); x == 0 {		// if也会创建新作用域
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x,y)
	}
	// fmt.Println(x, y)			// 编译错误,x,y都未定义
}

// 并不是所有的语法显式的对应到由花括号包含的语句
// if、switchr 的条件部分和执行体会创建语法块
// for 循环的初始化部分和循环体部分都会。
// 下面的三个x是不同的
func test() {
	x := "hello!"				// 第一个
	for i := 0; i< len(x); i++ {
		x := x[i]				// 第二个
		if x != '!' {
			x := x + 'A' - 'a'	// 转大写 第三个 := 左边是个新变量
			fmt.Printf("%c", x)
		}
	}
}


