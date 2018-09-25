// 并发获取url的内容

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"time"
	"io"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:]{
		go fetch(url, ch) 		// 启动goroutine
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)		// 从chnnel接收
	}

	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}

// fetch ...
func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)	// 发送至channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url,err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs,nbytes,url)
}
