/*
结构体是聚合数据类型，由零个或多个任意类型的值聚合成。
结构体变量的成员用点操作符访问。点操作符也可以和指向结构体的指针一起工作。
结构体成员名以大写开头的会被导出。

名为S的结构体类型不能再包含S类型的成员，但是可以包含 *S指针类型的成员，以便创建递归数据结构。

没有成员的结构体写作struct{}。它的大小为0，有些时候可以用它作为map的value来模拟set。

结构体字面值可以按顺序指定每个成员的值，也可以用成员名和值来初始化部分成员的值。两种方式不能混合使用。

如果成员都可以可比较的，那么结构体也可比较。==运算符将比较两个结构体的每个成员。

匿名成员：匿名成员的数据类型必须是命名的类型或指向一个命名类型的指针。访问匿名类型的成员时，不需要给出完整的路径。使用完整路径访问时，匿名成员的名字就是类型名。
匿名成员不能使用字面值表达。

因为类型名将作为匿名成员的隐式名字，因此不能同时包含两个类型相同的匿名成员，这将导致名字冲突。
即使不匿名，将成员名用小写字母开头，在包内部我们依然可以用简短形式访问匿名成员嵌套的成员。如 w.X=8代替w.circle.point.X = 8。

将没有任何成员的类型作为匿名成员是为了使用匿名成员类型的方法集。
*/

package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	//Name string 同类型的可以合并在一行定义
	Address, Name string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

type Point struct{X, Y int}

type Circle struct {
	Point						// 匿名成员
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var dilbert Employee
	// 点访问
	dilbert.Salary -= 5000
	
	// 获取成员指针
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert)

	// 结构体变量指针，可以用点操作符
	var emp *Employee = &dilbert
	emp.Position += " (aafdsafda)"
	// 相当于(*emp).Position += ...
	fmt.Println(*emp)

	fmt.Println(EmployeeById(100).Position)

	p := Point{1,2}
	fmt.Println(p)
	p = Point{X:20}
	fmt.Println(p)

	pp := &Point{1,2}
	fmt.Println(*pp)
	// 与下面的方式等价
	pp = new(Point)			// 不分配具体空间
	*pp = Point{1,2}
	fmt.Println(*pp)

	// 匿名成员
	var w Wheel
	w.X = 8						// 等价于完整路径访问：w.Circle.Point.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	// 匿名成员不能使用字面值表达
	//w = Wheel{8,8,5,20}			// 编译错误
	// 需要写为
	w = Wheel{Circle{Point{8,8},5},20}
	fmt.Printf("%#v\n", w)
	// 或
	w = Wheel{
		Circle: Circle {
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	w.X = 42
	fmt.Printf("%#v\n", w)
}

// EmployeeById 通过id获取员工信息结构体指针
func EmployeeById(id int) *Employee {
	var dilbert Employee
	dilbert.ID = id
	dilbert.Name = "Emp100"
	dilbert.Position = "Manager"
	return &dilbert
}

// 二叉树排序

type tree struct {
	value int
	left, right *tree
}

// Sort in place排序
func Sort(values []int)  {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues 按顺序将t中的元素添加到values中
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add ...
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
