import time
from multiprocessing import Process, Queue, Manager, Pool, Pipe


"""
    进程间通信
"""


def my_producer(queue):
    queue.put('bilibili')
    time.sleep(2)


def my_consumer(queue):
    time.sleep(2)
    data = queue.get()
    print(data)


if __name__ == '__main__':
    # 使用消息队列完成进程间通信
    queue = Queue(10)
    producer = Process(target=my_producer, args=(queue, ))
    consumer = Process(target=my_consumer, args=(queue, ))
    
    producer.start()
    producer.join()
    consumer.start()
    consumer.join()