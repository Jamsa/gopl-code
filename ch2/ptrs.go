// 指针
package main
import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var x1,y int
	// 只有指针指向同一变量，或者全部是nil时才相等
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil)

	//var p1 = f()
	fmt.Println(f() == f())		// 每次返回的指针是不同的

	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

// 返回局部变量的地址是安全的
func f() *int {
	v := 1
	return &v
}

// 通过指针修改参数，不推荐的做法
// *p是变量v的别名。很多引用类型，如slice,map,chan，甚至是结构体、数组、接口都会创建所引用变量的别名
func incr(p *int) int {
	// 只是修改指针指向的值，并不是指针运算
	*p++
	return *p
}
