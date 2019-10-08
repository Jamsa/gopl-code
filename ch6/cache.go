//嵌入结构体演示 cache 类型
package main

import (
	"fmt"
	"sync"
)

// cache 缓存对象定义
var cache = struct {
	sync.Mutex //嵌入结构体
	mapping    map[string]string
}{
	mapping: make(map[string]string),
}

// Lookup 查找缓存数据
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["aa"] = "aa"
	fmt.Println(Lookup("aa"))
}
