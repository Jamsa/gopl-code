/*
结构体是聚合数据类型，由零个或多个任意类型的值聚合成。
结构体变量的成员用点操作符访问。点操作符也可以和指向结构体的指针一起工作。
*/

package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	var dilbert Employee
	// 点访问
	dilbert.Salary -= 5000
	
	// 指针访问
	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Println(dilbert)

	// 结构体变量指针，可以用点操作符
	var emp *Employee = &dilbert
	emp.Position += " (aafdsafda)"
	// 相当于(*emp).Position += ...
	fmt.Println(*emp)
}
