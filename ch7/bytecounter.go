/*
* 接口示例
 */
package main

import "fmt"

// ByteCounter 位计数器
type ByteCounter int

// Write 方法满足 io.Writer类型
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

//类似的还有String() string方法，它满足了fmt.Stringer接口类型的要求

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 //重置计数
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
