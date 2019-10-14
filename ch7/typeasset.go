package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	// x.(T)类型断言，x表示一个接口类型和T表示类型。
	// 如果T是具体类型，检查x的类型和T是否相同。
	// 如果相同，则类型断言将得到它操作对象中的具体的值。
	// 如果检查失败，接下来的操作会抛出panic
	f := w.(*os.File) //success
	//c := w.(*bytes.Buffer) //panic
	println(f)
	//println(c)

	// 如果T是接口类型，则类型断言检查是否x的动态类型满足T。
	// 如果检查成功，结果仍然是一个有相同类型和值部分的接口值，但是结果有类型T。

}
