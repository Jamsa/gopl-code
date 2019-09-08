// 方法与函数
package main

import (
	"fmt"
	geo "gopl.io/ch6/geometry"
)

type Path []geo.Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := geo.Point{1, 2}
	q := geo.Point{3, 4}

	fmt.Println(geo.Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}
