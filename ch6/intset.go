/*
* IntSet 示例
* 每个word都有64 bit，用x/64的同作为位置下标，用x^64作为bit所在的位置
*
* 封装：大写首字母的标识符会从定义它们的包中被导出，上写字母的则不会。
* 这种基于名字的手段使得在语言中最小的封装单元是package，而不是类型。
* 一个struct类型的字段对同一个包的所有代码都是可见的，无论你的代码是写在一个函数还是一个方法里
 */
package main

import (
	"bytes"
	"fmt"
)

// IntSet 非负整数集合
// 它的零值表示空集合
type IntSet struct {
	words []uint64
}

// Has 检测集合中是否包含某个非负整数值
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加非负值至集合
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 合并两个集合
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 返回字符串表达
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(0)

	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}
