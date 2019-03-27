/*
Slice
代表变长的序列，序列中每个元素的类型都相同。一般写作[]T，与数组类似，只是没有固定长度。
它提供了访问数组子序（或全部）元素的功能，它的底层引用了一个数组对象。
slice由三部分组成：指针、长度和容量。指针指向slice（不是底层数组）的第一个元素，长度不能超过容量。
多个slice间可共享底层的数组，并且引用的数组区间可重叠。
切片操作s[i:j]，0<=i<=j<=caps，将创建一个新的slice，引用从第i个元素开始至第j-1间的j-i个元素。省略i时将用0代替，省略j时将用len(s)代替。s[:]引用整个数组。
超过caps将panic，发出len(s)则是扩展slice。
slice间不能比较，标准库提供了高度优化的bytes.Equal函数来判断两个[]byte是否相等。其它类型的，必须自己展开每个元素进行比较。唯一合法的比较操作是和nil比较。
创建方式make([]T, len)或make([]T, len,cap)，它返回的slice对象的底层有一个匿名数组。

尽管底层数组的元素是间接访问的，但是slice对应结构体本身的指针、长度和容量部分是直接访问的。从这个角度看，它并不是一个纯粹的引用类型。
*/
package main

import (
	"fmt"
)

func main() {
	months := [...]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	fmt.Println(months[:13])
	fmt.Println(months[:])
	m := months[:5]
	fmt.Println(m)
	//fmt.Println(months[:14]) 	// panic

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x,y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i) 	// 可直接写成x = appendInt(x, i)
		fmt.Printf("%d cap=%d \t %v\n", i, cap(y), y)
		x = y
	}

	fmt.Printf("%d\n", len(noneempty([]string{"Hello","","","world!"})))

	stackTest()
}

// appendInt 函数演示内置函数append函数的工作方式
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 容量足够
		z = x[:zlen]
	} else {
		// 空间不足。分配新数组
		// 分配双倍大小，避免多次操作时重复分配内存
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		// 内置的slice复制函数，z是目标,x是源，它的返回值是成功复制的元素的个数，等于两者中较小的slice的长度。
		copy(z, x)				
	}
	z[len(x)] = y
	return z
}

// 内存技巧一：通过合并元素避免重新分配内存
// noneempty 返回非空字符串
func noneempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i ++
		}
	}
	return strings[:i]
}

// 内存技巧二：模拟stack和利用copy和切片操作来删除中间的元素
// stack ...
func stackTest()  {
	strings := []string{"a","b","c"}
	
	stack := append(strings, "d") // push
	fmt.Printf("%v\n",stack)
	
	top := stack[len(stack)-1]
	fmt.Printf("%v\n",top)

	stack = stack[:len(stack)-1] // pop
	fmt.Printf("%v\n",stack)

	i := 1						// remove 索引为i的元素
	copy(stack[i:], stack[i+1:])
	stack = stack[:len(stack) - 1] // 去掉最后一个元素
	fmt.Printf("%v\n",stack)
}
