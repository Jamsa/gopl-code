/* 递归演示，带多值返回
* 不再依赖 fetch.go
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
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
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
	return visit(nil, doc), nil
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

/*
bare return:
如果一个函数所有的返回值都有显式的变量名，那么函数的 return 语句，可以省略操作数，如：
func aa(url string)(words, images int, err error){
	resp, err := http.Get(url)
	if err != nil{
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close() //确保资源被关闭
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	words, images = countXXX(doc)
	return
}
*/
