//package tempconv
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius  = 0
	BoilingC Celsius = 100
)

// CtoF ...
func CtoF(c Celsius) Fahrenheit {
	// 每个类型T都会有个类型转换操作T(x)，用于将x转为T类型。
	// 指针类型可能需要用小括号包装T，如: (*int)(0)。
	// 只有当从个类型的底层基础类型相同时，才允许这种转型操作。

	return Fahrenheit(c * 9 / 5 + 32)
}

// FtoC ...
func FtoC(f Fahrenheit) Celsius {
	return Celsius( (f-32) * 5 / 9)
}

func test(){
	// 如果x是可以赋值给T类型的值，那么x必然可以被转为T类型，但是一般不需要
	boilingF := CtoF(100)
	fmt.Printf("%g\n", boilingF - CtoF(FreezingC))

	fmt.Printf("%g\n", BoilingC - FreezingC)

	// == 和 < 也可以用于比较命名类型的变量和另一个相同类型的变量，或者有相同底层类型的未命名类型的值间做比较。量是如果两个值的命名类型不同，则不能直接比较
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f)  //编译错误
	fmt.Println(c == Celsius(f))
}

func main() {
	test()
}
