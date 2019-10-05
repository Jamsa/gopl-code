/*
* 通过嵌入结构体来扩展类型
 */
package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point // 将Point类型嵌入进来，提供X，Y属性
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1 //调用时不需要指定类型，直接当作自己的字段使用
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

}
