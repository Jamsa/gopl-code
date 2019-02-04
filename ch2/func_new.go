package main

import ("fmt")

// new的作用是创建一个匿名变量，初始化为零值
// 只是个语法糖，与定义变量后返回其指针并没有区别

// newInt ...
func newInt1() *int {
	return new(int)
}

// newInt ...
func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	// 每次返回的变量地址是不同的
	p := new(int)
	q := new(int)
	fmt.Println(p == q)

	// 如果两个类型是空的，即类型的大小是0，如struct{}和[0]int
	// 则它们有可能有相同的地址（依赖于语言实现）

	// new是预定义函数，不是关键字
}
