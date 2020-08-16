import threading


"""
    互斥锁，生产消费模式
"""


total = 0


# 生产者线程
class Produce(threading.Thread):
    def __init__(self, thread_name, lock):
        super().__init__(name=thread_name)
        self.lock = lock

    def run(self):
        global total
        for i in range(10000):
            # 上锁操作，如果当前锁没有被标记为上锁状态，当前对象在此锁上挂标记
            # 若之前已经被上锁，当前执行到这里的线程会阻塞等待锁的释放
            self.lock.acquire()
            total += i
            # 释放锁标记，代表此时锁可被任意线程上锁
            self.lock.release()


# 消费者线程
class Consumer(threading.Thread):
    def __init__(self, thread_name, lock):
        super().__init__(name=thread_name)
        self.lock = lock

    def run(self):
        global total
        for i in range(10000):
            self.lock.acquire()
            total -= i
            self.lock.release()


if __name__ == '__main__':
    # 互斥锁
    lock = threading.Lock()
    
    produce = Produce('produce', lock)
    consumer = Consumer('consumer', lock)

    # 生产者线程
    produce.start()
    produce.join()
    
    # 消费者线程
    consumer.start()
    consumer.join()

    print("total: {total}".format(total=total))