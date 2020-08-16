package pool

import "log"

type Job func()

// 工作节点
type Worker struct {
	pool *Pool     // 所属的池
	task chan Job  // 每个worker专属的任务队列
	quit chan bool // 退出标记管道
}

// 实例新工作节点
func NewWorker(pool *Pool) *Worker {
	return &Worker{
		pool: pool,
		task: make(chan Job, chanSize),
		quit: make(chan bool, 0),
	}
}

// 工作节点开始干活
func (w *Worker) start() {
	for {
		select {
		case job := <-w.task:	// task是每个worker专属的任务队列
			job()	// 执行任务的具体逻辑

			w.pool.workerQueue <- w	// worker执行完毕后重新入队pool
			
			// go程执行中的错误处理
			if p := recover(); p != nil {
				// 若用户自定义了错误处理函数则执行
				if w.pool.PanicHandler != nil {
					w.pool.PanicHandler(p)
				} else {
					// 否则默认错误处理
					log.Printf("worker exits from a panic: %v", p)
				}
			}
		case <-w.quit:
			return
		}
	}
}

// 停止worker
func (w *Worker) stop() {
	go func() {
		w.quit <- true	// 存入结束标记
	}()
}