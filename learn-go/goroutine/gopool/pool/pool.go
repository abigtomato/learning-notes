package pool

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
)

// 表示pool关闭状态的常量
const CLOSED = 1

var (
	chanSize = func() int {
		// 如果GOMAXPROCS为1时，使用阻塞channel
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// 如果GOMAXPROCS大于1时，使用非阻塞channel
		return 1
	}()
	PoolClosedError      = errors.New("this pool has been closed")
	InvalidPoolSizeError = errors.New("invalid size for pool")
)

// go程池（调度器）
type Pool struct {
	cap          int32        		// 池容量
	closed       int32       		// 关闭标记
	jobQueue     chan Job     		// 总任务队列
	workerQueue  chan *Worker 		// worker队列
	once         sync.Once	  		// 保证某些操作只执行一次
	PanicHandler func(interface{})	// 用户自定义错误处理
}

// 实例新pool
func NewPool(size int) (*Pool, error) {
	if size <= 0 {
		return nil, InvalidPoolSizeError
	}

	pool := &Pool{
		cap:      int32(size),
		jobQueue: make(chan Job, chanSize),
	}

	if chanSize != 0 {
		pool.workerQueue = make(chan *Worker, size)
	} else {
		pool.workerQueue = make(chan *Worker, chanSize)
	}

	return pool, nil
}

// 初始化pool
func initPool(size int) *Pool {
	pool, _ := NewPool(size)
	pool.run()
	
	return pool
}

// 提交任务给pool
func (p *Pool) submit(job Job) error {
	// 判断pool是否已经关闭
	if atomic.LoadInt32(&p.closed) == CLOSED {
		return PoolClosedError
	}

	p.jobQueue <- job	// 任务入队

	return nil
}

// 启动pool
func (p *Pool) run() {
	// 根据最大pool容量创建worker
	for i := 0; i < int(p.cap); i++ {
		worker := NewWorker(p)		// 创建worker实例
		go worker.start()			// 开启go程分支执行worker逻辑
		p.workerQueue <- worker		// worker入队
	}
	
	go p.scheduler()
}

// 开启调度
func (p *Pool) scheduler() {
	for {
		select {
		case job := <-p.jobQueue:		// 监听任务队列的数据
			worker := <-p.workerQueue	// 若有任务需要处理，出队一个worker进行处理
			worker.task <- job			// 存入worker的专属任务队列中
		}
	}
}

// 关闭pool
func (p *Pool) close() {
	/*
		1.Once包含一个bool变量和一个互斥量，bool变量记录逻辑初始化是否完成，互斥量负责保护bool变量和客户端的数据结构；
		2.Once的唯一方法Do以需要执行的初始化函数作为参数；
		3.每次调用Do时会先锁定互斥量并检查里边的bool变量，第一次调用时bool变量为false；
		4.Do会调用初始化函数并将变量置为true，后续的再次调用相当于空操作。
	*/
	p.once.Do(func() {
		defer func() {
			close(p.jobQueue)
			close(p.workerQueue)
		}()
		
		atomic.StoreInt32(&p.closed, 1)

		p.jobQueue = nil
		p.workerQueue = nil
	})
}