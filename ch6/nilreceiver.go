/*
* nil是合法的接收器类型
 */
package main

// IntList 是整数list，nil表示空list
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum 返回list中元素个数
func (list *IntList) Sum() int {

	//接收器可以为nil
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
