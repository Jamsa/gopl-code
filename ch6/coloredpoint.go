/*
* 通过嵌入结构体来扩展类型
 */
package main

import (
	"fmt"
	"image/color"

	geo "gopl.io/ch6/geometry"
)

// ColoredPoint 彩色Point
type ColoredPoint struct {
	geo.Point // 将Point类型嵌入进来，提供X，Y属性
	Color     color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1 //调用时不需要指定类型，直接当作自己的字段使用
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geo.Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{geo.Point{X: 5, Y: 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	//方法值：引用值上的方法
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q.Point))

	//方法表达式：引用类型上的方法
	distance := geo.Point.Distance
	fmt.Println(distance(p.Point, q.Point))

}
