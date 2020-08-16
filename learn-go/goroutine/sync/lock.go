package main

import (
	"fmt"
	"time"
	"sync"
)

/* 全局互斥锁 */

var (
	/*
		1.全局资源，开启20条协程计算1-20的阶乘，将结果写入Map中；
		2.若是没有互斥锁，多个协程操作一个资源会产生资源竞争问题。
	*/
	myMap = make(map[int]int, 20)
	
	/*
		全局互斥锁：
		1.可以对代码片段进行加锁，第一个协程获取锁对代码段操作；
		2.其他协程会尝试获取锁，但已经被第一个协程锁定，所以进入等待队列进行等待；
		3.当第一个协程执行完毕会释放锁，队列中的协程按出队顺序获取锁并执行，执行完之后再释放锁。
	*/
	lock sync.Mutex
)

func main() {
	for i := 1; i <= 20; i++ {
		// 开启Go协程
		go func(n int) {
			res := 1
			for i := 1; i <= n; i++ {
				res *= i
			}

			// 添加锁
			lock.Lock()
			// 释放锁
			defer lock.Unlock()

			myMap[n] = res
		}(i)
	}

	// 主线程睡眠等待全部协程结束（无法知道什么时候结束，只能设置大概等待时间）
	time.Sleep(time.Second * 5)

	// 主线程并不知道什么时候执行完全部的协程，可能会尝试去查看锁的状态，也会触发资源竞争，所有要添加互斥锁
	lock.Lock()
	defer lock.Unlock()

	for k, v := range myMap {
		fmt.Printf("[%v]=%v\n", k, v)
	}
}