// 方法与函数
package main

import (
	"fmt"

	geo "gopl.io/ch6/geometry"
)

//Path 路径
type Path []geo.Point

//Distance 距离计算
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := geo.Point{1, 2}
	q := geo.Point{3, 4}

	//函数调用
	fmt.Println(geo.Distance(p, q))
	//方法调用
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
	z := &perim
	fmt.Println(z.Distance()) //编译器会隐式的转换为 *z 进行调用

	//指针类型方法调用
	s := geo.Point{1, 2}
	r := &s
	s.ScaleBy(2) //编译器会隐式的转换为 &s 进行调用
	r.ScaleBy(2)
	fmt.Println(*r)
	s.ScaleBy(2)
	/*
		1. 不管方法调用的接收者是指针类型还是非指针类型，都可以通过指针/非指针类型进行调用，编译器帮你做类型转换
		2. 声明方法的接收者是指针还是非指针，需要考虑对象的大小
	*/

}
