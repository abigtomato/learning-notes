package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

/* 条件变量 */

var cond sync.Cond

func producer(out chan<- int, idx int) {
	for {
		func() {
			// 使用条件变量添加互斥锁
			cond.L.Lock()
			defer cond.L.Unlock()

			/*
				使用for循环判断条件变量是否满足（而不是使用if）:
				1.设想使用if的场景：第一个进来的go程通过了if判断，还未执行wait就失去了cpu的使用权；
				2.其他go程抢占到cpu则会循环再次判断是否满足，若是if其他go程会直接向下执行（因为第一个进来的go程已经通过了判断）。
			*/
			// 管道中的数据满了，就使当前的生产go程wait
			for len(out) == cap(out) {
				/*	
					Wait(): 
					1.使当前go程在条件变量cond上阻塞并等待条件变量的唤醒；
					2.释放当前go程持有的互斥锁（相当于进行cond.L.Unlock()），和第1步同为1个原子操作；
					3.当阻塞在此的go程被唤醒，Wait()函数返回时，该go程重新获取互斥锁（相当于进行cond.L.Lock()操作）。
				*/
				cond.Wait()
			}

			num := rand.Intn(1000)
			out <- num
			fmt.Printf("生产者%d号 -> %d\n", idx, num)
		}()

		// Signal(): 给一个在当前条件变量上阻塞的go程发送唤醒通知
		cond.Signal()
		time.Sleep(time.Second)
	}
}

func consumer(in <-chan int, idx int) {
	for {
		func() {
			cond.L.Lock()
			defer cond.L.Unlock()

			// 管道没有数据可以消费了，就使当前消费go程wait
			for len(in) == 0 {
				cond.Wait()
			}

			num := <-in
			fmt.Printf("消费者%d号 <- %d\n", idx, num)
		}()

		cond.Signal()
		time.Sleep(time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	/*
		type Cond struct {
			L Locker
			notify notifyList
			checker copyChecker
		}
	 */
	cond.L = new(sync.Mutex)	// 使用互斥锁初始化条件变量的锁字段
	
	// 数据管道
	numChan := make(chan int, 3)
	// 退出标记管道
	quitChan := make(chan bool)
	
	for i := 0; i < 5; i++ {
		go producer(numChan, i)
	}

	for i := 0; i < 5; i++ {
		go consumer(numChan, i)
	}

	for {
		_, ok := <-quitChan
		if !ok {
			break
		}
	}
}