/* 递归演示
*
* 使用：
* go build gopl.io/ch1/fetch
* go build gopl.io/ch5/outline
* ./fetch http://golang.org | ./outline
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	// 将每个元素推入堆栈
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) //n.Data为元素名
		fmt.Println(stack)
	}

	// 遍历 n 的子元素，将它收集到 links
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// stack 传递的是值，是当前值的拷贝。上面的 append 操作不会影响当前循环其它循环次数的值
		outline(stack, c)
	}
}
