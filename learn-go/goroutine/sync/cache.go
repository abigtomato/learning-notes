package main

import (
	"fmt"
	"sync"
	"ioutil"
	"net/http"
)

/* 并发非阻塞缓存示例（互斥锁版） */

// 需要缓存的函数结构
type Func func(key string) (interface{}, error)

// 函数缓存
type result struct {
	value 	interface{}
	err 	error
}

// 带条件的函数缓存，其中ready通道用于重复抑制
type entry struct {
	res 	result
	ready 	chan struct{}
}

// 函数记忆（对开销高昂的函数进行缓存）
type Memo struct {
	f 		Func
	mu 		sync.Mutex
	cache 	map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{
		f: 		f,
		cache: 	make(map[string]*entry),
	}
}

// 缓存获取
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()	// 用互斥量保护被多个go程访问的map
	e := memo.cache[key]
	
	if e == nil {
		// 1.entry不存在的情况则代表函数第一次执行，需要建立缓存
		e = &entry{
			ready: make(chan struct{}),
		}
		memo.cache[key] = e	// 该语句在互斥锁作用域中，是为了避免多go程出现竞态问题
		memo.mu.Unlock()

		// 调用函数获取结果
		e.res.value, e.res.err = memo.f(key)	// 互斥锁的作用域不包括该语句，因为需要此操作并发执行

		// 当函数调用完毕后再关闭ready通道，是为了让其他go程不重复调用函数
		close(e.ready)
	} else {
		memo.mu.Unlock()
		/*
			2.entry已存在的情况：
				2.1 可能表示内部的值还没准备好（另一个go程可能还在调用f）；
				2.2 需要等待entry准备好（也就是ready被close）才能继续执行读取entry中的result数据。
		*/
		<-e.ready
	}

	return e.res.value, e.res.err
}

// 需要被缓存的函数
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	return ioutil.ReadAll(resp.Body)
}

// 生成用于http get请求的url
func incomingURLs() []string {
	return []string{"http://www.baidu.com"}
}

func main() {
	m := memo.New(httpGetBody)
	var wg sync.WaitGroup
	
	for url := range incomingURLs() {
		wg.Add(1)
		
		// 获取函数缓存的go程
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