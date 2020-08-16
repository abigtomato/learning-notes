from queue import Queue
import time
import threading


"""
    使用队列实现线程间通信
"""

# 模拟抓取url存入队列
def get_detail_url(queue, lock):
    while True:
        print("get_detail_html start")
        time.sleep(2)
        for elem in range(20):
            # 为存储操作添加互斥锁
            lock.acquire()
            queue.put("http://pan.baidu.com/s/{id}".format(id=elem))
            lock.release()
        print("get_detail_html end")


# 模拟从队列中取出url请求获取html
def get_detail_html(queue, lock):
    while True:
        # 为读取操作添加互斥锁
        lock.acquire()
        url = queue.get()
        lock.release()

        print("get_detail_html start : {url}".format(url=url))
        time.sleep(2)
        print("get_detail_html end")


if __name__ == '__main__':
    # 实例一个用于存储url的队列，最大容量maxsize设置为1000
    detail_url_queue = Queue(maxsize=1000)
    lock = threading.Lock()

    start_time = time.time()

    # 开启一个线程执行url抓取操作
    threading.Thread(target=get_detail_url, args=(detail_url_queue, lock)).start()
    # 开启10个线程执行取出url获取html的操作
    for i in range(10):
        threading.Thread(target=get_detail_html, args=(detail_url_queue, lock)).start()

    # detail_url_queue.task_done()
    # detail_url_queue.join()

    end_time = time.time()

    print("last time : {time}".format(time=end_time-start_time))