/* 使用自定义参数类型时，只需要实现flag.Value接口
* 其String() string方法是个fmt.Stringer
* 它的Set方法解析它的字符串参数并且更新变量值
 */

package main

import (
	"flag"
	"fmt"

	tempconv "gopl.io/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
