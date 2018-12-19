// 华氏摄氏温度转换

package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g华氏度 = %g摄氏度\n", freezingF, fToC(freezingF))
	fmt.Printf("%g华氏度 = %g 摄氏度\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
