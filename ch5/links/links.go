/* 函数演示
* 使用匿名函数 深度优先遍历
 */
package main

import (
	"fmt"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := Extract(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "links: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() //确保资源被关闭
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close() //确保资源被关闭
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		// 如果 n 是 html 的 a 元素
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	var forEachNode func(n *html.Node, v func(n *html.Node))
	forEachNode = func(n *html.Node, v func(n *html.Node)) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			v(c)
			forEachNode(c, v)
		}
	}

	forEachNode(doc, visitNode)

	return links, nil
}

/*
* 捕获迭代变量的问题
* 以下代码示例是为了创建一些目录之后，再通过匿名函数删除。
* 由于for循环语句会创建新的语法块，捕获了迭代变量，会导致删除操作只能删除最后一次循环中捕获的dir变量。
*
var rmdirs;
for _, dir := range tempDirs() {
	os.MkdirAll(dir,0755) 			//错误
	rmdirs = append(rmdirs, func(){
		os.RemoveAll(dir)			//捕获的是迭代变量（的地址），最终它会变成最后一轮循环的值
	})
}

for _, rmdir := range rmdirs {
	rmdir() //调用匿名函数删除目录
}
* 正确的做法是
var rmdirs;
for _, dir := range tempDirs() {
	dir := dir //使用这个额外定义的dir，它的作用域只在单次循环中，不是一个迭代变量
	os.MkdirAll(dir,0755)
	rmdirs = append(rmdirs, func(){
		os.RemoveAll(dir)
	})
}
*/
