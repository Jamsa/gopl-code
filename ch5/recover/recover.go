/*
* recover的使用
* 如果在defered函数中调用了内置函数recover，并且定义了该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。
* 导致panic异常的函数不会继续运行，量能正常返回。未发生panic时调用recover时，recover会返回nil。
*
* 使用recover的一些惯例：
* 1. 不应该不加区分的恢复所有panic，因为panic后，无法保证包级变量的状态仍然与我们预期的一致。
* 2. 不应该试图去恢复其他包引起的panic，因为你无法确保这样做是否安全。
* 3. 公有API应该将函数的运行失败当作error返回，而不是panic。
* 4. 安全的做法是有选择性的recover。将panic value设置为特殊值，在处理时用于判断panic的类型，选择性的进行recover。
 */
package main

import (
	"fmt"
)

func testRecover(input string) (result string, err error) {
	defer func() {
		if p := recover(); p != nil { //recover panic
			err = fmt.Errorf("内部错误:%v", p)
		}
	}()

	panic("error message") //引发panic
	//return input + " result", nil
}

func main() {
	result, err := testRecover("hello")

	if err != nil {
		fmt.Printf("%v\n", err) //输出的是recover中返回的错误信息
		return
	}

	fmt.Printf("%v\n", result)
}
