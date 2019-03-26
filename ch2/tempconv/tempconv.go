// 包
// 变量声明、常量和方法放在此文件
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius  = 0
	BoilingC Celsius = 100
)

// String ...
func (c Celsius) String() string {
	return fmt.Sprintf("%g ºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g ºF", f)
}
