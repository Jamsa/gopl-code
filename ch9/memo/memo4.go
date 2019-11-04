// 缓存函数
// 带并行安全版本的版本
// 通过channel减少一次加锁
package memo

type entry struct{
	res result
	ready chan struct{}
}

//缓存结构
type Memo struct {
	f     Func
	cache map[string]result
	mu sync.Mutex
}

//被缓存函数
type Func func(key string) (interface{}, error)

//返回值结构
type result struct {
	value interface{}
	err   error
}

//创建带缓存函数
f New(f Func) *Memo{
	return &Memo{f: f, cache: make(map[string]entry)}
}
// 对memo.cache的访问是并行安全的
func (memo *Memo) Get(key string)(interface{}, error){
	memo.mu.Lock()
	e  := memo.cache[key]
	//memo.mu.Unlock() 不释放
	if e == nil{
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()// 产生entry后就可以释放锁了

		e.res.value, e.res.err = memo.f(key)
		//memo.mu.Lock() 不需要再次加锁
		close(e.ready)// 产生ready消息
	}else{
		memo.mu.Unlock()
		// 等侍ready消息，这个channel可以被多个客户端访问，而不需要互斥
		<-e.ready 
	}
	memo.mu.Unlock()
	return res.value, res.err
}
