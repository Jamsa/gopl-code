/* 函数参数演示
*
* 使用：
* outline2 https://golang.org
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get error:%v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() //确保资源被关闭
		fmt.Fprintf(os.Stderr, "getting %s: %s", url, resp.Status)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close() //确保资源被关闭
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %s as HTML: %v", url, err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}
