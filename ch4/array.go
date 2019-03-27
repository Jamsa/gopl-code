// 数组
/*
 * 数组长度固定，Slice可增长收缩
 */
package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	q := [3]int{1,2,3}
	idx := 100;
	//fmt.Printf("%v",q[idx])		// 越界访问panic
	r := [...]int{1,2,3,4,5}	// 长度自动设置
	// q = r						// 长度不同，无法赋值
	fmt.Printf("数组\n")
	fmt.Printf("%T\n",q)
	fmt.Printf("%T\n",r)
	// 默认使用顺序索引，也可以指定索引和值的方式初始化
	symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}
	fmt.Println(RMB, symbol[RMB])
	// 字面值形式中，初始化索引的顺序是无关紧要的，没有用到的索引可以省略
	r1 := [...]int{99: -1}		// 100个元素的数组，最后一个值是-1，其余默认为零值
	fmt.Printf("%T\n",r1)

	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{2,3}
	fmt.Println(a == b, a==c ,b==c)			// 长度和元素完全相等
}

/*
 * 函数调用时，Go语言是传值调用，参数变量是一个复制的副本。 
 * Go语言并不会像有些语言那样，隐式的将数组作为引用或指针传入。
 * 因此，传递大数组时可以显式的传入数组指针，以便将对数组的修改反馈给调用者。
 * 因为数组元素数量固定，并不能通过指针修改参数数组长度，所以Slice作为参数更为常见。
 */
func zero(ptr *[32]byte)  {
	*ptr = [32]byte{}
	/*或
    for i := range ptr {
		ptr[i] = 0
	}*/
}
