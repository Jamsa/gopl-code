/*接口值
* 接口值由两部分组成：接口类型和那个类型的值
* 它们被称为接口的动态类型和动态值。类型是编译期的概念。
* 接口与变量一样总是被一个定义明确的值所初始化。
* 接口的零值就是它的类型和值部分都是nil，调用空接口值上的任意方法都会产生panic
* 包含nil指针的接口不是nil接口
 */
package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	//接口值可以持有任意大的动态值
	var x interface{} = []int{1, 2, 3}
	//接口值用==和!=进行比较
	fmt.Println(x == x)

	var buf *bytes.Buffer //将它修改为非指针类型，不会产生这个错误
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf) //注意：debug=false时会报错，因为buf为nil指针，而不是nil接口
	if debug {

	}
}

// 如果 out不是nil，将输出信息至out
func f(out io.Writer) {
	//这里的判断不会起到保护作用。
	// 在运行时out是 *bytes.Buffer指针类型，指针的值是nil，out本身不是nil
	// 即：包含nil指针的接口，不是nil接口
	// (接口值由两部分组成，类型和值，这里指值是nil)
	if out != nil {
		out.Write([]byte("done!\n")) //panic：nil指针解引用报错
	}
}
