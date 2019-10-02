/* 函数应用示例
* 拓扑排序示例
 */

package main

import (
	"fmt"
	"sort"
)

// 依赖关系数据
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structure":        {"discrete math"},
	"databases":             {"data structure"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, coursor := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, coursor)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	//深度优先遍历
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
