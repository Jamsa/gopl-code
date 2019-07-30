/* 递归演示
* go 使用可变栈，栈的大小会按需增加（初始值很小）。这使得我们使用递归时不需要考虑溢出和安全问题
* 使用：
* go build gopl.io/ch1/fetch
* go build gopl.io/ch5/findlinks1
* ./fetch http://golang.org | ./findlinks
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

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	// 如果 n 是 html 的 a 元素
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// 如果这个元素有 href 属性
				links = append(links, a.Val)
			}
		}
	}

	// 遍历 n 的子元素，将它收集到 links
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
