package main

import (
	"fmt"
)

/* 并发非阻塞缓存示例（channel版） */

// Func是用于记忆的函数类型
type Func func(key string) (interface{}, error)

// 调用Func的返回结果
type result struct {
	value 	interface{}
	err 	error
}

// 对Func返回结果的包装
type entry struct {
	res 	result
	ready 	chan struct{}	// 当res准备好后关闭ready通道
}

// request是一条请求消息，key需要用Func来调用
type request struct {
	key 		string
	response 	chan<- result	// 客户端需要单个result
}

// 函数记忆
type Memo struct {
	requests chan request
}

// New返回f的函数记忆，客户端之后需要调用Close
func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}

	// 开启服务
	go memo.server(f)
	return memo
}

// 获取缓存
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	defer close(response)

	// requests由server消费，用于执行慢函数f(key)
	memo.requests <- request{
		key: 		key, 
		response: 	response,
	}

	// 阻塞等待server将慢函数f的执行结果存入
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	// cache只由一个go程所监控，不存在竞态问题
	cache := make(map[string]*entry)
	
	// 消费requests通道，获取的request包含函数f的参数key和用于保存结果的通道response
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// entry不存在的情况，实例新entry并调用函数f并关闭ready通道
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e

			// 调用f(key)，调用成功后通知数据准备完毕
			go e.call(f, req.key)	
		}
		
		// 等待entry准备完毕发送结果给客户端
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待该数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
}

// 慢函数f
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

// 生成url
func incomingURLs() []string {
	return []string{"http://www.baidu.com"}
}

func main() {
	m := New(httpGetBody)
	defer m.Close()

	var wg sync.WaitGroup
	for url := range incomingURLs() {
		wg.Add(1)
		
		go func(url string) {
			start := time.Now()
			
			// 获取函数缓存
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			
			wg.Done()
		}(url)
	}

	wg.Wait()
}