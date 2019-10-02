/*  defer 学习

 */
package main

import (
	"log"
	"os"
	"time"
)

// 利用defer输出函数的进入、退出、耗时日志信息
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s(%s)", msg, time.Since(start))
	}
}

// 由于defer语句中的函数会在return语句更新返回值变量后再执行，利用defer机制观察返回值，甚至修改返回值
func double(x int) (result int) {
	defer func() { log.Printf("double(%d) = %d\n", x, result); result++ }()
	return x + x
}

//循环体中的defer语句需要特别注意，只有在函数执行完后，这些被延迟的函数才会执行。
//以下的代码有可能耗尽文件描述符，在函数退出前，没有文件会被关闭
func fileDeferClose1(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() //这里可能耗尽文件描述符
	}
}

//一种解决方法是将Defer移到另一个函数
func fileDeferClose2(filenames []string) error {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
}

//函数返回时文件会被关闭
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
}

func main() {
	bigSlowOperation()

	log.Printf("%d\n", double(20))
}
