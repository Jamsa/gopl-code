// 缓存函数
// 带并行安全版本的版本
// 完全不使用锁
package memo

//代表函数调用请求
type request struct{
	key string
	response chan<- result
}

type entry struct{
	res result
	ready chan struct{}
}

//缓存结构
type Memo struct {
	requests chan request
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
	memo := &Memo{requests: make(chan request)} 
	go memo.server(f)
	return memo
}
// 对memo.cache的访问是并行安全的
func (memo *Memo) Get(key string)(interface{}, error){
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.alue, res.err
}

func (memo *Memo)Close(){ close(memo.requests) }

func (memo *Memo) server(f Func){
	cache := make(map[string]*entry)
	for req := range memo.requests{
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string){
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result){
	<- e.ready
	response <- e.res
}
