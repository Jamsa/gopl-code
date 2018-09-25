// 获取url的内容

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		// 自动添加http://
		if !strings.HasPrefix(url, "http://"){
			url = fmt.Sprintf("http://%s", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)

		fmt.Printf("===================================\n")

		resp, err = http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("状态码：%d\n", resp.StatusCode)
		// 另一种方式读取内容：使用io.Copy进行流复制
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
