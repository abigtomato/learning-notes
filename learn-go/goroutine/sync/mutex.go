package main

import (
	"fmt"
	"sync"
	"time"
)

/* 传统的同步机制（互斥锁） */

// 可以进行原子操作的Int
type AtomicInt struct {
	value 	int
	lock 	sync.Mutex		// 互斥锁
}

// 原子自增
func (this *AtomicInt) increment() {
	// 若想在函数中为一段代码加锁，可以使用匿名函数实现
	func() {
		this.lock.Lock()
		defer this.lock.Unlock()
		this.value++
	}()
}

// 原子获取
func (this *AtomicInt) get() int {
	// 为当前go程加锁，其他所有go程执行到此不能获取锁进入阻塞
	this.lock.Lock()	// 建议锁: 操作系统提供，建议在编程时使用的锁
	// 函数结束释放锁（自动唤醒阻塞在这把锁上的所有go程）
	defer this.lock.Unlock()
	return this.value
}

func main() {
	var a AtomicInt

	// 并发读写可能会发生同步问题
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}