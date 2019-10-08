package geometry

import "math"

type Point struct{ X, Y float64 }

//传统的函数定义
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Point类型的方法定义
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/*
* 基于指针对象的方法定义
* 它可以避免Point类型方法调用时复制参数对象的开销。
* 实际应用中，如果Point类型有一个指针作为接收器的方法，那么所有Point的方法都必须有指针接收器，即使是那些并不需要这个指针接收器的函数。（示例没有遵循这一习惯）。
 */
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
type P *int

// 编译错误：无效的接收器类型。
// 如果类型P本身就是指针，则P不允许出现在接收器中。
func (P) f() {

}
*/
