// 缓存函数
// 非并行安全版本
package memo

//缓存结构
type Memo struct {
	f     Func
	cache map[string]result
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
// 对memo.cache的访问是非并行安全的
func (memo *Memo) Get(key string)(interface{}, error){
	res, ok := memo.cache[key]
	if !ok{
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
