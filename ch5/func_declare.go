package main

import "fmt"

/*
函数声明
包括函数名、形参、返回值列表(可省略)及函数体
函数返回无名变量或无返回值时，返回值的括号可以省略
*/

/* 一组形参或返回值的类型相同时，不必为每个形参都写出参数类型，下面两个写法是等价的 */

func f1(i, j, k int, s, t string)                {}
func f2(i int, j int, k int, s string, t string) {}

func add(x, y int) int     { return x + y }
func sub(x, y int) (z int) { z = x + y; return }
func first(x, _ int) int   { return x } /* _ 为blank identifier，可以强调某个参数未被使用 */
func zero(int, int) int    { return 0 }

func main() {
	/*
		函数的类型被称为函数的标识符。如果两个函数形式参数和返回值列表上的变量类型一一对应，那么这两个函数被认为有相同的类型
		和标识符。
	*/
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

	/*
		每次函数调用都必须按声明顺序为所有参数提供实参。Go语言没有默认参数值，也不能通过参数名指定形参，形参和返回值变量名
		对于函数调用者没有意义。
		在函数体中形参被作为函数的局部变量，被初始化为调用者提供的值。形参有名返回值作为函数最外层的局部变量，被存储在相同
		的语法块中。
		实参通过值的方式传递，形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针、切片、map、
		function、channel等类型，实参可能会由于函数的间接引用被修改。
		没有函数体的函数声明表示该函数不是以Go实现的，这样的声明定义了函数标识。
	*/
}
