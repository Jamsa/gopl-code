/**
函数值
*/
package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func add1(r rune) rune { return r + 1 }

//匿名函数可访问完整的词法环境
func square2() func() int {
	var x int //闭包
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// f = product 报错，方法签名不一致

	var f1 func(int) int //函数的零值是nil
	// f1(3) 这行会报错
	if f1 != nil { // 可以与 nil 进行比较。但是函数间不可比较，也不能用函数值作为map的key
		f1(3)
	}

	//对字符串的每个字符都调用add1
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))

	//匿名函数
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))

	f2 := square2()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}
