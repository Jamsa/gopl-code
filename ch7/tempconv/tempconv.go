/*Package tempconv 温度转换示例*/
package tempconv

import (
	"flag"
	"fmt"

	tp "gopl.io/ch2/tempconv"
)

type celsiusFlag struct{ tp.Celsius }

// CelsiusFlag 自定义的 flag 类型
// 参数必须按 数量值和单位的格式输入，如："100C"
func CelsiusFlag(name string, value tp.Celsius, usage string) *tp.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "ºC":
		f.Celsius = tp.Celsius(value)
		return nil
	case "F", "ºF":
		f.Celsius = tp.FToC(tp.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
