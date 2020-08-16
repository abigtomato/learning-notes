import time
import multiprocessing


"""
    两个进程间的通信队列
"""


def producer(pipe):
    pipe.send('2233')


def consumer(pipe):
    print(pipe.recv())


if __name__ == '__main__':
    """适用于两个进程间通信的队列"""
    recv_pipe, send_pipe = multiprocessing.Pipe()
    producer_process = multiprocessing.Process(target=producer, args=(send_pipe, ))
    consumer_process = multiprocessing.Process(target=consumer, args=(recv_pipe, ))
    
    producer_process.start()
    producer_process.join()
    consumer_process.start()
    consumer_process.join()