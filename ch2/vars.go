//  变量
package main

import "fmt"

func main() {
	// 如果没有初始化值，将用零值初始化该变量
	var s string // 这里将是""
	fmt.Println(s)
	var i, j, k int                  // 声明和初始化多个变量
	var b, f, s1 = true, 2.3, "four" // 声明多个不同类型的变量
	// 函数内的简短声明方式，简短声明变量。
	// 简短声明变量和多重赋值一起使用时，简单声明变量至少要声明一个变量
	// 否则只能改用多重赋值
	t := 0.0
	fmt.Printf("%v %v %v %v %v %v %v\n", i, j, k, b, f, s1, t)

	i = 0
	// 是语句，不是表达式，没有返回值
	i ++ 						
	fmt.Printf("%v", i)

	// 多赋值时可以用下划线丢弃不需要的值
	_, a1, a2 := 1, 2, 3
	// 输出的是a1是12 (是bug吗？)
	fmt.Printf("%d, %d\n",a1 , a2)

	b1, b2, _, b4 := 1,2,3,4
	fmt.Printf("%d, %d, %d\n",b1 , b2, b4)

	// 在向函数传参或者给slice的元素赋值等操作时都会发生隐式赋值。
	// 左右两边需要有相同的数据类型。
	// 用 == 或 != 进行相等性检查时，也和可赋值能力有关，第二个值必须对第一个值的类型是可赋值的，以之亦然。
	
}
