// 并发的退出
// 利用channel被关闭后 <-操作始终不阻塞的原理进行广播
// 通过关闭done channel的方式终止正在进行的walkDir、dirents和主goroutine
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//目录遍历
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	//如果已经取消就不再继续遍历
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 程序高峰期可能会创建成百上千的goroutine
// sema 是用于限制并发数量的计数器，利用channel buffer
var sema = make(chan struct{}, 20)

//获取目录内容
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // 获取计数
	case <-done: //信号量获取后也可以被取消
		return nil
	}

	defer func() { <-sema }() //释放计数
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

// 查询退出状态
func cancelled() bool {
	select {
	// 被关闭的channel上读取永远不会阻塞，可以多次调用
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取1 byte数据
		close(done)                    // channel关闭时将产生广播信号，各个监听端都能得到消息
	}()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) //并行遍历目录
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			//排空fileSizes让已经启动的goroutines结束
			for range fileSizes {
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 已经被关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			// 每500ms打印一次
			printDiskUsage(nfiles, nbytes)
		}

	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
