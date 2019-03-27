/*
Map是无序的key/value集合，写作map[K]V，K和V的类型可以不同，但是所有的key有相同的类型,value也有相同的类型。
key必须是支持 == 比较运算符的类型。虽然浮点也支持相等运算，但是用浮点做key类型是不合适的，因为它可能为NaN，NaN与任何浮点类型都不相等。

元素可以用自增运算符，不存在的key会创建。
元素不是变量，不能取址。因为map元素可能会随元素数量的增长而重新分配更内存空间，导致地址失效。

用range迭代时顺序是不确定的，语言故意为之，强制要求程序不依赖于迭代顺序。如果需要按顺序遍历key/value，应该对key进行排序。如，使用sort包的Strings函数对字符串slice排序。

map类型的零值是nil，可以调用查找、删除、len和range等，但是不能向它存入元素

通过map[key]访问时，key不存在时将得到value对应类型的零值。如果value本身存在零值的情况，可以用v,ok := map[key]的方式，检查布尔型变量ok的值。

与slice一样，map间也不能进行相等性比较，除非与nil进行比较。

golang中没有set，可以用value为bool的map代替。

当需要slice这类无法使用 == 进行比较的类型作为key时，可以定义专用函数将slice转为string类型的，以string值作为map的key。
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 字面值
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}

	fmt.Printf("%v\n",ages)

	// 用make创建
	ages = make(map[string]int)
	ages["alice"] = 35
	ages["aaa"] = 36
	fmt.Printf("%v\n",ages)

	// 删除元素
	delete(ages,"aaa")
	fmt.Printf("%v\n",ages)

	// 查找不到时将返回value类型的零值
	fmt.Printf("%v\n",ages["ccc"])

	// 可以用自增运算符，不存在的key会创建
	ages["ccc"]++
	ages["ccc"] += 1
	fmt.Printf("%v  %v\n",ages,ages["ccc"])

	// 元素不是变量，不能取址。
	//_ = &ages["ccc"]

	ages["baf"] = 200

	// 遍历的顺序不确定
	fmt.Println("==========")
	for name, age := range ages {
		fmt.Printf("%s  \t%d\n", name, age)
	}
	fmt.Println("==========")
	for name, age := range ages {
		fmt.Printf("%s  \t%d\n", name, age)
	}
	fmt.Println("==========")

	// 排序key
	names := make([]string, len(ages))
	idx := 0
	for name,_ := range(ages) {
		names[idx] = name
		idx ++
	}
	sort.Strings(names)
	fmt.Printf("%v\n",names)

	fmt.Println("==========")
	for _, name := range names {
		fmt.Printf("%s  \t%d\n", name, ages[name])
	}
	fmt.Println("==========")

	var tmp map[string]int
	fmt.Println(tmp == nil)
	fmt.Println(len(tmp) == 0)
	//tmp["hello"] = 0			// panic

	fmt.Println(ages["cdf"])	// 不存在的key得到零值
	// 如果本身就存在value为0的情况，可以用这种方式
	if age, ok := ages["cdf"]; !ok {
		fmt.Println("cdf 不存在")
	} else {
		fmt.Println(age)
	}
}
