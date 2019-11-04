// 缓存函数
// 带并行安全版本的版本
// 但是每次调用Get时都将被锁，对函数的调用变成了串行，损失了性能
package memo

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
	return &Memo{f: f, cache: make(map[string]result)}
}
// 对memo.cache的访问是并行安全的
func (memo *Memo) Get(key string)(interface{}, error){
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok{
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
