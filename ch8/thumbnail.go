//并发的循环

package main

import (
	"fmt"
	"log"
	"sync"
	"os"
)

// ImageFile 代表生成缩略图的函数
func ImageFile(infile string) (string, error) {
	fmt.Println("ImageFile:%s", infile)
	return infile, nil
}

// makeThumbnails 非并行
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// makeThumbnails2 易于并行的处理方式，只要在调用的时候在前面加个go关键字
// 这个版本的问题是：用go关键字调用它时，它很快会结束，工作没完成就返回了。
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) // 忽略了错误处理
	}
}

// makeThumbnails3
// 通过channels等侍各个goroutines执行完毕
// 缺点是各个goroutines无法向主goroutine返回值
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// 这里使用f参数，而不是闭包中直接使用f，
		// 是因为循环变量f是被所有匿名函数共享的，会被连续迭代所更新
		go func(f string) {
			ImageFile(f) // 忽略了错误处理
			ch <- struct{}{}
		}(f)
	}

	//等侍goroutines结束
	for range filenames {
		<-ch
	}
}

// makeThumbnails4
// 从goroutines通过channel返回错误信息至主goroutine
// 问题是可能造成goroutines泄漏
func makeThumbnails4(filenames []string) {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	//等侍goroutines结束
	for range filenames {
		if err := <-errors; err != nil {
			// 可能造成goroutines泄漏。
			// 当channels的存在非nil值时，函数返回。
			// 没有机会继续排空errors channel。其余的goroutines在向
			// 这个channel发送值时，都会永远地阻塞下去，不会退出。
			return err
		}
	}
	return nil
}

// makeThumbnails5
// 带有完善的从goroutines返回至主goroutine的机制
// 通过创建带缓存的channels避免goroutines泄漏
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	//与filenames大小相同的channel，避免被阻塞
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err := ImageFile(f)
			ch <- it
		}(f)
	}

	//等侍goroutines结束
	for range filenames {
		it := <- ch
		if it.err != nil{
			thumbfiles = append(thumbfiles,it.thumbfile)
		}
	}
	return thumbfiles, nil
}


// makeThumbnails6
// 能返回文件总大小的版本
// 
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup //goroutines计数

	for f := range filenames {
		wg.Add(1) //为计数器加1
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // 忽略错误
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func(){
		wg.Wait()
		close(sizes) //最后结束位置
	}()

	var total int64
	for size := range sizes { // 循环在整个程序执行过程中持续运行
		total += size
	}
	return total
}

func main() {

}
